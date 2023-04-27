package inmemory

import (
	"context"
	"sync"

	"webinar-testing/pkg/models/cart"
	"webinar-testing/pkg/models/errs"
)

type storage struct {
	mtx  sync.RWMutex
	data map[cart.UserID]cart.Cart
}

func (s *storage) Add(_ context.Context, order cart.Order) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	ct, ok := s.data[order.UserID]
	if !ok {
		s.data[order.UserID] = cart.Cart{
			UserID: "",
			Goods:  map[cart.GoodID]int{order.Good: order.Count},
		}

		return nil
	}

	val := ct.Goods[order.Good]
	ct.Goods[order.Good] = order.Count + val

	return nil
}

func (s *storage) GetByUserID(_ context.Context, id cart.UserID) (cart.Cart, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	ct, ok := s.data[id]
	if !ok {
		return cart.Cart{}, errs.ErrRecordNotFound
	}

	return ct, nil
}

func New() *storage {
	return &storage{
		data: make(map[cart.UserID]cart.Cart),
	}
}
