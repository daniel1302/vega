package pow

import (
	"context"

	"code.vegaprotocol.io/vega/types"
	"github.com/golang/protobuf/proto"
)

func (e *Engine) Namespace() types.SnapshotNamespace {
	return types.PoWSnapshot
}

func (e *Engine) Keys() []string {
	return e.hashKeys
}

func (e *Engine) Stopped() bool {
	return false
}

// get the serialised form and hash of the given key.
func (e *Engine) serialise(k string) ([]byte, error) {
	payload := types.Payload{
		Data: &types.PayloadProofOfWork{
			BlockHeight:   e.blockHeight,
			BlockHash:     e.blockHash,
			SeenTx:        e.seenTx,
			HeightToTx:    e.heightToTx,
			BannedParties: e.bannedParties,
		},
	}

	data, err := proto.Marshal(payload.IntoProto())
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (e *Engine) HasChanged(k string) bool {
	return true
}

func (e *Engine) GetState(k string) ([]byte, []types.StateProvider, error) {
	state, err := e.serialise(k)
	return state, nil, err
}

func (e *Engine) LoadState(ctx context.Context, p *types.Payload) ([]types.StateProvider, error) {
	if e.Namespace() != p.Data.Namespace() {
		return nil, types.ErrInvalidSnapshotNamespace
	}
	pl := p.Data.(*types.PayloadProofOfWork)
	e.bannedParties = pl.BannedParties
	e.blockHash = pl.BlockHash
	e.blockHeight = pl.BlockHeight
	e.heightToTx = pl.HeightToTx
	e.seenTx = pl.SeenTx
	return nil, nil
}
