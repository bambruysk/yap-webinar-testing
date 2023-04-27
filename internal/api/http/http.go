package http

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"

	"webinar-testing/pkg/models/cart"
)

type Options struct {
	Addr string
}

type Usecase interface {
	Add(ctx context.Context, order cart.Order) error
	Get(ctx context.Context, user cart.UserID) (cart cart.Cart, err error)
}

type server struct {
	opts    *Options
	logger  *zap.Logger
	router  *chi.Mux
	usecase Usecase
}

func (s *server) initMiddleware() {
	s.router.Use(
		middleware.RequestID,
		middleware.Logger,
	)
}

func (s *server) initHandlers() {
	s.router.Post("/", s.addHandler)
	s.router.Get("/{id}", s.getHandler)
}

func New(opts *Options, logger *zap.Logger, usecase Usecase) *server {
	s := &server{
		opts:    opts,
		logger:  logger,
		router:  chi.NewRouter(),
		usecase: usecase,
	}
	r := chi.NewRouter()

	r.Use()

	s.initMiddleware()
	s.initHandlers()

	return s
}

func (s *server) Run() {
	go func() {
		s.logger.Fatal("unable run server", zap.Error(
			http.ListenAndServe(s.opts.Addr, s.router),
		))
	}()
}
