package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"test-smartway/internal/api/middleware"
	"test-smartway/internal/entity"
)

type schemaHandler struct {
	schemaService entity.ISchemaService
	logger        *zap.Logger
}

func RegisterSchemaHandlers(router *mux.Router, service entity.ISchemaService, logger *zap.Logger, mw *middleware.Middleware) {
	handler := new(schemaHandler)
	handler.schemaService = service
	handler.logger = logger

	r := router.PathPrefix("/schemas").Subrouter()

	r.Use(mw.PanicRecovery)
	r.Use(mw.RequestId)
	r.Use(mw.ContentTypeJSON)
	r.Use(mw.DebugLogger)

	r.HandleFunc("", handler.AddSchema).Methods(http.MethodPost)
	r.HandleFunc("/{name}", handler.GetSchema).Methods(http.MethodGet)
	r.HandleFunc("", handler.PatchSchema).Methods(http.MethodPatch)
	r.HandleFunc("/{id}", handler.DeleteSchema).Methods(http.MethodDelete)
}

func (h schemaHandler) AddSchema(w http.ResponseWriter, r *http.Request) {
	schema := new(entity.Schema)

	err := json.NewDecoder(r.Body).Decode(schema)
	if err != nil {
		resp, code := entity.HandleError(r.Context(), h.logger, entity.ErrorBadRequest)
		w.WriteHeader(code)
		w.Write(resp)
		return
	}

	schema, err = h.schemaService.AddSchema(r.Context(), schema)
	if err != nil {
		resp, code := entity.HandleError(r.Context(), h.logger, err)
		w.WriteHeader(code)
		w.Write(resp)
		return
	}

	resp, _ := json.Marshal(schema)
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}

func (h schemaHandler) GetSchema(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]

	schema, err := h.schemaService.GetSchema(r.Context(), name)
	if err != nil {
		resp, code := entity.HandleError(r.Context(), h.logger, err)
		w.WriteHeader(code)
		w.Write(resp)
		return
	} else if schema.Id == 0 {
		resp, code := entity.HandleError(r.Context(), h.logger, entity.ErrorNotFound)
		w.WriteHeader(code)
		w.Write(resp)
		return
	}

	resp, _ := json.Marshal(schema)
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (h schemaHandler) PatchSchema(w http.ResponseWriter, r *http.Request) {
	schema := new(entity.Schema)

	err := json.NewDecoder(r.Body).Decode(schema)
	if err != nil || schema.Id == 0 {
		resp, code := entity.HandleError(r.Context(), h.logger, entity.ErrorBadRequest)
		w.WriteHeader(code)
		w.Write(resp)
		return
	}

	schema, err = h.schemaService.PatchSchema(r.Context(), schema)
	if err != nil {
		resp, code := entity.HandleError(r.Context(), h.logger, err)
		w.WriteHeader(code)
		w.Write(resp)
		return
	} else if schema.Id == 0 {
		resp, code := entity.HandleError(r.Context(), h.logger, entity.ErrorNotFound)
		w.WriteHeader(code)
		w.Write(resp)
		return
	}

	resp, _ := json.Marshal(schema)
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (h schemaHandler) DeleteSchema(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := h.schemaService.DeleteSchema(r.Context(), id)
	if err != nil {
		resp, code := entity.HandleError(r.Context(), h.logger, err)
		w.WriteHeader(code)
		w.Write(resp)
		return
	}

	w.WriteHeader(http.StatusOK)
}
