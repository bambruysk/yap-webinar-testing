package cart

import (
	"context"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"webinar-testing/pkg/models/cart"
	"webinar-testing/pkg/models/errs"
)

// CartStorager
//
//go:generate mockery --name CartStorager --with-expecter
type CartStorager interface {
	Add(ctx context.Context, order cart.Order) error
	GetByUserID(ctx context.Context, id cart.UserID) (cart.Cart, error)
}

// CartStorager
//
//go:generate mockery --name WarehouseConnector --with-expecter
type WarehouseConnector interface {
	Check(ctx context.Context, goodId cart.GoodID) (n int, err error)
}

type service struct {
	logger    *zap.Logger
	storage   CartStorager
	warehouse WarehouseConnector
}

func New(logger *zap.Logger, storage CartStorager, wh WarehouseConnector) *service {
	return &service{logger: logger, storage: storage, warehouse: wh}
}

func (s *service) Add(ctx context.Context, order cart.Order) error {
	if err := s.checkWarehouseGood(ctx, order); err != nil {
		return err
	}

	return s.storage.Add(ctx, order)
}

func (s *service) Get(ctx context.Context, user cart.UserID) (cart cart.Cart, err error) {
	return s.storage.GetByUserID(ctx, user)
}

func (s *service) checkWarehouseGood(ctx context.Context, order cart.Order) error {
	n, err := 1000, error(nil) // s.warehouse.Check(ctx, order.Good)
	if err != nil {
		return errors.Wrap(errs.ErrWarehouseConnect, err.Error())
	}

	if n < order.Count {
		return errs.ErrRecordNotFound
	}

	return nil
}
