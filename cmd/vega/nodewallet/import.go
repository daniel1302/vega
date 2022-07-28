// Copyright (c) 2022 Gobalsky Labs Limited
//
// Use of this software is governed by the Business Source License included
// in the LICENSE file and at https://www.mariadb.com/bsl11.
//
// Change Date: 18 months from the later of the date of the first publicly
// available Distribution of this version of the repository, and 25 June 2022.
//
// On the date above, in accordance with the Business Source License, use
// of this software will be governed by version 3 or later of the GNU General
// Public License.

package nodewallet

import (
	"encoding/json"
	"errors"
	"fmt"

	vgfs "code.vegaprotocol.io/shared/libs/fs"
	vgjson "code.vegaprotocol.io/shared/libs/json"
	"code.vegaprotocol.io/shared/paths"
	"code.vegaprotocol.io/vega/core/config"
	vgfmt "code.vegaprotocol.io/vega/core/libs/fmt"
	"code.vegaprotocol.io/vega/core/nodewallets"
	"code.vegaprotocol.io/vega/logging"

	"github.com/jessevdk/go-flags"
	tmconfig "github.com/tendermint/tendermint/config"
)

var (
	ErrOneOfTendermintFlagIsRequired       = errors.New("one of --tendermint-home or --tendermint-pubkey flag is required")
	ErrTendermintFlagsAreMutuallyExclusive = errors.New("--tendermint-home and --tendermint-pubkey are mutually exclusive")
)

type importCmd struct {
	config.OutputFlag

	Config nodewallets.Config

	WalletPassphrase   config.Passphrase   `long:"wallet-passphrase-file"`
	ClefAccountAddress config.PromptString `long:"clef-account-address" description:"The Ethereum account address to be imported by Vega from Clef. In hex."`

	Chain      string              `short:"c" long:"chain" required:"true" description:"The chain to be imported" choice:"vega" choice:"ethereum" choice:"tendermint"`
	WalletPath config.PromptString `long:"wallet-path" description:"The path to the wallet file to import"`
	Force      bool                `long:"force" description:"Should the command re-write an existing nodewallet file if it exists"`

	// tendermint flags
	TendermintPubkey string `long:"tendermint-pubkey" description:"The tendermint pubkey of the tendermint validator node"`
	TendermintHome   string `long:"tendermint-home" description:"The tendermint home from which to look the for the pubkey"`
}

func (opts *importCmd) Execute(_ []string) error {
	output, err := opts.GetOutput()
	if err != nil {
		return err
	}

	log := logging.NewLoggerFromConfig(logging.NewDefaultConfig())
	defer log.AtExit()

	registryPass, err := rootCmd.PassphraseFile.Get("node wallet", false)
	if err != nil {
		return err
	}

	vegaPaths := paths.New(rootCmd.VegaHome)

	_, conf, err := config.EnsureNodeConfig(vegaPaths)
	if err != nil {
		return err
	}

	opts.Config = conf.NodeWallet

	if _, err := flags.NewParser(opts, flags.Default|flags.IgnoreUnknown).Parse(); err != nil {
		return err
	}

	clefEnabled := opts.Config.ETH.ClefAddress != ""

	var walletPass, walletPath string
	if opts.Chain == vegaChain || (opts.Chain == ethereumChain && !clefEnabled) {
		walletPass, err = opts.WalletPassphrase.Get("blockchain wallet", false)
		if err != nil {
			return err
		}
		walletPath, err = opts.WalletPath.Get("wallet path", "wallet-path")
		if err != nil {
			return err
		}
	}

	var data map[string]string
	switch opts.Chain {
	case ethereumChain:
		var accountAddress string
		if clefEnabled {
			accountAddress, err = opts.ClefAccountAddress.Get("Clef account address", "clef-account-address")
			if err != nil {
				return err
			}
		}

		data, err = nodewallets.ImportEthereumWallet(
			opts.Config.ETH,
			vegaPaths,
			registryPass,
			walletPass,
			accountAddress,
			walletPath,
			opts.Force,
		)
		if err != nil {
			return fmt.Errorf("couldn't import Ethereum node wallet: %w", err)
		}
	case vegaChain:
		data, err = nodewallets.ImportVegaWallet(vegaPaths, registryPass, walletPass, walletPath, opts.Force)
		if err != nil {
			return fmt.Errorf("couldn't import Vega node wallet: %w", err)
		}
	case tendermintChain:
		if len(opts.TendermintHome) == 0 && len(opts.TendermintPubkey) == 0 {
			return ErrOneOfTendermintFlagIsRequired
		}
		if len(opts.TendermintHome) > 0 && len(opts.TendermintPubkey) > 0 {
			return ErrTendermintFlagsAreMutuallyExclusive
		}

		tendermintPubkey := opts.TendermintPubkey
		if len(opts.TendermintHome) > 0 {
			tendermintPubkey, err = getLocalTendermintPubkey(opts.TendermintHome)
			if err != nil {
				return fmt.Errorf("couldn't retrieve tendermint public key: %w", err)
			}
		}

		data, err = nodewallets.ImportTendermintPubkey(vegaPaths, registryPass, tendermintPubkey, opts.Force)
		if err != nil {
			return fmt.Errorf("couldn't import Tendermint pubkey: %w", err)
		}
	}

	if output.IsHuman() {
		fmt.Println(green("import successful:"))
		vgfmt.PrettyPrint(data)
	} else if output.IsJSON() {
		if err := vgjson.Print(data); err != nil {
			return err
		}
	}

	return nil
}

func getLocalTendermintPubkey(tendermintHome string) (string, error) {
	tmConfig := tmconfig.DefaultConfig()
	tmConfig.SetRoot(tendermintHome)
	genesisFilePath := tmConfig.PrivValidator.KeyFile()

	data, err := vgfs.ReadFile(genesisFilePath)
	if err != nil {
		return "", err
	}

	privValidatorKey := struct {
		PubKey struct {
			Value string `json:"value"`
		} `json:"pub_key"`
	}{}

	if err = json.Unmarshal(data, &privValidatorKey); err != nil {
		return "", fmt.Errorf("could not read priv_validator_key.json: %w", err)
	}

	return privValidatorKey.PubKey.Value, nil
}
