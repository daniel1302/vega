package banking_test

import (
	"context"
	"testing"

	"code.vegaprotocol.io/protos/vega"
	"code.vegaprotocol.io/vega/banking"
	"code.vegaprotocol.io/vega/events"
	"code.vegaprotocol.io/vega/types"
	"code.vegaprotocol.io/vega/types/num"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCancelTransfer(t *testing.T) {
	e := getTestEngine(t)
	defer e.ctrl.Finish()

	// let's do a massive fee, easy to test
	e.OnTransferFeeFactorUpdate(context.Background(), num.NewDecimalFromFloat(0.5))
	e.OnEpoch(context.Background(), types.Epoch{Seq: 7, Action: vega.EpochAction_EPOCH_ACTION_START})
	e.OnEpoch(context.Background(), types.Epoch{Seq: 7, Action: vega.EpochAction_EPOCH_ACTION_END})

	var endEpoch13 uint64 = 11
	transferID := "TRANSFERID"
	partyID := "from"
	ctx := context.Background()
	transfer := &types.TransferFunds{
		Kind: types.TransferCommandKindRecurring,
		Recurring: &types.RecurringTransfer{
			TransferBase: &types.TransferBase{
				ID:   transferID,
				From: partyID,

				FromAccountType: types.AccountTypeGeneral,
				To:              "to",
				ToAccountType:   types.AccountTypeGlobalReward,
				Asset:           "eth",
				Amount:          num.NewUint(100),
				Reference:       "someref",
			},
			StartEpoch: 10,
			EndEpoch:   &endEpoch13,
			Factor:     num.MustDecimalFromString("0.9"),
		},
	}

	e.assets.EXPECT().Get(gomock.Any()).Times(1).Return(nil, nil)
	e.broker.EXPECT().Send(gomock.Any()).Times(1)
	assert.NoError(t, e.TransferFunds(ctx, transfer))

	// now we try to cancel an non-existing transfer
	assert.EqualError(t,
		e.CancelTransferFunds(ctx, &types.CancelTransferFunds{TransferID: "NOPE"}),
		banking.ErrRecurringTransferDoesNotExists.Error(),
	)

	// now we try to cancel the right transfer, but with the wrong party
	assert.EqualError(t,
		e.CancelTransferFunds(ctx, &types.CancelTransferFunds{
			TransferID: transferID,
			Party:      "NOPE",
		}),
		banking.ErrCannotCancelOtherPartiesRecurringTransfers.Error(),
	)

	// now we move in time just a bit so we get some transfer processed, but then cancel before
	// then end of the transfer
	fromAcc := types.Account{
		Balance: num.NewUint(1000),
	}

	// asset exists
	e.col.EXPECT().GetPartyGeneralAccount(gomock.Any(), gomock.Any()).Times(1).Return(&fromAcc, nil)

	// assert the calculation of fees and transfer request are correct
	e.col.EXPECT().TransferFunds(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(1).DoAndReturn(
		func(ctx context.Context,
			transfers []*types.Transfer,
			accountTypes []types.AccountType,
			references []string,
			feeTransfers []*types.Transfer,
			feeTransfersAccountTypes []types.AccountType,
		) ([]*types.TransferResponse, error,
		) {
			t.Run("ensure transfers are correct", func(t *testing.T) {
				// transfer is done fully instantly, we should have 2 transfer
				assert.Len(t, transfers, 2)
				assert.Equal(t, transfers[0].Owner, "from")
				assert.Equal(t, transfers[0].Amount.Amount, num.NewUint(100))
				assert.Equal(t, transfers[0].Amount.Asset, "eth")

				// 1 account types too
				assert.Len(t, accountTypes, 2)
				assert.Equal(t, accountTypes[0], types.AccountTypeGeneral)
			})

			t.Run("ensure fee transfers are correct", func(t *testing.T) {
				assert.Len(t, feeTransfers, 1)
				assert.Equal(t, feeTransfers[0].Owner, "from")
				assert.Equal(t, feeTransfers[0].Amount.Amount, num.NewUint(50))
				assert.Equal(t, feeTransfers[0].Amount.Asset, "eth")

				// then the fees account types
				assert.Len(t, feeTransfersAccountTypes, 1)
				assert.Equal(t, accountTypes[0], types.AccountTypeGeneral)
			})

			return nil, nil
		})

	e.OnEpoch(context.Background(), types.Epoch{Seq: 10, Action: vega.EpochAction_EPOCH_ACTION_START})
	e.OnEpoch(context.Background(), types.Epoch{Seq: 10, Action: vega.EpochAction_EPOCH_ACTION_END})

	// now we cancel it, we should get no error and and event
	e.broker.EXPECT().Send(gomock.Any()).DoAndReturn(func(evt events.Event) {
		t.Run("ensure transfer is done", func(t *testing.T) {
			e, ok := evt.(*events.TransferFunds)
			assert.True(t, ok, "unexpected event from the bus")
			assert.Equal(t, e.Proto().Status, types.TransferStatusCancelled)
		})
	})

	assert.NoError(t,
		e.CancelTransferFunds(ctx, &types.CancelTransferFunds{
			TransferID: transferID,
			Party:      partyID,
		}),
	)

	// now we move in time, the recurring transfer was suppose to go
	// 'til epoch 11, but it's not cancelled, and nothing should happen
	e.OnEpoch(context.Background(), types.Epoch{Seq: 11, Action: vega.EpochAction_EPOCH_ACTION_START})
	e.OnEpoch(context.Background(), types.Epoch{Seq: 11, Action: vega.EpochAction_EPOCH_ACTION_END})
}