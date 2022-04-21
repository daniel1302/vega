package sqlstore_test

import (
	"context"
	"fmt"
	"testing"

	"code.vegaprotocol.io/data-node/entities"
	"code.vegaprotocol.io/data-node/sqlstore"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testAssetCount int = 0

func addTestAsset(t *testing.T, as *sqlstore.Assets, block entities.Block) entities.Asset {
	// Make an asset
	testAssetCount += 1
	totalSupply, _ := decimal.NewFromString("1000000000000000000001")

	asset := entities.Asset{
		ID:            entities.NewAssetID(generateID()),
		Name:          fmt.Sprint("my test asset", testAssetCount),
		Symbol:        fmt.Sprint("TEST", testAssetCount),
		TotalSupply:   totalSupply,
		Decimals:      5,
		Quantum:       10,
		ERC20Contract: "",
		VegaTime:      block.VegaTime,
	}

	// Add it to the database
	err := as.Add(context.Background(), asset)
	require.NoError(t, err)
	return asset
}

func TestAsset(t *testing.T) {
	defer DeleteEverything()

	bs := sqlstore.NewBlocks(connectionSource)
	block := addTestBlock(t, bs)

	as := sqlstore.NewAssets(connectionSource)
	ctx := context.Background()

	// Get all assets, there shouldn't be any yet
	assets, err := as.GetAll(ctx)
	require.NoError(t, err)
	require.Empty(t, assets)

	asset := addTestAsset(t, as, block)

	// Add it again, we should get a primary key violation
	err = as.Add(context.Background(), asset)
	assert.Error(t, err)

	// Query and check we've got back an asset the same as the one we put in
	fetchedAsset, err := as.GetByID(ctx, asset.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, asset, fetchedAsset)

	// Get all assets and make sure there's one more than there was to begin with
	assets, err = as.GetAll(ctx)
	assert.NoError(t, err)
	assert.Len(t, assets, 1)
}
