package api

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"time-tracker/entity"
)

type errorResponse struct {
	Error string `json:"error"`
}

func sendError(ctx context.Context, w http.ResponseWriter, err error) {
	l := ctx.Value(entity.CtxLogger{}).(*slog.Logger)

	errCode := http.StatusInternalServerError

	switch {
	case errors.Is(err, entity.ErrNotFound):
		errCode = http.StatusNotFound
	case errors.Is(err, entity.ErrBadRequest):
		errCode = http.StatusBadRequest
	case errors.Is(err, entity.ErrValidate):
		errCode = http.StatusBadRequest
	}

	w.WriteHeader(errCode)
	resp := errorResponse{Error: err.Error()}
	_, _ = w.Write([]byte(http.StatusText(errCode)))
	_ = sendJson(w, resp)

	l.Error("sendError", "error", err)
}

func sendJson(w http.ResponseWriter, body any) error {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		return err
	}
	return nil
}
