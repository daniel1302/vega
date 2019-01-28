package markets

import (
	"vega/internal/storage"
	"vega/msg"
	"context"
	"vega/log"
	"vega/internal/logging"
)

//Service provides the interface for VEGA markets business logic.
type Service interface {
	// CreateMarket stores the given market.
	CreateMarket(ctx context.Context, market *msg.Market) error
	// GetByName searches for the given market by name.
	GetByName(ctx context.Context, name string) (*msg.Market, error)
	// GetAll returns all markets.
	GetAll(ctx context.Context) ([]*msg.Market, error)
	// GetDepth returns the market depth for the given market.
	GetDepth(ctx context.Context, market string) (marketDepth msg.MarketDepth, err error)
	// ObserveMarket provides a way to listen to changes on VEGA markets.
	ObserveMarkets(ctx context.Context) (markets <-chan []msg.Market, ref uint64)
	// ObserveDepth provides a way to listen to changes on the Depth of Market for a given market.
	ObserveDepth(ctx context.Context, market string) (depth <-chan msg.MarketDepth, ref uint64)
}

type marketService struct {
	*Config
	marketStore storage.MarketStore
	orderStore storage.OrderStore
}

// NewService creates an market service with the necessary dependencies
func NewService(marketStore storage.MarketStore, orderStore storage.OrderStore) Service {
	config := NewConfig()
	return &marketService{
		config,
		marketStore,
		orderStore,
	}
}

// CreateMarket stores the given market.
func (s *marketService) CreateMarket(ctx context.Context, party *msg.Market) error {
	return s.marketStore.Post(party)
}

// GetByName searches for the given market by name.
func (s *marketService) GetByName(ctx context.Context, name string) (*msg.Market, error) {
	p, err := s.marketStore.GetByName(name)
	return p, err
}

// GetAll returns all markets.
func (s *marketService) GetAll(ctx context.Context) ([]*msg.Market, error) {
	p, err := s.marketStore.GetAll()
	return p, err
}

// GetDepth returns the market depth for the given market.
func (s *marketService) GetDepth(ctx context.Context, market string) (marketDepth msg.MarketDepth, err error) {
	m, err := s.marketStore.GetByName(market)
	if err != nil {
		return msg.MarketDepth{}, err
	}
	return s.orderStore.GetMarketDepth(m.Name)
}

// ObserveDepth provides a way to listen to changes on the Depth of Market for a given market.
func (s *marketService) ObserveDepth(ctx context.Context, market string) (<-chan msg.MarketDepth, uint64) {
	depth := make(chan msg.MarketDepth)
	internal := make(chan []msg.Order)
	ref := s.orderStore.Subscribe(internal)

	go func(id uint64, internal chan []msg.Order, ctx context.Context) {
		ip := logging.IPAddressFromContext(ctx)
		<-ctx.Done()
		log.Debugf("MarketService -> depth closed connection: %d [%s]", id, ip)
		err := s.orderStore.Unsubscribe(id)
		if err != nil {
			log.Errorf("Error un-subscribing depth when context.Done() on MarketService for subscriber %d [%s]: %s", id, ip, err)
		}
	}(ref, internal, ctx)

	go func(id uint64, ctx context.Context) {
		ip := logging.IPAddressFromContext(ctx)
		for range internal {
			d, err := s.orderStore.GetMarketDepth(market)
			if err != nil {
				log.Errorf("Error calculating market depth for subscriber %d [%s]: %s", ref, ip, err)
			} else {
				select {
				case depth <- d:
					log.Debugf("MarketService -> depth for subscriber %d [%s] sent successfully", ref, ip)
				default:
					log.Debugf("MarketService -> depth for subscriber %d [%s] not sent", ref, ip)
				}
			}
		}
		log.Debugf("MarketService -> Channel for depth subscriber %d [%s] has been closed", ref, ip)
	}(ref, ctx)

	return depth, ref
}

func (s *marketService) ObserveMarkets(ctx context.Context) (markets <-chan []msg.Market, ref uint64) {
	 return nil, 0
}
