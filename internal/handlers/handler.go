package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"

    apperrors "floodRiskManagementService/internal/errors"
    "floodRiskManagementService/internal/response"
    "floodRiskManagementService/internal/service"
    "floodRiskManagementService/internal/validation"

    "github.com/gorilla/mux"
)

type FloodHandler struct {
    service *service.FloodService
}

func NewFloodHandler(service *service.FloodService) *FloodHandler {
    return &FloodHandler{
        service: service,
    }
}

func (h *FloodHandler) SaveFloodRisk(w http.ResponseWriter, r *http.Request) {
    var req floodRiskRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        response.Error(w, apperrors.NewBadRequestError("Invalid JSON payload", err))
        return
    }

    if err := validation.ValidateCoordinates(req.Latitude, req.Longitude); err != nil {
        response.Error(w, err)
        return
    }

    floodRisk, err := h.service.AssessFloodRisk(
        r.Context(), service.FloodRiskRequest{
            Latitude:  req.Latitude,
            Longitude: req.Longitude,
        },
    )
    if err != nil {
        response.Error(w, err)
        return
    }

    response.JSON(w, http.StatusOK, floodRisk)
}

func (h *FloodHandler) GetAllFloodRisks(w http.ResponseWriter, r *http.Request) {
    risks, err := h.service.GetAllFloodRisks(r.Context())
    if err != nil {
        response.Error(w, err)
        return
    }

    response.JSON(w, http.StatusOK, risks)
}

func (h *FloodHandler) GetFloodRiskByID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    idStr := vars["id"]

    id, err := strconv.Atoi(idStr)
    if err != nil {
        response.Error(w, apperrors.NewBadRequestError("Invalid ID format", err))
        return
    }

    if err := validation.ValidateID(id); err != nil {
        response.Error(w, err)
        return
    }

    floodRisk, err := h.service.GetFloodRiskByID(r.Context(), id)
    if err != nil {
        response.Error(w, err)
        return
    }

    response.JSON(w, http.StatusOK, floodRisk)
}

func (h *FloodHandler) SaveFloodReport(w http.ResponseWriter, r *http.Request) {
    var req floodReport
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        response.Error(w, apperrors.NewBadRequestError("Invalid JSON payload", err))
        return
    }

    if err := validation.ValidateFloodReport(req.AffectedArea, req.Actions); err != nil {
        response.Error(w, err)
        return
    }

    report := service.FloodReport{
        AffectedArea: req.AffectedArea,
        Actions:      req.Actions,
    }

    if err := h.service.AddFloodReport(r.Context(), report); err != nil {
        response.Error(w, err)
        return
    }

    response.JSON(
        w, http.StatusCreated, map[string]string{
            "message": "Flood report added successfully",
            "area":    req.AffectedArea,
        },
    )
}
