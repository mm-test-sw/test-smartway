package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"test-smartway/internal/api/middleware"
	"test-smartway/internal/entity"
)

type accountHandler struct {
	accountService entity.IAccountService
	logger         *zap.Logger
}

func RegisterAccountHandlers(router *mux.Router, service entity.IAccountService, logger *zap.Logger, mw *middleware.Middleware) {
	handler := new(accountHandler)
	handler.accountService = service
	handler.logger = logger

	r := router.PathPrefix("/accounts").Subrouter()

	r.Use(mw.PanicRecovery)
	r.Use(mw.Timeout)
	r.Use(mw.RequestId)
	r.Use(mw.ContentTypeJSON)
	r.Use(mw.DebugLogger)

	r.HandleFunc("", handler.AddAccount).Methods(http.MethodPost)
	r.HandleFunc("/{id}", handler.DeleteAccount).Methods(http.MethodDelete)
	r.HandleFunc("", handler.PutAccount).Methods(http.MethodPut)
	r.HandleFunc("/{id}/airlines", handler.GetAirlines).Methods(http.MethodGet)
}

func (h accountHandler) AddAccount(w http.ResponseWriter, r *http.Request) {
	account := new(entity.Account)

	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		resp, code := entity.HandleError(r.Context(), h.logger, entity.ErrorBadRequest)
		w.WriteHeader(code)
		w.Write(resp)
		return
	}

	account, err = h.accountService.AddAccount(r.Context(), account)
	if err != nil {
		resp, code := entity.HandleError(r.Context(), h.logger, err)
		w.WriteHeader(code)
		w.Write(resp)
		return
	}

	resp, _ := json.Marshal(account)
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}

func (h accountHandler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := h.accountService.DeleteAccount(r.Context(), id)
	if err != nil {
		resp, code := entity.HandleError(r.Context(), h.logger, err)
		w.WriteHeader(code)
		w.Write(resp)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h accountHandler) PutAccount(w http.ResponseWriter, r *http.Request) {
	account := new(entity.Account)

	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		resp, code := entity.HandleError(r.Context(), h.logger, entity.ErrorBadRequest)
		w.WriteHeader(code)
		w.Write(resp)
		return
	}

	account, err = h.accountService.PutAccount(r.Context(), account)
	if err != nil {
		resp, code := entity.HandleError(r.Context(), h.logger, err)
		w.WriteHeader(code)
		w.Write(resp)
		return
	}

	resp, _ := json.Marshal(account)
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (h accountHandler) GetAirlines(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	airlines, err := h.accountService.GetAirlines(r.Context(), id)
	if err != nil {
		resp, code := entity.HandleError(r.Context(), h.logger, err)
		w.WriteHeader(code)
		w.Write(resp)
		return
	}

	resp, _ := json.Marshal(airlines)
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
