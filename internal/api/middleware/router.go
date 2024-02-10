package middleware

import (
	"context"
	"fmt"
	uuid2 "github.com/google/uuid"
	"net/http"
	"runtime/debug"
	"test-smartway/internal/app/config"
	"test-smartway/internal/entity"
	"time"

	"go.uber.org/zap"
)

type Middleware struct {
	logger *zap.Logger
	cfg    *config.Config
}

func NewMiddleware(logger *zap.Logger, cfg *config.Config) *Middleware {
	return &Middleware{
		logger: logger,
		cfg:    cfg,
	}
}

func (m *Middleware) PanicRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()
		defer func() {
			if err := recover(); err != nil {

				resp := []byte("{\"error\": \"InternalServerError\"}")
				respCode := 500

				// Ответ клиенту
				w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				w.WriteHeader(respCode)
				w.Write(resp)

				m.logger.DPanic("Panic Recovery",
					zap.Any("RequestId", r.Context().Value(entity.RequestIDKey)),
					zap.String("LeadTime", fmt.Sprintf("%.3f", time.Duration(time.Now().UnixNano()-timeStart.UnixNano()).Seconds())),
					zap.String("RequestMethod", r.Method),
					zap.Any("LogicError", err),
					zap.String("URL", r.URL.RequestURI()),
					zap.Int32("ResponseCode", int32(respCode)),
					zap.String("ResponseBody", string(resp)),
					zap.String("RemoteAddr", r.RemoteAddr),
					zap.String("UserAgent", r.UserAgent()),
					zap.String("StackTrace", string(debug.Stack())),
				)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func (m *Middleware) ContentTypeJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")

			next.ServeHTTP(w, r)
		},
	)
}

func (m *Middleware) RequestId(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			uuid, _ := uuid2.NewUUID()
			r = r.WithContext(context.WithValue(r.Context(), entity.RequestIDKey, uuid.String()))

			next.ServeHTTP(w, r)
		},
	)
}

func (m *Middleware) Timeout(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx, cancel := context.WithTimeout(r.Context(), m.cfg.RequestTimeout)
			defer cancel()
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		},
	)
}

func (m *Middleware) DebugLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			timeStart := time.Now()

			next.ServeHTTP(w, r)

			m.logger.Debug("Request Logger",
				zap.Any("RequestId", r.Context().Value(entity.RequestIDKey)),
				zap.String("LeadTime", fmt.Sprintf("%.3f", time.Duration(time.Now().UnixNano()-timeStart.UnixNano()).Seconds())),
				zap.String("RequestMethod", r.Method),
				zap.String("URL", r.RequestURI),
				zap.String("IP", r.RemoteAddr),
			)
		},
	)
}
