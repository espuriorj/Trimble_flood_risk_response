package service

import "time"

type (
    FloodRisk struct {
        ID           int
        Latitude     float64
        Longitude    float64
        RiskLevel    string
        DischargeM3S float64
        AssessedAt   time.Time
    }

    FloodReport struct {
        ID           int
        AffectedArea string
        Actions      string
        ReportedAt   time.Time
    }

    FloodRiskRequest struct {
        Latitude  float64
        Longitude float64
    }

    FloodAPIData struct {
        LatestDischarge float64  `json:"latest_discharge"` // Most recent discharge value
        Timestamps      []string `json:"timestamps"`       // Time data from API
    }
)
