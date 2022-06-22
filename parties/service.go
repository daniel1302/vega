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

package parties

import (
	"context"

	"code.vegaprotocol.io/data-node/logging"
	types "code.vegaprotocol.io/protos/vega"
)

// PartyStore ...
//go:generate go run github.com/golang/mock/mockgen -destination mocks/part_store_mock.go -package mocks code.vegaprotocol.io/data-node/parties PartyStore
type PartyStore interface {
	Post(party *types.Party) error
	GetByID(id string) (*types.Party, error)
	GetAll() ([]*types.Party, error)
}

// Svc represents the party service
type Svc struct {
	Config
	log   *logging.Logger
	store PartyStore
}

// NewService creates a Parties service with the necessary dependencies
func NewService(log *logging.Logger, config Config, store PartyStore) (*Svc, error) {
	// setup logger
	log = log.Named(namedLogger)
	log.SetLevel(config.Level.Get())

	// create the network party, as it's a builtin party
	// and required from the apis + the network can create orders which are stored
	// in the orders db
	err := store.Post(&types.Party{
		Id: "network",
	})
	if err != nil {
		return nil, err
	}

	return &Svc{
		log:    log,
		Config: config,
		store:  store,
	}, nil
}

// ReloadConf updates the internal configuration of the service
func (s *Svc) ReloadConf(cfg Config) {
	s.log.Info("reloading configuration")
	if s.log.GetLevel() != cfg.Level.Get() {
		s.log.Info("updating log level",
			logging.String("old", s.log.GetLevel().String()),
			logging.String("new", cfg.Level.String()),
		)
		s.log.SetLevel(cfg.Level.Get())
	}

	s.Config = cfg
}

// CreateParty stores the given party.
func (s *Svc) CreateParty(ctx context.Context, party *types.Party) error {
	return s.store.Post(party)
}

// GetByID searches for the given party by id.
func (s *Svc) GetByID(ctx context.Context, id string) (*types.Party, error) {
	return s.store.GetByID(id)
}

// GetAll returns all parties.
func (s *Svc) GetAll(ctx context.Context) ([]*types.Party, error) {
	return s.store.GetAll()
}
