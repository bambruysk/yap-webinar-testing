package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"

	"webinar-testing/pkg/models/cart"
	"webinar-testing/pkg/models/errs"
)

func (s *server) addHandler(w http.ResponseWriter, r *http.Request) {
	req := cart.Order{}
	if err := unmarshall(w, r, &req); err != nil {
		return
	}

	if err := s.usecase.Add(r.Context(), req); err != nil {
		switch {
		case errors.Is(err, errs.ErrWarehouseConnect):
			responseError(w, fmt.Sprintf("unable connect to warehouse: %v", err), http.StatusInternalServerError)
		case errors.Is(err, errs.ErrWarehouseNotHasGood):
			responseError(w, fmt.Sprintf("not enough goods: %v", err), http.StatusBadRequest)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s *server) getHandler(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	if len(userId) == 0 {
		responseError(w, "bad user id", http.StatusBadRequest)

		return
	}

	resp, err := s.usecase.Get(r.Context(), cart.NewUserIdFromString(userId))
	if err != nil {
		switch {
		case errors.Is(err, errs.ErrRecordNotFound):
			responseError(w, fmt.Sprintf("%v", err), http.StatusNotFound)

		}
		return
	}

	responseJSON(w, resp, http.StatusOK)

}

func responseError(w http.ResponseWriter, msg string, status int) {
	w.WriteHeader(status)
	w.Write([]byte(msg))
}

func responseJSON(w http.ResponseWriter, msg any, status int) {
	body, err := json.Marshal(msg)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

func unmarshall(w http.ResponseWriter, r *http.Request, v any) error {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		responseError(w, "unable read body", http.StatusInternalServerError)

		return err
	}

	if err := json.Unmarshal(body, v); err != nil {
		responseError(w, "unable read body", http.StatusInternalServerError)

		return err
	}

	return nil
}
