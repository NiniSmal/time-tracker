package api

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"time-tracker/entity"
)

func sendError(ctx context.Context, w http.ResponseWriter, err error) {
	l, ok := ctx.Value(entity.CtxLogger{}).(*slog.Logger)
	if !ok {
		panic("no login in context")
	}

	errCode := http.StatusInternalServerError

	switch {
	case errors.Is(err, entity.ErrNotFound):
		errCode = http.StatusNotFound
	case errors.Is(err, entity.ErrBadRequest):
		errCode = http.StatusBadRequest
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errCode)
	_, _ = w.Write([]byte(http.StatusText(errCode)))

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
