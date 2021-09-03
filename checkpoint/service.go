package checkpoint

import (
	"sort"

	"code.vegaprotocol.io/data-node/logging"
	protoapi "code.vegaprotocol.io/protos/data-node/api/v1"
	eventspb "code.vegaprotocol.io/protos/vega/events/v1"
)

type CheckpointStore interface {
	GetAll() ([]*eventspb.CheckpointEvent, error)
}

type Svc struct {
	Config
	store CheckpointStore
	log   *logging.Logger
}

func NewService(log *logging.Logger, cfg Config, store CheckpointStore) *Svc {
	log = log.Named(namedLogger)
	log.SetLevel(cfg.Level.Get())
	return &Svc{
		Config: cfg,
		store:  store,
		log:    log,
	}
}

// ReloadConf update the internal configuration of the order service
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

// GetAll fetches all checkpoints from storage, sorts, and converts them for API use
func (s *Svc) GetAll() ([]*protoapi.Checkpoint, error) {
	cps, err := s.store.GetAll()
	if err != nil {
		return nil, err
	}

	// default sort by block height, descending
	sort.SliceStable(cps, func(i, j int) bool {
		return cps[i].BlockHeight > cps[j].BlockHeight
	})

	// convert and return
	ret := make([]*protoapi.Checkpoint, 0, len(cps))
	for _, cp := range cps {
		ret = append(ret, &protoapi.Checkpoint{
			Hash:      cp.Hash,
			BlockHash: cp.BlockHash,
			AtBlock:   cp.BlockHeight,
		})
	}
	return ret, nil
}
