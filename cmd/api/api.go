package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/DNature/ecom/service/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()

	apiRouter := http.NewServeMux()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(apiRouter)

	// add api subrouter with prefix to main router
	router.Handle("/api/v1/", http.StripPrefix("/api/v1", apiRouter))

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
