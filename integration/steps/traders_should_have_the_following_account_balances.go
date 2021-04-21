package steps

import (
	"fmt"

	"github.com/cucumber/godog/gherkin"

	"code.vegaprotocol.io/vega/integration/stubs"
	types "code.vegaprotocol.io/vega/proto"
)

func TradersShouldHaveTheFollowingAccountBalances(
	broker *stubs.BrokerStub,
	table *gherkin.DataTable,
) error {
	for _, r := range TableWrapper(*table).Parse() {
		row := accountBalancesRow{row: r}
		var hasError bool

		generalAccount, err := broker.GetTraderGeneralAccount(row.trader(), row.asset())
		if err != nil {
			return errCannotGetTraderGeneralAccount(row.trader(), row.asset(), err)
		}
		if generalAccount.GetBalance() != row.generalAccountBalance() {
			hasError = true
		}

		marginAccount, err := broker.GetTraderMarginAccount(row.trader(), row.marketID())
		if err != nil {
			return errCannotGetTraderMarginAccount(row.trader(), row.asset(), err)
		}
		if marginAccount.GetBalance() != row.marginAccountBalance() {
			hasError = true
		}

		if hasError {
			return errMismatchedAccountBalances(row, marginAccount, generalAccount)
		}

	}
	return nil
}

func errCannotGetTraderGeneralAccount(trader, asset string, err error) error {
	return fmt.Errorf("couldn't get general account for trader(%s) and asset(%s): %s",
		trader, asset, err.Error(),
	)
}

func errCannotGetTraderMarginAccount(trader, asset string, err error) error {
	return fmt.Errorf("couldn't get margin account for trader(%s) and asset(%s): %s",
		trader, asset, err.Error(),
	)
}

func errMismatchedAccountBalances(row accountBalancesRow, marginAccount types.Account, generalAccount types.Account) error {
	return formatDiff(
		fmt.Sprintf("account balances did not match for party(%s)", row.trader()),
		map[string]string{
			"margin account balance":  u64ToS(row.marginAccountBalance()),
			"general account balance": u64ToS(row.generalAccountBalance()),
		},
		map[string]string{
			"margin account balance": u64ToS(marginAccount.GetBalance()),
			"general account balance": u64ToS(generalAccount.GetBalance()),
		},
	)
}

type accountBalancesRow struct {
	row RowWrapper
}

func (r accountBalancesRow) trader() string {
	return r.row.MustStr("trader")
}

func (r accountBalancesRow) asset() string {
	return r.row.MustStr("asset")
}

func (r accountBalancesRow) marketID() string {
	return r.row.MustStr("market id")
}

func (r accountBalancesRow) marginAccountBalance() uint64 {
	return r.row.MustU64("margin")
}

func (r accountBalancesRow) generalAccountBalance() uint64 {
	return r.row.MustU64("general")
}
