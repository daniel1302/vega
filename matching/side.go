package matching

import (
	"fmt"
	"sort"

	"code.vegaprotocol.io/vega/logging"
	"code.vegaprotocol.io/vega/metrics"
	types "code.vegaprotocol.io/vega/proto"
	"github.com/pkg/errors"
)

var (
	// ErrPriceNotFound signals that a price was not found on the book side
	ErrPriceNotFound = errors.New("price-volume pair not found")
	// ErrNoOrder signals that there's no orders on the book side.
	ErrNoOrder = errors.New("no orders in the book side")
)

// OrderBookSide reprenset a side of the book, either Sell or Buy
type OrderBookSide struct {
	log *logging.Logger
	// Config
	levels []*PriceLevel
}

func (s *OrderBookSide) addOrder(o *types.Order, side types.Side) {
	// update the price-volume map

	s.getPriceLevel(o.Price, side).addOrder(o)
}

func (s *OrderBookSide) getHighestOrderPrice(side types.Side) (uint64, error) {
	if len(s.levels) <= 0 {
		return 0, ErrNoOrder
	}
	// sell order descending
	if side == types.Side_SIDE_SELL {
		return s.levels[0].price, nil
	}
	// buy order ascending
	return s.levels[len(s.levels)-1].price, nil
}

func (s *OrderBookSide) getLowestOrderPrice(side types.Side) (uint64, error) {
	if len(s.levels) <= 0 {
		return 0, ErrNoOrder
	}
	// sell order descending
	if side == types.Side_SIDE_SELL {
		return s.levels[len(s.levels)-1].price, nil
	}
	// buy order ascending
	return s.levels[0].price, nil
}

func (s OrderBookSide) BestPriceAndVolume(side types.Side) (uint64, uint64) {
	if len(s.levels) <= 0 {
		return 0, 0
	}
	return s.levels[len(s.levels)-1].price, s.levels[len(s.levels)-1].volume
}

func (s *OrderBookSide) amendOrder(orderAmend *types.Order) error {
	priceLevelIndex := -1
	orderIndex := -1
	var oldOrder *types.Order

	for idx, priceLevel := range s.levels {
		if priceLevel.price == orderAmend.Price {
			priceLevelIndex = idx
			for j, order := range priceLevel.orders {
				if order.Id == orderAmend.Id {
					orderIndex = j
					oldOrder = order
					break
				}
			}
			break
		}
	}

	if oldOrder == nil || priceLevelIndex == -1 || orderIndex == -1 {
		return types.ErrOrderNotFound
	}

	if oldOrder.PartyID != orderAmend.PartyID {
		return types.ErrOrderAmendFailure
	}

	if oldOrder.Remaining < orderAmend.Size {
		return types.ErrOrderAmendFailure
	}

	if oldOrder.Reference != orderAmend.Reference {
		return types.ErrOrderAmendFailure
	}

	reduceBy := oldOrder.Remaining - orderAmend.Size
	*s.levels[priceLevelIndex].orders[orderIndex] = *orderAmend
	s.levels[priceLevelIndex].reduceVolume(reduceBy)
	return nil
}

// RemoveOrder will remove an order from the book
func (s *OrderBookSide) RemoveOrder(o *types.Order) (*types.Order, error) {
	// first  we try to find the pricelevel of the order
	var i int
	if o.Side == types.Side_SIDE_BUY {
		i = sort.Search(len(s.levels), func(i int) bool { return s.levels[i].price >= o.Price })
	} else {
		// sell side levels should be ordered in ascending
		i = sort.Search(len(s.levels), func(i int) bool { return s.levels[i].price <= o.Price })
	}
	// we did not found the level
	// then the order do not exists in the price level
	if i >= len(s.levels) {
		return nil, types.ErrOrderNotFound
	}

	// orders are order by timestamp (CreatedAt)
	oidx := sort.Search(len(s.levels[i].orders), func(j int) bool {
		return s.levels[i].orders[j].CreatedAt >= o.CreatedAt
	})
	// we did not find the order
	if oidx >= len(s.levels[i].orders) {
		return nil, types.ErrOrderNotFound
	}

	// now we may have a few orders with the same timestamp
	// lets iterate over them in order to find the right one
	finaloidx := -1
	for oidx < len(s.levels[i].orders) && s.levels[i].orders[oidx].CreatedAt == o.CreatedAt {
		if s.levels[i].orders[oidx].Id == o.Id {
			finaloidx = oidx
			break
		}
		oidx++
	}

	var order *types.Order
	// remove the order from the
	if finaloidx != -1 {
		order = s.levels[i].orders[finaloidx]
		s.levels[i].removeOrder(finaloidx)
	} else {
		// We could not find the matching order, return an error
		return nil, types.ErrOrderNotFound
	}

	if len(s.levels[i].orders) <= 0 {
		s.levels = s.levels[:i+copy(s.levels[i:], s.levels[i+1:])]
	}

	return order, nil
}

func (s *OrderBookSide) getPriceLevel(price uint64, side types.Side) *PriceLevel {
	var i int
	if side == types.Side_SIDE_BUY {
		// buy side levels should be ordered in descending
		i = sort.Search(len(s.levels), func(i int) bool { return s.levels[i].price >= price })
	} else {
		// sell side levels should be ordered in ascending
		i = sort.Search(len(s.levels), func(i int) bool { return s.levels[i].price <= price })
	}

	// we found the level just return it.
	if i < len(s.levels) && s.levels[i].price == price {
		return s.levels[i]
	}

	// append new elem first to make sure we have enough place
	// this would reallocate sufficiently then
	// no risk of this being a empty order, as it's overwritten just next with
	// the slice insert
	level := NewPriceLevel(price)
	s.levels = append(s.levels, nil)
	copy(s.levels[i+1:], s.levels[i:])
	s.levels[i] = level
	return level
}

func (s *OrderBookSide) uncross(agg *types.Order) ([]*types.Trade, []*types.Order, uint64, error) {
	timer := metrics.NewTimeCounter("-", "matching", "OrderBookSide.uncross")

	var (
		trades                  []*types.Trade
		impactedOrders          []*types.Order
		lastTradedPrice         uint64
		totalVolumeToFill       uint64
		totalPrice, totalVolume uint64
	)

	if agg.TimeInForce == types.Order_TIF_FOK {
		totalVolume = agg.Remaining

		if agg.Side == types.Side_SIDE_SELL {
			for _, level := range s.levels {
				// in case of network trades, we want to calculate an accurate average price to return
				if agg.Type == types.Order_TYPE_NETWORK {
					totalVolumeToFill += level.volume
					factor := totalVolume
					if level.volume < totalVolume {
						factor = level.volume
						totalVolume -= level.volume
					}
					totalPrice += level.price * factor
				} else if level.price >= agg.Price || agg.Type == types.Order_TYPE_MARKET {
					totalVolumeToFill += level.volume
				}
				// No need to keep looking once we pass the required amount
				if totalVolumeToFill >= agg.Remaining {
					break
				}
			}
		}

		if agg.Side == types.Side_SIDE_BUY {
			for _, level := range s.levels {
				// in case of network trades, we want to calculate an accurate average price to return
				if agg.Type == types.Order_TYPE_NETWORK {
					totalVolumeToFill += level.volume
					factor := totalVolume
					if level.volume < totalVolume {
						factor = level.volume
						totalVolume -= level.volume
					}
					totalPrice += level.price * factor
				} else if level.price <= agg.Price || agg.Type == types.Order_TYPE_MARKET {
					totalVolumeToFill += level.volume
				}
				// No need to keep looking once we pass the required amount
				if totalVolumeToFill >= agg.Remaining {
					break
				}
			}
		}

		if agg.Type == types.Order_TYPE_NETWORK {
			// set avg price for order
			agg.Price = totalPrice / agg.Remaining
		}

		if s.log.GetLevel() == logging.DebugLevel {
			s.log.Debug(fmt.Sprintf("totalVolumeToFill %d until price %d, remaining %d\n", totalVolumeToFill, agg.Price, agg.Remaining))
		}

		if totalVolumeToFill < agg.Remaining {
			timer.EngineTimeCounterAdd()
			return trades, impactedOrders, 0, nil
		}
	}

	var (
		idx     = len(s.levels) - 1
		filled  bool
		ntrades []*types.Trade
		nimpact []*types.Order
		err     error
	)

	if agg.Side == types.Side_SIDE_SELL {
		// in here we iterate from the end, as it's easier to remove the
		// price levels from the back of the slice instead of from the front
		// also it will allow us to reduce allocations
		for !filled && idx >= 0 {
			if s.levels[idx].price >= agg.Price || agg.Type == types.Order_TYPE_MARKET {
				filled, ntrades, nimpact, err = s.levels[idx].uncross(agg)
				trades = append(trades, ntrades...)
				impactedOrders = append(impactedOrders, nimpact...)
				// break if a wash trade is detected
				if err != nil && err == ErrWashTrade {
					break
				}
				if len(s.levels[idx].orders) <= 0 {
					idx--
				}
			} else {
				break
			}
		}

		// now we nil the price levels that have been completely emptied out
		// then we resize the slice
		if idx < 0 || len(s.levels[idx].orders) > 0 {
			// do not remove this one as it's not emptied already
			idx++
		}
		if idx < len(s.levels) {
			// nil out the pricelevels so they get collected at some point
			for i := idx; i < len(s.levels); i++ {
				s.levels[i] = nil
			}
			s.levels = s.levels[:idx]
		}
	}

	if agg.Side == types.Side_SIDE_BUY {
		// in here we iterate from the end, as it's easier to remove the
		// price levels from the back of the slice instead of from the front
		// also it will allow us to reduce allocations
		for !filled && idx >= 0 {
			if s.levels[idx].price <= agg.Price || agg.Type == types.Order_TYPE_MARKET {
				filled, ntrades, nimpact, err = s.levels[idx].uncross(agg)
				trades = append(trades, ntrades...)
				impactedOrders = append(impactedOrders, nimpact...)
				if err != nil && err == ErrWashTrade {
					break
				}
				if len(s.levels[idx].orders) <= 0 {
					idx--
				}
			} else {
				break
			}
		}

		// now we nil the price levels that have been completely emptied out
		// then we resize the slice
		// idx can be < to 0 if we went through all price levels
		if idx < 0 || len(s.levels[idx].orders) > 0 {
			// do not remove this one as it's not emptied already
			idx++
		}
		if idx < len(s.levels) {
			// nil out the pricelevels so they get collected at some point
			for i := idx; i < len(s.levels); i++ {
				s.levels[i] = nil
			}
			s.levels = s.levels[:idx]
		}
	}

	if len(trades) > 0 {
		lastTradedPrice = trades[len(trades)-1].Price
	}
	timer.EngineTimeCounterAdd()
	return trades, impactedOrders, lastTradedPrice, err
}

func (s *OrderBookSide) getLevels() []*PriceLevel {
	return s.levels
}
