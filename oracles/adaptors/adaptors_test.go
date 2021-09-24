package adaptors_test

import (
	"encoding/hex"
	"encoding/json"
	"testing"

	commandspb "code.vegaprotocol.io/protos/vega/commands/v1"
	"code.vegaprotocol.io/vega/crypto"
	"code.vegaprotocol.io/vega/oracles"
	"code.vegaprotocol.io/vega/oracles/adaptors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAdaptors(t *testing.T) {
	t.Run("Creating adaptors succeeds", testCreatingAdaptorsSucceeds)
	t.Run("Normalising data from unknown oracle fails", testAdaptorsNormalisingDataFromUnknownOracleFails)
	t.Run("Normalising data from known oracle succeeds", testAdaptorsNormalisingDataFromKnownOracleSucceeds)
}

func testCreatingAdaptorsSucceeds(t *testing.T) {
	// when
	as := adaptors.New()

	// then
	assert.NotNil(t, as)
}

func testAdaptorsNormalisingDataFromUnknownOracleFails(t *testing.T) {
	// given
	pubKeyB := []byte("0xdeadbeef")
	pubKey := crypto.NewPublicKeyOrAddress(hex.EncodeToString(pubKeyB), pubKeyB)
	rawData := commandspb.OracleDataSubmission{
		Source:  commandspb.OracleDataSubmission_ORACLE_SOURCE_UNSPECIFIED,
		Payload: dummyOraclePayload(),
	}

	// when
	normalisedData, err := stubbedAdaptors().Normalise(pubKey, rawData)

	// then
	require.Error(t, err)
	assert.EqualError(t, err, adaptors.ErrUnknownOracleSource.Error())
	assert.Nil(t, normalisedData)
}

func testAdaptorsNormalisingDataFromKnownOracleSucceeds(t *testing.T) {
	tcs := []struct {
		name   string
		source commandspb.OracleDataSubmission_OracleSource
	}{
		{
			name:   "with Open Oracle source",
			source: commandspb.OracleDataSubmission_ORACLE_SOURCE_OPEN_ORACLE,
		}, {
			name:   "with JSON source",
			source: commandspb.OracleDataSubmission_ORACLE_SOURCE_JSON,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			// given
			pubKeyB := []byte("0xdeadbeef")
			pubKey := crypto.NewPublicKeyOrAddress(hex.EncodeToString(pubKeyB), pubKeyB)
			rawData := commandspb.OracleDataSubmission{
				Source:  tc.source,
				Payload: dummyOraclePayload(),
			}

			// when
			normalisedData, err := stubbedAdaptors().Normalise(pubKey, rawData)

			// then
			require.NoError(t, err)
			assert.NotNil(t, normalisedData)
		})
	}
}

func stubbedAdaptors() *adaptors.Adaptors {
	return &adaptors.Adaptors{
		Adaptors: map[commandspb.OracleDataSubmission_OracleSource]adaptors.Adaptor{
			commandspb.OracleDataSubmission_ORACLE_SOURCE_OPEN_ORACLE: &dummyOracleAdaptor{},
			commandspb.OracleDataSubmission_ORACLE_SOURCE_JSON:        &dummyOracleAdaptor{},
		},
	}
}

func dummyOraclePayload() []byte {
	payload, err := json.Marshal(map[string]string{
		"field_1": "value_1",
		"field_2": "value_2",
	})
	if err != nil {
		panic("failed to generate random oracle payload in tests")
	}

	return payload
}

type dummyOracleAdaptor struct {
}

func (d *dummyOracleAdaptor) Normalise(_ crypto.PublicKeyOrAddress, payload []byte) (*oracles.OracleData, error) {
	data := &oracles.OracleData{}
	err := json.Unmarshal(payload, data)
	return data, err
}
