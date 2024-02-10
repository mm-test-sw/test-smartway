package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"test-smartway/internal/api/middleware"
	"test-smartway/internal/entity"
)

type providerHandler struct {
	providerService entity.IProviderService
	logger          *zap.Logger
}

func RegisterProviderHandlers(router *mux.Router, service entity.IProviderService, logger *zap.Logger, mw *middleware.Middleware) {
	handler := new(providerHandler)
	handler.providerService = service
	handler.logger = logger

	r := router.PathPrefix("/providers").Subrouter()

	r.Use(mw.PanicRecovery)
	r.Use(mw.Timeout)
	r.Use(mw.RequestId)
	r.Use(mw.ContentTypeJSON)
	r.Use(mw.DebugLogger)

	r.HandleFunc("", handler.AddProvider).Methods(http.MethodPost)
	r.HandleFunc("/{id}", handler.DeleteProvider).Methods(http.MethodDelete)
	r.HandleFunc("/{id}/airlines", handler.GetAirlines).Methods(http.MethodGet)
}

func (h providerHandler) AddProvider(w http.ResponseWriter, r *http.Request) {
	provider := new(entity.Provider)

	err := json.NewDecoder(r.Body).Decode(provider)
	if err != nil {
		resp, code := entity.HandleError(r.Context(), h.logger, entity.ErrorBadRequest)
		w.WriteHeader(code)
		w.Write(resp)
		return
	}

	provider, err = h.providerService.AddProvider(r.Context(), provider)
	if err != nil {
		resp, code := entity.HandleError(r.Context(), h.logger, err)
		w.WriteHeader(code)
		w.Write(resp)
		return
	}

	resp, _ := json.Marshal(provider)
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}

func (h providerHandler) DeleteProvider(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := h.providerService.DeleteProvider(r.Context(), id)
	if err != nil {
		resp, code := entity.HandleError(r.Context(), h.logger, err)
		w.WriteHeader(code)
		w.Write(resp)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h providerHandler) GetAirlines(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	airlines, err := h.providerService.GetAirlines(r.Context(), id)
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
