package warehouse

import (
	"context"

	"github.com/go-resty/resty/v2"

	"webinar-testing/pkg/models/cart"
)

type Options struct {
	Addr string
}

type warehouse struct {
	client *resty.Client
}

func New(opts *Options) *warehouse {
	cl := resty.New().SetBaseURL(opts.Addr)

	return &warehouse{client: cl}
}

func (w *warehouse) Check(_ context.Context, goodId cart.GoodID) (n int, err error) {
	var res int

	_, err = w.client.R().SetResult(&res).Get(string(goodId))
	if err != nil {
		return 0, err
	}

	return res, nil
}
