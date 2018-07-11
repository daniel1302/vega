package blockchain

import (
	"fmt"
	"vega/core"

	"github.com/tendermint/tendermint/abci/server"
	cmn "github.com/tendermint/tmlibs/common"
)

// Starts up a Vega blockchain server.
func Start(vega *core.Vega) error {
	fmt.Println("Starting Vega blockchain socket...")
	blockchain := NewBlockchain(vega)
	srv, err := server.NewServer("127.0.0.1:26658", "socket", blockchain)
	if err != nil {
		return err
	}

	if err := srv.Start(); err != nil {
		return err
	}
	// Wait forever
	cmn.TrapSignal(func() {
		srv.Stop()
	})
	return nil
}
