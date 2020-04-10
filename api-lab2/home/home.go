package home

import (
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"time"
)

const message = "Hello Github!"

type Handlers struct {
	logger   *log.Logger
	database *sqlx.DB
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {

	// TODO Database
	// h.database.ExecContext(r.Context(), "SELECT ....")

	w.Header().Set("Content-Type", "text/plain;chartset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}

func (h *Handlers) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer h.logger.Printf("request processed in %s\n", time.Now().Sub(startTime))
		next(w, r)
	}
}

func (h *Handlers) SetUpRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", h.Logger(h.Home))
}

func NewHandlers(logger *log.Logger, database *sqlx.DB) *Handlers {
	return &Handlers{
		logger:   logger,
		database: database,
	}
}
