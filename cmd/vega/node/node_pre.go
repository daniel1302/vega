package node

import (
	"context"
	"fmt"

	"code.vegaprotocol.io/vega/checkpoint"
	"code.vegaprotocol.io/vega/rewards"

	proto "code.vegaprotocol.io/protos/vega"
	"code.vegaprotocol.io/vega/accounts"
	"code.vegaprotocol.io/vega/assets"
	"code.vegaprotocol.io/vega/banking"
	"code.vegaprotocol.io/vega/blockchain"
	"code.vegaprotocol.io/vega/blockchain/abci"
	"code.vegaprotocol.io/vega/blockchain/recorder"
	"code.vegaprotocol.io/vega/broker"
	"code.vegaprotocol.io/vega/candles"
	"code.vegaprotocol.io/vega/collateral"
	"code.vegaprotocol.io/vega/config"
	"code.vegaprotocol.io/vega/delegation"
	"code.vegaprotocol.io/vega/epochtime"
	"code.vegaprotocol.io/vega/evtforward"
	"code.vegaprotocol.io/vega/execution"
	"code.vegaprotocol.io/vega/fee"
	"code.vegaprotocol.io/vega/genesis"
	"code.vegaprotocol.io/vega/governance"
	"code.vegaprotocol.io/vega/limits"
	"code.vegaprotocol.io/vega/liquidity"
	"code.vegaprotocol.io/vega/logging"
	"code.vegaprotocol.io/vega/markets"
	"code.vegaprotocol.io/vega/netparams"
	"code.vegaprotocol.io/vega/netparams/checks"
	"code.vegaprotocol.io/vega/netparams/dispatch"
	"code.vegaprotocol.io/vega/nodewallet"
	"code.vegaprotocol.io/vega/notary"
	"code.vegaprotocol.io/vega/oracles"
	oracleAdaptors "code.vegaprotocol.io/vega/oracles/adaptors"
	"code.vegaprotocol.io/vega/orders"
	"code.vegaprotocol.io/vega/parties"
	"code.vegaprotocol.io/vega/plugins"
	"code.vegaprotocol.io/vega/pprof"
	"code.vegaprotocol.io/vega/processor"
	"code.vegaprotocol.io/vega/risk"
	"code.vegaprotocol.io/vega/staking"
	"code.vegaprotocol.io/vega/stats"
	"code.vegaprotocol.io/vega/storage"
	"code.vegaprotocol.io/vega/subscribers"
	"code.vegaprotocol.io/vega/trades"
	"code.vegaprotocol.io/vega/transfers"
	"code.vegaprotocol.io/vega/types"
	"code.vegaprotocol.io/vega/validators"
	"code.vegaprotocol.io/vega/vegatime"

	"github.com/cenkalti/backoff"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/prometheus/common/log"
	"github.com/spf13/afero"
	tmtypes "github.com/tendermint/tendermint/abci/types"
)

func (l *NodeCommand) persistentPre(args []string) (err error) {
	// this shouldn't happen...
	if l.cancel != nil {
		l.cancel()
	}
	// ensure we cancel the context on error
	defer func() {
		if err != nil {
			l.cancel()
		}
	}()
	l.ctx, l.cancel = context.WithCancel(context.Background())

	conf := l.cfgwatchr.Get()

	if flagProvided("--no-chain") {
		conf.Blockchain.ChainProvider = "noop"
	}

	if flagProvided("--no-stores") {
		conf.StoresEnabled = false
	}

	// reload logger with the setup from configuration
	l.Log = logging.NewLoggerFromConfig(conf.Logging)

	if conf.Pprof.Enabled {
		l.Log.Info("vega is starting with pprof profile, this is not a recommended setting for production")
		l.pproffhandlr, err = pprof.New(l.Log, conf.Pprof)
		if err != nil {
			return
		}
		l.cfgwatchr.OnConfigUpdate(
			func(cfg config.Config) { l.pproffhandlr.ReloadConf(cfg.Pprof) },
		)
	}

	l.Log.Info("Starting Vega",
		logging.String("config-path", l.configPath),
		logging.String("version", l.Version),
		logging.String("version-hash", l.VersionHash))

	// this doesn't fail
	l.timeService = vegatime.New(l.conf.Time)

	// Set ulimits
	if err = l.SetUlimits(); err != nil {
		l.Log.Warn("Unable to set ulimits",
			logging.Error(err))
	} else {
		l.Log.Debug("Set ulimits",
			logging.Uint64("nofile", l.conf.UlimitNOFile))
	}

	l.stats = stats.New(l.Log, l.conf.Stats, l.Version, l.VersionHash)

	// set up storage, this should be persistent
	if err := l.setupStorages(); err != nil {
		return err
	}
	l.setupSubscibers()

	if !l.conf.StoresEnabled {
		l.Log.Info("node setted up without badger store support")
	} else {
		l.Log.Info("node setted up with badger store support")
	}

	// instantiate the ETHClient
	ethClient, err := ethclient.Dial(l.conf.NodeWallet.ETH.Address)
	if err != nil {
		return err
	}

	// nodewallet
	if l.nodeWallet, err = nodewallet.New(l.Log, l.conf.NodeWallet, l.nodeWalletPassphrase, ethClient, l.configPath); err != nil {
		return err
	}

	l.ethClient = ethClient

	return l.nodeWallet.Verify()
}

func (l *NodeCommand) setupSubscibers() {
	l.transferSub = subscribers.NewTransferResponse(l.ctx, l.transferResponseStore, l.Log, true)
	l.marketEventSub = subscribers.NewMarketEvent(l.ctx, l.conf.Subscribers, l.Log, false)
	l.orderSub = subscribers.NewOrderEvent(l.ctx, l.conf.Subscribers, l.Log, l.orderStore, true)
	l.accountSub = subscribers.NewAccountSub(l.ctx, l.accounts, l.Log, true)
	l.partySub = subscribers.NewPartySub(l.ctx, l.partyStore, l.Log, true)
	l.tradeSub = subscribers.NewTradeSub(l.ctx, l.tradeStore, l.Log, true)
	l.marginLevelSub = subscribers.NewMarginLevelSub(l.ctx, l.riskStore, l.Log, true)
	l.governanceSub = subscribers.NewGovernanceDataSub(l.ctx, l.Log, true)
	l.voteSub = subscribers.NewVoteSub(l.ctx, false, true, l.Log)
	l.marketDataSub = subscribers.NewMarketDataSub(l.ctx, l.marketDataStore, l.Log, true)
	l.newMarketSub = subscribers.NewMarketSub(l.ctx, l.marketStore, l.Log, true)
	l.marketUpdatedSub = subscribers.NewMarketUpdatedSub(l.ctx, l.marketStore, l.Log, true)
	l.candleSub = subscribers.NewCandleSub(l.ctx, l.candleStore, l.Log, true)
	l.marketDepthSub = subscribers.NewMarketDepthBuilder(l.ctx, l.Log, true)
	l.riskFactorSub = subscribers.NewRiskFactorSub(l.ctx, l.riskStore, l.Log, true)
}

func (l *NodeCommand) setupStorages() (err error) {
	l.marketDataStore = storage.NewMarketData(l.Log, l.conf.Storage)
	l.riskStore = storage.NewRisks(l.Log, l.conf.Storage)

	// always enabled market,parties etc stores as they are in memory or boths use them
	if l.marketStore, err = storage.NewMarkets(l.Log, l.conf.Storage, l.cancel); err != nil {
		return
	}

	if l.partyStore, err = storage.NewParties(l.conf.Storage); err != nil {
		return
	}
	if l.transferResponseStore, err = storage.NewTransferResponses(l.Log, l.conf.Storage); err != nil {
		return
	}

	// if stores are not enabled, initialise the noop stores and do nothing else
	if !l.conf.StoresEnabled {
		l.orderStore = storage.NewNoopOrders(l.Log, l.conf.Storage)
		l.tradeStore = storage.NewNoopTrades(l.Log, l.conf.Storage)
		l.accounts = storage.NewNoopAccounts(l.Log, l.conf.Storage)
		l.candleStore = storage.NewNoopCandles(l.Log, l.conf.Storage)
		return
	}

	if l.candleStore, err = storage.NewCandles(l.Log, l.conf.Storage, l.cancel); err != nil {
		return
	}

	if l.orderStore, err = storage.NewOrders(l.Log, l.conf.Storage, l.cancel); err != nil {
		return
	}
	if l.tradeStore, err = storage.NewTrades(l.Log, l.conf.Storage, l.cancel); err != nil {
		return
	}
	if l.accounts, err = storage.NewAccounts(l.Log, l.conf.Storage, l.cancel); err != nil {
		return
	}

	l.cfgwatchr.OnConfigUpdate(
		func(cfg config.Config) { l.accounts.ReloadConf(cfg.Storage) },
		func(cfg config.Config) { l.tradeStore.ReloadConf(cfg.Storage) },
		func(cfg config.Config) { l.orderStore.ReloadConf(cfg.Storage) },
		func(cfg config.Config) { l.candleStore.ReloadConf(cfg.Storage) },
		func(cfg config.Config) { l.transferResponseStore.ReloadConf(cfg.Storage) },
		func(cfg config.Config) { l.partyStore.ReloadConf(cfg.Storage) },
		func(cfg config.Config) { l.riskStore.ReloadConf(cfg.Storage) },
		func(cfg config.Config) { l.marketDataStore.ReloadConf(cfg.Storage) },
		func(cfg config.Config) { l.marketStore.ReloadConf(cfg.Storage) },
		func(cfg config.Config) { l.stats.ReloadConf(cfg.Stats) },
	)

	return
}

// UponGenesis loads all asset from genesis state
func (l *NodeCommand) UponGenesis(ctx context.Context, rawstate []byte) (err error) {
	l.Log.Debug("Entering node.NodeCommand.UponGenesis")
	defer func() {
		if err != nil {
			l.Log.Debug("Failure in node.NodeCommand.UponGenesis", logging.Error(err))
		} else {
			l.Log.Debug("Leaving node.NodeCommand.UponGenesis without error")
		}
	}()

	state, err := assets.LoadGenesisState(rawstate)
	if err != nil {
		return err
	}
	if state == nil {
		return nil
	}

	for k, v := range state {
		err := l.loadAsset(k, v)
		if err != nil {
			return err
		}
	}

	return nil
}

func (l *NodeCommand) loadAsset(id string, v *proto.AssetDetails) error {
	aid, err := l.assets.NewAsset(id, types.AssetDetailsFromProto(v))
	if err != nil {
		return fmt.Errorf("error instanciating asset %v", err)
	}

	asset, err := l.assets.Get(aid)
	if err != nil {
		return fmt.Errorf("unable to get asset %v", err)
	}

	// just a simple backoff here
	err = backoff.Retry(
		func() error {
			err := asset.Validate()
			if !asset.IsValid() {
				return err
			}
			return nil
		},
		backoff.WithMaxRetries(backoff.NewExponentialBackOff(), 5),
	)
	if err != nil {
		return fmt.Errorf("unable to instantiate new asset err=%v, asset-source=%s", err, v.String())
	}
	if err := l.assets.Enable(aid); err != nil {
		l.Log.Error("invalid genesis asset",
			logging.String("asset-details", v.String()),
			logging.Error(err))
		return fmt.Errorf("unable to enable asset: %v", err)
	}

	assetD := asset.Type()
	if err := l.collateral.EnableAsset(context.Background(), *assetD); err != nil {
		return fmt.Errorf("unable to enable asset in collateral: %v", err)
	}

	l.Log.Info("new asset added successfully",
		logging.String("asset", asset.String()))

	return nil
}

func (l *NodeCommand) startABCI(ctx context.Context, commander *nodewallet.Commander) (*processor.App, error) {
	app := processor.NewApp(
		l.Log,
		l.conf.Processor,
		l.cancel,
		l.assets,
		l.banking,
		l.broker,
		l.witness,
		l.evtfwd,
		l.executionEngine,
		commander,
		l.genesisHandler,
		l.governance,
		l.notary,
		l.stats.Blockchain,
		l.timeService,
		l.epochService,
		l.topology,
		l.netParams,
		&processor.Oracle{
			Engine:   l.oracle,
			Adaptors: l.oracleAdaptors,
		},
		l.delegation,
		l.limits,
		l.stakeVerifier,
		l.checkpoint,
	)

	var abciApp tmtypes.Application
	tmCfg := l.conf.Blockchain.Tendermint
	if path := tmCfg.ABCIRecordDir; path != "" {
		rec, err := recorder.NewRecord(path, afero.NewOsFs())
		if err != nil {
			return nil, err
		}

		// closer
		go func() {
			<-ctx.Done()
			rec.Stop()
		}()

		abciApp = recorder.NewApp(app.Abci(), rec)
	} else {
		abciApp = app.Abci()
	}

	srv := abci.NewServer(l.Log, l.conf.Blockchain, abciApp)
	if err := srv.Start(); err != nil {
		return nil, err
	}
	l.abciServer = srv

	if path := tmCfg.ABCIReplayFile; path != "" {
		rec, err := recorder.NewReplay(path, afero.NewOsFs())
		if err != nil {
			return nil, err
		}

		// closer
		go func() {
			<-ctx.Done()
			rec.Stop()
		}()

		go func() {
			if err := rec.Replay(abciApp); err != nil {
				log.Fatalf("replay: %v", err)
			}
		}()
	}

	abciClt, err := abci.NewClient(l.conf.Blockchain.Tendermint.ClientAddr)
	if err != nil {
		return nil, err
	}
	l.blockchainClient = blockchain.NewClient(abciClt)
	commander.SetChain(l.blockchainClient)

	return app, nil
}

// we've already set everything up WRT arguments etc... just bootstrap the node
func (l *NodeCommand) preRun(_ []string) (err error) {
	// ensure that context is cancelled if we return an error here
	defer func() {
		if err != nil {
			l.cancel()
		}
	}()

	// plugins
	l.settlePlugin = plugins.NewPositions(l.ctx)
	l.notaryPlugin = plugins.NewNotary(l.ctx)
	l.assetPlugin = plugins.NewAsset(l.ctx)
	l.withdrawalPlugin = plugins.NewWithdrawal(l.ctx)
	l.depositPlugin = plugins.NewDeposit(l.ctx)
	l.netParamsService = netparams.NewService(l.ctx)
	l.liquidityService = liquidity.NewService(l.ctx, l.Log, l.conf.Liquidity)
	l.oracleService = oracles.NewService(l.ctx)

	l.genesisHandler = genesis.New(l.Log, l.conf.Genesis)
	l.genesisHandler.OnGenesisTimeLoaded(l.timeService.SetTimeNow)

	l.broker, err = broker.New(l.ctx, l.Log, l.conf.Broker)
	if err != nil {
		log.Error("unable to initialise broker", logging.Error(err))
		return err
	}

	l.broker.SubscribeBatch(
		l.marketEventSub, l.transferSub, l.orderSub, l.accountSub,
		l.partySub, l.tradeSub, l.marginLevelSub, l.governanceSub,
		l.voteSub, l.marketDataSub, l.notaryPlugin, l.settlePlugin,
		l.newMarketSub, l.assetPlugin, l.candleSub, l.withdrawalPlugin,
		l.depositPlugin, l.marketDepthSub, l.riskFactorSub, l.netParamsService,
		l.liquidityService, l.marketUpdatedSub, l.oracleService)

	now := l.timeService.GetTimeNow()

	l.assets = assets.New(l.Log, l.conf.Assets, l.nodeWallet, l.timeService)
	l.collateral = collateral.New(l.Log, l.conf.Collateral, l.broker, now)
	l.oracle = oracles.NewEngine(l.Log, l.conf.Oracles, now, l.broker, l.timeService)
	l.timeService.NotifyOnTick(l.oracle.UpdateCurrentTime)
	l.oracleAdaptors = oracleAdaptors.New()

	// instantiate the execution engine
	l.executionEngine = execution.NewEngine(
		l.Log,
		l.conf.Execution,
		l.timeService,
		l.collateral,
		l.oracle,
		l.broker,
	)
	// we cannot pass the Chain dependency here (that's set by the blockchain)
	wal, _ := l.nodeWallet.Get(nodewallet.Vega)
	commander, err := nodewallet.NewCommander(l.Log, nil, wal, l.stats)
	if err != nil {
		return err
	}

	l.limits = limits.New(l.Log, l.conf.Limits)
	l.timeService.NotifyOnTick(l.limits.OnTick)

	l.topology = validators.NewTopology(l.Log, l.conf.Validators, wal, l.broker)

	l.witness = validators.NewWitness(l.Log, l.conf.Validators, l.topology, commander, l.timeService)

	l.netParams = netparams.New(l.Log, l.conf.NetworkParameters, l.broker)

	l.governance = governance.NewEngine(l.Log, l.conf.Governance, l.collateral, l.broker, l.assets, l.witness, l.netParams, now)

	l.stakingAccounts, l.stakeVerifier = staking.New(
		l.Log, l.conf.Staking, l.broker, l.timeService, l.witness, l.ethClient, l.netParams,
	)

	// checkpoint engine
	l.checkpoint, err = checkpoint.New(l.Log, l.conf.Checkpoint, l.assets, l.collateral, l.governance, l.netParams)
	if err != nil {
		panic(err)
	}

	l.genesisHandler.OnGenesisAppStateLoaded(
		// be sure to keep this in order.
		// the node upon genesis will load all asset first in the node
		// state. This is important to happened first as we will load the
		// asset which will be considered as the governance token.
		l.UponGenesis,
		// This needs to happen always after, as it defined the network
		// parameters, one of them is  the Governance Token asset ID.
		// which if not loaded in the previous state, then will make the node
		// panic at startup.
		l.netParams.UponGenesis,
		l.topology.LoadValidatorsOnGenesis,
		l.limits.UponGenesis,
		l.checkpoint.UponGenesis,
	)

	l.notary = notary.New(l.Log, l.conf.Notary, l.topology, l.broker, commander)
	l.evtfwd = evtforward.New(l.Log, l.conf.EvtForward, commander, l.timeService, l.topology)
	l.banking = banking.New(l.Log, l.conf.Banking, l.collateral, l.witness, l.timeService, l.assets, l.notary, l.broker)

	// now instantiate the blockchain layer
	if l.app, err = l.startABCI(l.ctx, commander); err != nil {
		return err
	}

	// start services
	if l.candleService, err = candles.NewService(l.Log, l.conf.Candles, l.candleStore); err != nil {
		return
	}

	if l.orderService, err = orders.NewService(l.Log, l.conf.Orders, l.orderStore, l.timeService); err != nil {
		return
	}

	if l.tradeService, err = trades.NewService(l.Log, l.conf.Trades, l.tradeStore, l.settlePlugin); err != nil {
		return
	}
	if l.marketService, err = markets.NewService(l.Log, l.conf.Markets, l.marketStore, l.orderStore, l.marketDataStore, l.marketDepthSub); err != nil {
		return
	}
	l.riskService = risk.NewService(l.Log, l.conf.Risk, l.riskStore, l.marketStore, l.marketDataStore)
	l.governanceService = governance.NewService(l.Log, l.conf.Governance, l.broker, l.governanceSub, l.voteSub, l.netParams)

	// last assignment to err, no need to check here, if something went wrong, we'll know about it
	l.feeService = fee.NewService(l.Log, l.conf.Execution.Fee, l.marketStore, l.marketDataStore)
	l.partyService, err = parties.NewService(l.Log, l.conf.Parties, l.partyStore)
	l.accountsService = accounts.NewService(l.Log, l.conf.Accounts, l.accounts)
	l.transfersService = transfers.NewService(l.Log, l.conf.Transfers, l.transferResponseStore)
	l.notaryService = notary.NewService(l.Log, l.conf.Notary, l.notaryPlugin)
	l.assetService = assets.NewService(l.Log, l.conf.Assets, l.assetPlugin)
	l.eventService = subscribers.NewService(l.broker)
	l.epochService = epochtime.NewService(l.Log, l.conf.Epoch, l.timeService, l.broker)

	// TODO replace with actual implementation
	stakingAccount := delegation.NewDummyStakingAccount(l.collateral)
	l.netParams.Watch(netparams.WatchParam{
		Param:   netparams.GovernanceVoteAsset,
		Watcher: stakingAccount.GovAssetUpdated,
	})

	l.delegation = delegation.New(l.Log, delegation.NewDefaultConfig(), l.broker, l.topology, stakingAccount, l.epochService)
	l.netParams.Watch(netparams.WatchParam{
		Param:   netparams.DelegationMinAmount,
		Watcher: l.delegation.OnMinAmountChanged,
	})

	// setup rewards engine
	l.rewards = rewards.New(l.Log, l.conf.Rewards, l.broker, l.delegation, l.epochService, l.collateral, l.timeService)

	// setup config reloads for all engines / services /etc
	l.setupConfigWatchers()
	l.timeService.NotifyOnTick(l.cfgwatchr.OnTimeUpdate)

	// setup some network parameters runtime validations
	// and network parameters updates dispatches
	return l.setupNetParameters()
}

func (l *NodeCommand) setupNetParameters() error {
	// now we are going to setup some network parameters which can be done
	// through runtime checks
	// e.g: changing the governance asset require the Assets and Collateral engines, so we can ensure any changes there are made for a valid asset
	if err := l.netParams.AddRules(
		netparams.ParamStringRules(
			netparams.GovernanceVoteAsset,
			checks.GovernanceAssetUpdate(l.Log, l.assets, l.collateral),
		),
	); err != nil {
		return err
	}

	// now add some watcher for our netparams
	return l.netParams.Watch(
		netparams.WatchParam{
			Param:   netparams.GovernanceVoteAsset,
			Watcher: dispatch.GovernanceAssetUpdate(l.Log, l.assets),
		},
		netparams.WatchParam{
			Param:   netparams.MarketMarginScalingFactors,
			Watcher: l.executionEngine.OnMarketMarginScalingFactorsUpdate,
		},
		netparams.WatchParam{
			Param:   netparams.MarketFeeFactorsMakerFee,
			Watcher: l.executionEngine.OnMarketFeeFactorsMakerFeeUpdate,
		},
		netparams.WatchParam{
			Param:   netparams.MarketFeeFactorsInfrastructureFee,
			Watcher: l.executionEngine.OnMarketFeeFactorsInfrastructureFeeUpdate,
		},
		netparams.WatchParam{
			Param:   netparams.MarketLiquidityStakeToCCYSiskas,
			Watcher: l.executionEngine.OnSuppliedStakeToObligationFactorUpdate,
		},
		netparams.WatchParam{
			Param:   netparams.MarketValueWindowLength,
			Watcher: l.executionEngine.OnMarketValueWindowLengthUpdate,
		},
		netparams.WatchParam{
			Param:   netparams.MarketTargetStakeScalingFactor,
			Watcher: l.executionEngine.OnMarketTargetStakeScalingFactorUpdate,
		},
		netparams.WatchParam{
			Param:   netparams.MarketTargetStakeTimeWindow,
			Watcher: l.executionEngine.OnMarketTargetStakeTimeWindowUpdate,
		},
		netparams.WatchParam{
			Param:   netparams.BlockchainsEthereumConfig,
			Watcher: l.nodeWallet.OnEthereumConfigUpdate,
		},
		netparams.WatchParam{
			Param:   netparams.MarketLiquidityProvidersFeeDistribitionTimeStep,
			Watcher: l.executionEngine.OnMarketLiquidityProvidersFeeDistributionTimeStep,
		},
		netparams.WatchParam{
			Param:   netparams.MarketLiquidityProvisionShapesMaxSize,
			Watcher: l.executionEngine.OnMarketLiquidityProvisionShapesMaxSizeUpdate,
		},
		netparams.WatchParam{
			Param:   netparams.MarketLiquidityMaximumLiquidityFeeFactorLevel,
			Watcher: l.executionEngine.OnMarketLiquidityMaximumLiquidityFeeFactorLevelUpdate,
		},
		netparams.WatchParam{
			Param:   netparams.MarketLiquidityBondPenaltyParameter,
			Watcher: l.executionEngine.OnMarketLiquidityBondPenaltyUpdate,
		},
		netparams.WatchParam{
			Param:   netparams.MarketLiquidityTargetStakeTriggeringRatio,
			Watcher: l.executionEngine.OnMarketLiquidityTargetStakeTriggeringRatio,
		},
		netparams.WatchParam{
			Param:   netparams.MarketAuctionMinimumDuration,
			Watcher: l.executionEngine.OnMarketAuctionMinimumDurationUpdate,
		},
		netparams.WatchParam{
			Param:   netparams.MarketProbabilityOfTradingTauScaling,
			Watcher: l.executionEngine.OnMarketProbabilityOfTradingTauScalingUpdate,
		},
		netparams.WatchParam{
			Param:   netparams.MarketMinProbabilityOfTradingForLPOrders,
			Watcher: l.executionEngine.OnMarketMinProbabilityOfTradingForLPOrdersUpdate,
		},
		netparams.WatchParam{
			Param:   netparams.ValidatorsEpochLength,
			Watcher: l.epochService.OnEpochLengthUpdate,
		},
		netparams.WatchParam{
			Param:   netparams.GovernanceVoteAsset,
			Watcher: l.rewards.UpdateAssetForStakingAndDelegationRewardScheme,
		},
		netparams.WatchParam{
			Param:   netparams.StakingAndDelegationRewardPayoutFraction,
			Watcher: l.rewards.UpdatePayoutFractionForStakingRewardScheme,
		},
		netparams.WatchParam{
			Param:   netparams.StakingAndDelegationRewardPayoutDelay,
			Watcher: l.rewards.UpdatePayoutDelayForStakingRewardScheme,
		},
		netparams.WatchParam{
			Param:   netparams.StakingAndDelegationRewardMaxPayoutPerParticipant,
			Watcher: l.rewards.UpdateMaxPayoutPerParticipantForStakingRewardScheme,
		},
		netparams.WatchParam{
			Param:   netparams.StakingAndDelegationRewardDelegatorShare,
			Watcher: l.rewards.UpdateDelegatorShareForStakingRewardScheme,
		},
		netparams.WatchParam{
			Param:   netparams.StakingAndDelegationRewardMinimumValidatorStake,
			Watcher: l.rewards.UpdateMinimumValidatorStakeForStakingRewardScheme,
		},
		netparams.WatchParam{
			Param:   netparams.ValidatorsVoteRequired,
			Watcher: l.witness.OnDefaultValidatorsVoteRequiredUpdate,
		},
		netparams.WatchParam{
			Param:   netparams.ValidatorsVoteRequired,
			Watcher: l.notary.OnDefaultValidatorsVoteRequiredUpdate,
		},
		netparams.WatchParam{
			Param:   netparams.NetworkCheckpointTimeElapsedBetweenCheckpoints,
			Watcher: l.checkpoint.OnTimeElapsedUpdate,
		},
	)
}

func (l *NodeCommand) setupConfigWatchers() {
	l.cfgwatchr.OnConfigUpdate(
		func(cfg config.Config) { l.executionEngine.ReloadConf(cfg.Execution) },
		func(cfg config.Config) { l.notary.ReloadConf(cfg.Notary) },
		func(cfg config.Config) { l.evtfwd.ReloadConf(cfg.EvtForward) },
		func(cfg config.Config) { l.abciServer.ReloadConf(cfg.Blockchain) },
		func(cfg config.Config) { l.topology.ReloadConf(cfg.Validators) },
		func(cfg config.Config) { l.witness.ReloadConf(cfg.Validators) },
		func(cfg config.Config) { l.assets.ReloadConf(cfg.Assets) },
		func(cfg config.Config) { l.banking.ReloadConf(cfg.Banking) },
		func(cfg config.Config) { l.governance.ReloadConf(cfg.Governance) },
		func(cfg config.Config) { l.nodeWallet.ReloadConf(cfg.NodeWallet) },
		func(cfg config.Config) { l.app.ReloadConf(cfg.Processor) },

		// services
		func(cfg config.Config) { l.candleService.ReloadConf(cfg.Candles) },
		func(cfg config.Config) { l.orderService.ReloadConf(cfg.Orders) },
		func(cfg config.Config) { l.liquidityService.ReloadConf(cfg.Liquidity) },
		func(cfg config.Config) { l.tradeService.ReloadConf(cfg.Trades) },
		func(cfg config.Config) { l.marketService.ReloadConf(cfg.Markets) },
		func(cfg config.Config) { l.riskService.ReloadConf(cfg.Risk) },
		func(cfg config.Config) { l.governanceService.ReloadConf(cfg.Governance) },
		func(cfg config.Config) { l.assetService.ReloadConf(cfg.Assets) },
		func(cfg config.Config) { l.notaryService.ReloadConf(cfg.Notary) },
		func(cfg config.Config) { l.transfersService.ReloadConf(cfg.Transfers) },
		func(cfg config.Config) { l.accountsService.ReloadConf(cfg.Accounts) },
		func(cfg config.Config) { l.partyService.ReloadConf(cfg.Parties) },
		func(cfg config.Config) { l.feeService.ReloadConf(cfg.Execution.Fee) },
	)
}
