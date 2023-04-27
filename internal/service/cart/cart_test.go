package cart

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"webinar-testing/internal/service/cart/mocks"
	"webinar-testing/pkg/models/cart"
	"webinar-testing/pkg/models/errs"
)

func Test_service_checkWarehouseGood(t *testing.T) {
	type fields struct {
		logger    *zap.Logger
		warehouse *mocks.WarehouseConnector
	}
	type args struct {
		ctx   context.Context
		order cart.Order
	}
	type whChecks struct {
		count int
		err   error
	}

	tests := []struct {
		name     string
		fields   fields
		args     args
		whChecks whChecks
		wantErr  error
	}{
		{
			name: "Test Success",
			fields: fields{
				logger:    zap.NewNop(),
				warehouse: mocks.NewWarehouseConnector(t),
			},
			args: args{
				ctx: nil,
				order: cart.Order{
					Good:  "Phone",
					Count: 1000,
				},
			},
			whChecks: whChecks{
				count: 5000,
				err:   nil,
			},
			wantErr: nil,
		},
		{
			name: "Test Not Enough",
			fields: fields{
				logger:    zap.NewNop(),
				warehouse: mocks.NewWarehouseConnector(t),
			},
			args: args{
				ctx: nil,
				order: cart.Order{
					Good:  "Phone",
					Count: 5000,
				},
			},
			whChecks: whChecks{
				count: 100,
				err:   nil,
			},
			wantErr: errs.ErrWarehouseNotHasGood,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				logger:    tt.fields.logger,
				warehouse: tt.fields.warehouse,
			}

			tt.fields.warehouse.EXPECT().
				Check(tt.args.ctx, tt.args.order.Good).
				Return(tt.whChecks.count, tt.whChecks.err)

			err := s.checkWarehouseGood(tt.args.ctx, tt.args.order)

			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}
