package api

import (
	"net/http"
	"github.com/EAS-Kalem/moto_go/storage"
	"github.com/EAS-Kalem/moto_go/types"
	"github.com/EAS-Kalem/moto_go/api"
);

type Server struct {
	listenAddr string
	storage storage.Storage
}

func NewServer (listenAddr string ) *Server {
	return &Server {
		listenAddr: listenAddr,
		store: store,
	}
}


func (s *Server) Start() error {
	http.HandleFunc("/user", s.handleGetUserByID)
	return http.ListenAndServe(s.listenAddr, nil)

}

func (s *Server) handleGetUserByID(w http.ResponseWriter, r *http.Request){
	user := s.store.Get(10)
	json.NewEncoder(w).Encode(user)
}