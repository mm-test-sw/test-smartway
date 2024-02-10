package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"test-smartway/internal/api/middleware"
	"test-smartway/internal/entity"
)

type airlineHandler struct {
	airlineService entity.IAirlineService
	logger         *zap.Logger
}

func RegisterAirlineHandlers(router *mux.Router, service entity.IAirlineService, logger *zap.Logger, mw *middleware.Middleware) {
	handler := new(airlineHandler)
	handler.airlineService = service
	handler.logger = logger

	r := router.PathPrefix("/airlines").Subrouter()

	r.Use(mw.PanicRecovery)
	r.Use(mw.Timeout)
	r.Use(mw.RequestId)
	r.Use(mw.ContentTypeJSON)
	r.Use(mw.DebugLogger)

	r.HandleFunc("", handler.AddAirline).Methods(http.MethodPost)
	r.HandleFunc("/{code}", handler.DeleteAirline).Methods(http.MethodDelete)
	r.HandleFunc("/providers", handler.PutAirlineProviders).Methods(http.MethodPut)
}

func (h airlineHandler) AddAirline(w http.ResponseWriter, r *http.Request) {
	airline := new(entity.Airline)

	err := json.NewDecoder(r.Body).Decode(airline)
	if err != nil {
		resp, code := entity.HandleError(r.Context(), h.logger, entity.ErrorBadRequest)
		w.WriteHeader(code)
		w.Write(resp)
		return
	}

	airline, err = h.airlineService.AddAirline(r.Context(), airline)
	if err != nil {
		resp, code := entity.HandleError(r.Context(), h.logger, err)
		w.WriteHeader(code)
		w.Write(resp)
		return
	}

	resp, _ := json.Marshal(airline)
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}

func (h airlineHandler) DeleteAirline(w http.ResponseWriter, r *http.Request) {
	code := mux.Vars(r)["code"]

	err := h.airlineService.DeleteAirline(r.Context(), code)
	if err != nil {
		resp, code := entity.HandleError(r.Context(), h.logger, err)
		w.WriteHeader(code)
		w.Write(resp)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h airlineHandler) PutAirlineProviders(w http.ResponseWriter, r *http.Request) {

	airlineProviders := new(entity.AirlineProviders)

	err := json.NewDecoder(r.Body).Decode(airlineProviders)
	if err != nil {
		resp, code := entity.HandleError(r.Context(), h.logger, entity.ErrorBadRequest)
		w.WriteHeader(code)
		w.Write(resp)
		return
	}

	airlineProviders, err = h.airlineService.PutAirlineProviders(r.Context(), airlineProviders)
	if err != nil {
		resp, code := entity.HandleError(r.Context(), h.logger, err)
		w.WriteHeader(code)
		w.Write(resp)
		return
	}

	resp, _ := json.Marshal(airlineProviders)
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
