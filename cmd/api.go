package main

import (
	"net/http"
	"time"

	"github.com/dimplesY/goose_test/internal/accounts"
	database "github.com/dimplesY/goose_test/internal/db"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
)

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("server is running"))
	})

	accountsService := accounts.NewAccountService(database.New(app.db))
	accountsHandler := accounts.NewAccountHandler(accountsService)

	r.Post("/login", accountsHandler.LoginByNameAndPassword)

	return r
}

func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	return srv.ListenAndServe()
}

type application struct {
	db *pgx.Conn
}
