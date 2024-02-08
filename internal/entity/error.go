package entity

import (
	"context"
	"encoding/json"
	"errors"
	chiM "github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
	"net/http"
	"test-smartway/internal/api/middleware"
)

var ErrorInternalServerMarshal, _ = json.Marshal(LogicError{ResponseMessage: "internal server error"})
var ErrorInternalServer = &LogicError{ResponseMessage: "internal server error", Code: 500}
var ErrorNotFound = &LogicError{ResponseMessage: "not found", Code: 404}
var ErrorBadRequestMarshal, _ = json.Marshal(LogicError{ResponseMessage: "bad request"})
var ErrorBadRequest = &LogicError{ResponseMessage: "bad request", Code: 400}

type LogicError struct {
	ResponseMessage string `json:"error"`
	Err             error  `json:"-"`
	Code            int    `json:"-"`
}

func NewLogicError(err error, responseMessage string, code int) *LogicError {
	return &LogicError{
		ResponseMessage: responseMessage,
		Err:             err,
		Code:            code,
	}
}

func (e *LogicError) Error() string {
	if e == nil || e.Err == nil {
		return ""
	}

	return e.Err.Error()
}

func (e *LogicError) JsonMarshal() []byte {
	if e == nil || len(e.ResponseMessage) == 0 {
		return nil
	}

	b, _ := json.Marshal(e)
	return b
}

func HandleError(ctx context.Context, logger *zap.Logger, err error) ([]byte, int) {
	var logicErr *LogicError
	if errors.As(err, &logicErr) {
		logger.Error(logicErr.Error(), zap.String("RequestId", chiM.GetReqID(ctx)), zap.Int("ResponseCode", logicErr.Code))
		return logicErr.JsonMarshal(), logicErr.Code
	}

	logger.Error(err.Error(), zap.Any("RequestId", ctx.Value(middleware.RequestIDKey)), zap.Int("ResponseCode", http.StatusBadRequest))

	return ErrorBadRequestMarshal, http.StatusBadRequest
}
