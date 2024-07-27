package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"
	"time-tracker/api"
	"time-tracker/config"
	"time-tracker/service"
	"time-tracker/storage"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, cfg.Postgres)
	if err != nil {
		logger.Error("connect to database", "error", err)
		return
	}
	logger.Info("connected to postgres -ok")
	err = conn.Ping(ctx)
	if err != nil {
		logger.Error("ping to database", "error", err)
	}

	defer conn.Close(ctx)

	ur := storage.NewUserRepo(conn)
	tr := storage.NewTaskRepo(conn)

	us := service.NewUserService(ur)
	ts := service.NewTaskService(tr)

	uh := api.NewUserHandler(us)
	th := api.NewTaskHandler(ts)
	mw := api.NewMiddleware(logger)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /users", uh.CreateUser)
	mux.HandleFunc("GET /users", uh.Users)
	mux.HandleFunc("DELETE /users", uh.DeleteUser)
	mux.HandleFunc("PUT /users", uh.UpdateUser)
	mux.HandleFunc("GET /info", uh.UserByPassport)

	mux.HandleFunc("POST /tasks", th.CreateTask)
	mux.HandleFunc("PUT /tasks", th.UpdateStatus)
	mux.HandleFunc("GET /labor_costs", th.TaskTimeByUserID)
	server := http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.Port),
		Handler:           mw.Logging(mux),
		ReadTimeout:       time.Second,
		ReadHeaderTimeout: time.Second,
		WriteTimeout:      time.Second,
	}
	err = server.ListenAndServe()
	if err != nil {
		logger.Error("listen and serve", "error", err)
		return
	}
	logger.Info("listen and serve - ok")
}
