package cmd_test

import (
	"testing"

	vgrand "code.vegaprotocol.io/vega/libs/rand"
	cmd "code.vegaprotocol.io/vega/cmd/vegawallet/commands"
	"code.vegaprotocol.io/vega/cmd/vegawallet/commands/flags"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRevokePermissionsFlags(t *testing.T) {
	t.Run("Valid flags succeeds", testRevokePermissionsFlagsValidFlagsSucceeds)
	t.Run("Missing flags fails", testRevokePermissionsFlagsMissingFlagsFails)
}

func testRevokePermissionsFlagsValidFlagsSucceeds(t *testing.T) {
	// given
	testDir := t.TempDir()
	passphrase, passphraseFilePath := NewPassphraseFile(t, testDir)
	f := &cmd.RevokePermissionsFlags{
		Wallet:         vgrand.RandomStr(10),
		Hostname:       vgrand.RandomStr(10),
		PassphraseFile: passphraseFilePath,
		Force:          true,
	}

	// when
	req, err := f.Validate()

	// then
	require.NoError(t, err)
	require.NotNil(t, req)
	assert.Equal(t, f.Hostname, req.Hostname)
	assert.Equal(t, f.Wallet, req.Wallet)
	assert.Equal(t, passphrase, req.Passphrase)
}

func testRevokePermissionsFlagsMissingFlagsFails(t *testing.T) {
	testDir := t.TempDir()
	walletName := vgrand.RandomStr(10)
	hostname := vgrand.RandomStr(10)
	_, passphraseFilePath := NewPassphraseFile(t, testDir)

	tcs := []struct {
		name        string
		flags       *cmd.RevokePermissionsFlags
		missingFlag string
	}{
		{
			name: "without hostname",
			flags: &cmd.RevokePermissionsFlags{
				Wallet:         walletName,
				Hostname:       "",
				PassphraseFile: passphraseFilePath,
			},
			missingFlag: "hostname",
		}, {
			name: "without wallet",
			flags: &cmd.RevokePermissionsFlags{
				Wallet:         "",
				Hostname:       hostname,
				PassphraseFile: passphraseFilePath,
			},
			missingFlag: "wallet",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			// when
			req, err := tc.flags.Validate()

			// then
			assert.ErrorIs(t, err, flags.FlagMustBeSpecifiedError(tc.missingFlag))
			require.Nil(t, req)
		})
	}
}
