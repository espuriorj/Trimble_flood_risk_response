package handlers

import "time"

type (
    floodRiskRequest struct {
        Latitude  float64 `json:"latitude" validate:"required,min=-90,max=90"`
        Longitude float64 `json:"longitude" validate:"required,min=-180,max=180"`
    }

    floodReport struct {
        ID           int       `json:"id"`
        AffectedArea string    `json:"affected_area"`
        Actions      string    `json:"actions"`
        ReportedAt   time.Time `json:"reported_at"`
    }
)
