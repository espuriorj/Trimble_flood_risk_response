package repository

import "time"

type FloodRisk struct {
    ID           int       `db:"id"`
    Latitude     float64   `db:"latitude"`
    Longitude    float64   `db:"longitude"`
    RiskLevel    string    `db:"risk_level"`
    DischargeM3S float64   `db:"discharge_m3s"`
    AssessedAt   time.Time `db:"assessed_at"`
}

type FloodReport struct {
    ID           int       `db:"id"`
    AffectedArea string    `db:"affected_area"`
    Actions      string    `db:"actions"`
    ReportedAt   time.Time `db:"reported_at"`
}
