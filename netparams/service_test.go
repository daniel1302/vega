// Copyright (c) 2022 Gobalsky Labs Limited
//
// Use of this software is governed by the Business Source License included
// in the LICENSE file and at https://www.mariadb.com/bsl11.
//
// Change Date: 18 months from the later of the date of the first publicly
// available Distribution of this version of the repository, and 25 June 2022.
//
// On the date above, in accordance with the Business Source License, use
// of this software will be governed by version 3 or later of the GNU General
// Public License.

package netparams_test

import (
	"context"
	"testing"
	"time"

	"code.vegaprotocol.io/data-node/netparams"
	types "code.vegaprotocol.io/protos/vega"
	"code.vegaprotocol.io/vega/events"
	"github.com/stretchr/testify/assert"
)

type serviceTest struct {
	*netparams.Service
	ctx   context.Context
	cfunc context.CancelFunc
}

func getServiceTest(t *testing.T) *serviceTest {
	ctx, cfunc := context.WithCancel(context.Background())
	s := netparams.NewService(ctx)
	return &serviceTest{
		Service: s,
		ctx:     ctx,
		cfunc:   cfunc,
	}
}

func TestGetAllNetParams(t *testing.T) {
	svc := getServiceTest(t)
	evts := []*events.NetworkParameter{
		events.NewNetworkParameterEvent(svc.ctx, "key1", "value1"),
		events.NewNetworkParameterEvent(svc.ctx, "key2", "value2"),
		events.NewNetworkParameterEvent(svc.ctx, "key3", "value3"),
	}

	svc.Push(evts[0], evts[1], evts[2])

	var (
		hasError = true
		retries  = 50
	)

	for hasError && retries > 0 {
		retries -= 1
		all := svc.GetAll()
		// we expect 3 elements to be returned
		hasError = len(all) != 3
		time.Sleep(50 * time.Millisecond)
	}

	hasNP := func(nps []types.NetworkParameter, k, v string) bool {
		for _, np := range nps {
			if np.Key == k && np.Value == v {
				return true
			}
		}
		return false
	}

	all := svc.GetAll()
	assert.Len(t, all, 3)
	assert.True(t, hasNP(all, "key1", "value1"))
	assert.True(t, hasNP(all, "key2", "value2"))
	assert.True(t, hasNP(all, "key3", "value3"))
}
