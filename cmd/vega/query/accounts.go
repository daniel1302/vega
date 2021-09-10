package query

import (
	"errors"
	"fmt"

	coreapipb "code.vegaprotocol.io/protos/vega/coreapi/v1"

	"github.com/golang/protobuf/jsonpb"
)

type AccountsCmd struct {
	Party   AccountsPartyCmd   `command:"party" description:"List accounts for a given party"`
	Market  AccountsMarketCmd  `command:"market" description:"List accounts for a given market"`
	Network AccountsNetworkCmd `command:"network" description:"List accounts owned by the network"`
	Help    bool               `short:"h" long:"help" description:"Show this help message"`
}

type AccountsPartyCmd struct {
	NodeAddress string `long:"node-address" description:"The address of the vega node to use" default:"0.0.0.0:3002"`
	Market      string `long:"market" description:"An optional market"`
	Help        bool   `short:"h" long:"help" description:"Show this help message"`
}

type AccountsMarketCmd struct {
	NodeAddress string `long:"node-address" description:"The address of the vega node to use" default:"0.0.0.0:3002"`
	Help        bool   `short:"h" long:"help" description:"Show this help message"`
}

type AccountsNetworkCmd struct {
	NodeAddress string `long:"node-address" description:"The address of the vega node to use" default:"0.0.0.0:3002"`
	Help        bool   `short:"h" long:"help" description:"Show this help message"`
}

func (opts *AccountsPartyCmd) Execute(params []string) error {
	if len(params) > 1 {
		return errors.New("only one party needs to be specified")
	}

	if len(params) < 1 {
		return errors.New("one party is required")
	}

	req := coreapipb.ListAccountsRequest{
		Party:  params[0],
		Market: opts.Market, // at most empty anyway
	}

	return getPrintAccounts(opts.NodeAddress, &req)
}

func (opts *AccountsMarketCmd) Execute(params []string) error {
	if len(params) > 1 {
		return errors.New("only one market needs to be specified")
	}

	return nil
}

func (opts *AccountsNetworkCmd) Execute(params []string) error {
	req := coreapipb.ListAccountsRequest{}
	return getPrintAccounts(opts.NodeAddress, &req)
}

func getPrintAccounts(nodeAddress string, req *coreapipb.ListAccountsRequest) error {
	clt, err := getClient(nodeAddress)
	if err != nil {
		return fmt.Errorf("could not connect to the vega node: %w", err)
	}

	ctx, cancel := timeoutContext()
	defer cancel()
	res, err := clt.ListAccounts(ctx, req)
	if err != nil {
		return fmt.Errorf("error querying the vega node: %w", err)
	}

	m := jsonpb.Marshaler{
		Indent: "  ",
	}
	buf, err := m.MarshalToString(res)
	if err != nil {
		return fmt.Errorf("invalid response from vega node: %w", err)
	}

	fmt.Printf("%v", string(buf))

	return nil
}
