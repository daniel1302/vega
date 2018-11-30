package blockchain

import (
	"encoding/binary"
	"vega/core"
	"vega/log"
	"vega/msg"
	"github.com/tendermint/tendermint/abci/example/code"
	"github.com/tendermint/tendermint/abci/types"
	"github.com/golang/protobuf/proto"
)

type Blockchain struct {
	types.BaseApplication
	vega *core.Vega
	previousTimestamp int64
}

func NewBlockchain(vegaApp *core.Vega) *Blockchain {
	return &Blockchain{vega: vegaApp}
}

func (app *Blockchain) BeginBlock(beginBlock types.RequestBeginBlock) types.ResponseBeginBlock {
	//log.Debugf(fmt.Sprintf("Begin block time report (%d txs):", beginBlock.Header.NumTxs))
	//log.Debugf("------------------------")
	//log.Debugf(fmt.Sprintf("Gossip time: %v", beginBlock.Header.Time))
	//log.Debugf(fmt.Sprintf("Unix epoch+nano: %d", beginBlock.Header.Time.UnixNano()))
	//log.Debugf("------------------------")

	// We need to cache the last timestamp so we can distribute trades
	// in a block evenly between last timestamp and current timestamp
	if app.vega.State.Timestamp > 0 {
		app.previousTimestamp = app.vega.State.Timestamp
	}

	// Store the timestamp info that we receive from the block chain provider (Tendermint)
	app.vega.State.Datetime = beginBlock.Header.Time
	app.vega.State.Timestamp = beginBlock.Header.Time.UnixNano()

	// Ensure we always set app.previousTimestamp it'll be 0 on the first block
	if app.previousTimestamp < 1 {
		app.previousTimestamp = app.vega.State.Timestamp
	}

	app.vega.StartCandleBuffer()

	return types.ResponseBeginBlock{}
}

//func (app *Blockchain) EndBlock(endBlock types.RequestEndBlock) types.ResponseEndBlock {
//	//fmt.Println(fmt.Sprintf("%v", endBlock))
//	return types.ResponseEndBlock{}
//}

// Mempool Connection
//
// A transaction is received by a validator from a client into its own
// (*one node*) mempool. We need to check whether we consider it
// "legal" (validly formatted, containing non-crazy data from a business
// perspective). If so, send it through to the consensus round.
//
// From the Tendermint docs:
//
// [The mempool connection is used only for CheckTx requests. Transactions are
// run using CheckTx in the same order they were received by the validator. If
// the CheckTx returns OK, the transaction is kept in memory and relayed to
// other peers in the same order it was received. Otherwise, it is discarded.
//
// CheckTx requests run concurrently with block processing; so they should run
// against a copy of the main application state which is reset after every block.
// This copy is necessary to track transitions made by a sequence of CheckTx
// requests before they are included in a block. When a block is committed,
// the application must ensure to reset the mempool state to the latest
// committed state. Tendermint Core will then filter through all transactions
// in the mempool, removing any that were included in the block, and re-run
// the rest using CheckTx against the post-Commit mempool state]
//
// FIXME: For the moment, just let everything through.
func (app *Blockchain) CheckTx(tx []byte) types.ResponseCheckTx {
	//log.Debugf("CheckTx: %s", string(tx))
	return types.ResponseCheckTx{Code: code.CodeTypeOK}
}

var tx_averages []int
var ob_averages []int
var current_ob int
var current_tb int
var tx_per_block uint64

// Consensus Connection
// Step 1: DeliverTx
//
// A transaction has been accepted by more than 2/3 of
// validator nodes. At this step, we can execute our business logic (or,
// in Ethereum terms, this is where the smart contract code lives).
//
// Every honest validator node will run state changes according to what
// happens in this function.
//
// From the Tendermint docs:
//
// [DeliverTx is the workhorse of the blockchain. Tendermint sends the DeliverTx
// requests asynchronously but in order, and relies on the underlying socket
// protocol (ie. TCP) to ensure they are received by the app in order. They
// have already been ordered in the global consensus by the Tendermint protocol.
//
// DeliverTx returns a abci.Result, which includes a Code, Data, and Log. The
// code may be non-zero (non-OK), meaning the corresponding transaction should
// have been rejected by the mempool, but may have been included in a block by
// a Byzantine proposer.
//
// The block header will be updated (TODO) to include some commitment to the
// results of DeliverTx, be it a bitarray of non-OK transactions, or a merkle
// root of the data returned by the DeliverTx requests, or both]
func (app *Blockchain) DeliverTx(tx []byte) types.ResponseDeliverTx {
	txLength := len(tx)
	//log.Debugf("DeliverTx: %s [%v]", string(tx), txLength)
	tx_per_block++

	if app.vega.Statistics.Status == msg.AppStatus_CHAIN_NOT_FOUND {
		app.vega.Statistics.Status = msg.AppStatus_CHAIN_REPLAYING
	}

	if tx_averages == nil {
		tx_averages = make([]int, 0)
	}
	tx_averages = append(tx_averages, txLength)

	// Decode payload and command
	value, cmd, err := VegaTxDecode(tx)
	if err != nil {
		log.Errorf("Invalid tx when decoding DeliverTx payload")
		return types.ResponseDeliverTx{Code: code.CodeTypeEncodingError}
	}

	// Process known command types
	switch cmd {
		case CreateOrderCommand:

			// deserialize proto msg to struct
			order := msg.OrderPool.Get().(*msg.Order)
			e := proto.Unmarshal(value, order)
			if e != nil {
				log.Errorf("Error: Decoding order to proto: ", e.Error())
				return types.ResponseDeliverTx{Code: code.CodeTypeEncodingError}
			}

			log.Debugf("ABCI received a CREATE ORDER command after consensus")

			// Submit the create new order request to the Vega trading core
			confirmationMessage, errorMessage := app.vega.SubmitOrder(order)
			if confirmationMessage != nil {
				log.Infof("ABCI order confirmation message:")
				log.Infof("- aggressive order: %+v", confirmationMessage.Order)
				log.Debugf("- trades: %+v", confirmationMessage.Trades)
				log.Infof("- passive orders affected: %+v", confirmationMessage.PassiveOrdersAffected)

				current_tb += len(confirmationMessage.Trades)
			}
			if errorMessage != msg.OrderError_NONE {
				log.Infof("ABCI order error message (create):")
				log.Infof("- error: %+v", errorMessage.String())
			}
			current_ob++

		case CancelOrderCommand:

			// deserialize proto msg to struct
			order := msg.OrderPool.Get().(*msg.Order)
			e := proto.Unmarshal(value, order)
			if e != nil {
				log.Errorf("Error: Decoding order to proto: ", e.Error())
				return types.ResponseDeliverTx{Code: code.CodeTypeEncodingError}
			}

			log.Debugf("ABCI received a CANCEL ORDER command after consensus")

			// Submit the cancel new order request to the Vega trading core
			cancellationMessage, errorMessage := app.vega.CancelOrder(order)
			if cancellationMessage != nil {
				log.Infof("ABCI order cancellation message:")
				log.Infof("- cancelled order: %+v", cancellationMessage.Order)
			}
			if errorMessage != msg.OrderError_NONE {
				log.Infof("ABCI order error message (cancel):")
				log.Infof("- error: %+v", errorMessage.String())
			}

		case AmendmentOrderCommand:

			// deserialize proto msg to struct
			amendment := &msg.Amendment{}
			e := proto.Unmarshal(value, amendment)
			if e != nil {
				log.Errorf("Error: Decoding order to proto: ", e.Error())
				return types.ResponseDeliverTx{Code: code.CodeTypeEncodingError}
			}

			log.Debugf("ABCI received a Amendment command after consensus")

			// Submit the Amendment new order request to the Vega trading core
			confirmationMessage, errorMessage := app.vega.AmendOrder(amendment)
			if confirmationMessage != nil {
				log.Debugf("ABCI reports it received an order amendment message from vega:\n")
				log.Debugf("- cancelled order: %+v\n", confirmationMessage.Order)
			}
			if errorMessage != msg.OrderError_NONE {
				log.Debugf("ABCI reports it received an order error message from vega:\n")
				log.Debugf("- error: %+v\n", errorMessage.String())
			}

		default:
			log.Errorf("UNKNOWN command received after consensus: %v", cmd)
	}

	if len(tx_averages) > 0 {
		totaltx := 0
		for _, itx := range tx_averages {
			totaltx += itx
		}
		averageTx := totaltx / len(tx_averages)
		log.Debugf("Stats: Current tx average size = %v bytes", averageTx)
		app.vega.Statistics.AverageTxBytes = uint64(averageTx)

		// MAX sample size for avg calculation is 5000 txs
		if len(tx_averages) == 5000 {
			tx_averages = nil
		}
	}

	app.vega.State.Size += 1
	return types.ResponseDeliverTx{Code: code.CodeTypeOK}
}


// Consensus Connection
// Step 2: Commit the block and persist to disk.
//
// From the Tendermint docs:
//
// [Once all processing of the block is complete, Tendermint sends the Commit
// request and blocks waiting for a response. While the mempool may run
// concurrently with block processing (the BeginBlock, DeliverTxs, and
// EndBlock), it is locked for the Commit request so that its state can be
// safely reset during Commit. This means the app MUST NOT do any blocking
// communication with the mempool (ie. broadcast_tx) during Commit, or there
// will be deadlock. Note also that all remaining transactions in the mempool
// are replayed on the mempool connection (CheckTx) following a commit.
//
// The app should respond to the Commit request with a byte array, which is
// the deterministic state root of the application. It is included in the
// header of the next block. It can be used to provide easily verified
// Merkle-proofs of the state of the application.
//
// It is expected that the app will persist state to disk on Commit.
// The option to have all transactions replayed from some previous block is
// the job of the Handshake.
//
func (app *Blockchain) Commit() types.ResponseCommit {
	app.vega.RemoveExpiringOrdersAtTimestamp(uint64(app.vega.State.Timestamp))

	// Using a memdb - just return the big endian size of the db
	appHash := make([]byte, 8)
	binary.PutVarint(appHash, app.vega.State.Size)
	app.vega.State.AppHash = appHash
	app.vega.State.Height += 1

	app.vega.Commit()
	app.vega.GenerateCandles()

	app.vega.Statistics.OrdersPerSecond = uint64(current_ob)
	app.vega.Statistics.TradesPerSecond = uint64(current_tb)
	app.vega.Statistics.TxPerBlock = tx_per_block
	tx_per_block = 0

	// Calculate total orders per block average
	if current_ob > 0 {
		if ob_averages == nil {
			ob_averages = make([]int, 0)
		}
		ob_averages = append(ob_averages, current_ob)
		if len(ob_averages) > 0 {
			totalob := 0
			for _, itx := range ob_averages {
				totalob += itx
			}
			averageOb := totalob / len(ob_averages)
			log.Debugf("Stats: Current orders/block average = %v", averageOb)
			app.vega.Statistics.AverageOrdersPerBlock = uint64(averageOb)

			// MAX sample size for avg calculation is 5000 blocks
			if len(ob_averages) == 5000 {
				ob_averages = nil
			}
		}
	}

	current_ob = 0
	current_tb = 0

	return types.ResponseCommit{Data: appHash}
}


