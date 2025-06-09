package internal

import (
    "net/http"

    "floodRiskManagementService/internal/handlers"
    "floodRiskManagementService/internal/middleware"
    "floodRiskManagementService/internal/service"

    "github.com/gorilla/mux"
)

// SetupRoutes configures all application routes with middleware
func SetupRoutes(service *service.FloodService) *mux.Router {
    router := mux.NewRouter()

    router.Use(middleware.Logger)
    router.Use(middleware.Recovery)
    router.Use(middleware.CORS)

    h := handlers.NewFloodHandler(service)

    router.HandleFunc("/flood/risk", h.GetAllFloodRisks).Methods(http.MethodGet)
    router.HandleFunc("/flood/risk", h.SaveFloodRisk).Methods(http.MethodPost)
    router.HandleFunc("/flood/risk/{id:[0-9]+}", h.GetFloodRiskByID).Methods(http.MethodGet)
    router.HandleFunc("/flood/report", h.SaveFloodReport).Methods(http.MethodPost)

    return router
}
