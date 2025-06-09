package repository

import (
    "context"
    "fmt"

    "github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
    db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
    return &Repository{db: db}
}

func (r *Repository) SaveRisk(ctx context.Context, risk *FloodRisk) error {
    query := `INSERT INTO flood_risks (latitude, longitude, risk_level, discharge_m3s, assessed_at) 
        VALUES ($1, $2, $3, $4, $5) RETURNING id`

    err := r.db.QueryRow(
        ctx, query,
        risk.Latitude,
        risk.Longitude,
        risk.RiskLevel,
        risk.DischargeM3S,
        risk.AssessedAt,
    ).Scan(&risk.ID)

    return err
}

func (r *Repository) FindAllRisks(ctx context.Context) ([]FloodRisk, error) {
    query := `SELECT id, latitude, longitude, risk_level, discharge_m3s, assessed_at FROM flood_risks`
    rows, err := r.db.Query(ctx, query)
    if err != nil {
        return nil, fmt.Errorf("failed to query flood risks: %w", err)
    }
    defer rows.Close()

    var risks []FloodRisk
    for rows.Next() {
        var risk FloodRisk
        err := rows.Scan(
            &risk.ID, &risk.Latitude, &risk.Longitude, &risk.RiskLevel, &risk.DischargeM3S, &risk.AssessedAt,
        )
        if err != nil {
            return nil, fmt.Errorf("failed to scan flood risk row: %w", err)
        }
        risks = append(risks, risk)
    }
    return risks, nil
}

func (r *Repository) FindRiskByID(ctx context.Context, id int) (FloodRisk, error) {
    query := `SELECT id, latitude, longitude, risk_level, discharge_m3s, assessed_at FROM flood_risks WHERE id = $1`
    var risk FloodRisk
    err := r.db.QueryRow(ctx, query, id).Scan(
        &risk.ID, &risk.Latitude, &risk.Longitude, &risk.RiskLevel, &risk.DischargeM3S, &risk.AssessedAt,
    )
    if err != nil {
        return FloodRisk{}, fmt.Errorf("failed to query flood risk by ID: %w", err)
    }
    return risk, nil
}

func (r *Repository) SaveReport(ctx context.Context, report *FloodReport) error {
    query := `INSERT INTO flood_reports (affected_area, actions, reported_at) VALUES ($1, $2, $3) RETURNING id`

    err := r.db.QueryRow(
        ctx, query,
        report.AffectedArea,
        report.Actions,
        report.ReportedAt,
    ).Scan(&report.ID)

    return err
}
