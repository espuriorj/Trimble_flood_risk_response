package internal

import (
    "net/http"

    "floodRiskManagementService/db"
    "floodRiskManagementService/internal/repository"
    "floodRiskManagementService/internal/service"
)

type Server struct {
    mux *http.ServeMux
}

func NewServer() (*Server, error) {
    s := &Server{
        mux: http.NewServeMux(),
    }

    repo := repository.NewRepository(db.DB)
    svc := service.NewFloodService(repo, &http.Client{})
    handler := SetupRoutes(svc)
    s.mux.Handle("/", handler)

    return s, nil
}

func (s *Server) Start(addr string) error {
    return http.ListenAndServe(addr, s.mux)
}
