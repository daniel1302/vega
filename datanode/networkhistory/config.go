package networkhistory

import (
	"time"

	"code.vegaprotocol.io/vega/datanode/config/encoding"
	"code.vegaprotocol.io/vega/datanode/networkhistory/snapshot"
	"code.vegaprotocol.io/vega/datanode/networkhistory/store"
	"code.vegaprotocol.io/vega/logging"
)

type Config struct {
	Level         encoding.LogLevel `long:"log-level"`
	Enabled       encoding.Bool     `long:"enabled" description:"set to false to disable network history"`
	WipeOnStartup encoding.Bool     `long:"wipe-on-startup" description:"deprecated and ignored, use data-node unsafe_reset_all command instead"`

	Publish encoding.Bool `long:"publish" description:"if true this node will create and publish network history segments"`

	Store    store.Config    `group:"Store" namespace:"store"`
	Snapshot snapshot.Config `group:"Snapshot" namespace:"snapshot"`

	FetchRetryMax int               `long:"fetch-retry-max" description:"maximum number of times to retry fetching segments - default 10"`
	RetryTimeout  encoding.Duration `long:"retry-timeout" description:"time to wait between retries, increases with each retry - default 5s"`

	Initialise InitializationConfig `group:"Initialise" namespace:"initialise"`
}

// NewDefaultConfig creates an instance of the package specific configuration, given a
// pointer to a logger instance to be used for logging within the package.
func NewDefaultConfig() Config {
	return Config{
		Level:         encoding.LogLevel{Level: logging.InfoLevel},
		Enabled:       true,
		Publish:       true,
		Store:         store.NewDefaultConfig(),
		Snapshot:      snapshot.NewDefaultConfig(),
		FetchRetryMax: 10,
		RetryTimeout:  encoding.Duration{Duration: 5 * time.Second},
		Initialise:    NewDefaultInitializationConfig(),
	}
}

func NewDefaultInitializationConfig() InitializationConfig {
	return InitializationConfig{
		MinimumBlockCount: 1,
		TimeOut:           encoding.Duration{Duration: 1 * time.Minute},
		GrpcAPIPorts:      []int{},
		ToSegment:         "",
	}
}

type InitializationConfig struct {
	ToSegment         string            `long:"to-segment" description:"the segment to initialise up to, if omitted the datanode will attempt to fetch the latest segment from the network"`
	MinimumBlockCount int64             `long:"block-count" description:"the minimum number of blocks to fetch"`
	TimeOut           encoding.Duration `long:"timeout" description:"maximum time allowed to auto-initialise the node"`
	GrpcAPIPorts      []int             `long:"grpc-api-ports" description:"list of additional ports to check to for api connection when getting latest segment"`
}
