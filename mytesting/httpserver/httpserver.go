package httpserver

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	dbCli DBClient
}

func (s *Server) Handler() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/ping", s.Ping).Methods("GET")

	return r
}

func (s *Server) Ping(w http.ResponseWriter, r *http.Request) {
	if err := s.dbCli.Ping(); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(200)
}

type DBClient interface {
	Ping() error
}

type mem struct{}

func (*mem) Ping() error { return nil }
