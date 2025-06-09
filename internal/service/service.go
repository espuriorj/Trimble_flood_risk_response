package service

import (
    "context"
    "encoding/json"
    "fmt"
    "net/http"
    "time"

    apperrors "floodRiskManagementService/internal/errors"
    "floodRiskManagementService/internal/repository"
    "floodRiskManagementService/internal/validation"
)

const (
    FloodAPITimeout = 10 * time.Second
    FloodAPIBaseURL = "https://flood-api.open-meteo.com/v1/flood"
)

type FloodService struct {
    repo       *repository.Repository
    httpClient *http.Client
}

func NewFloodService(repo *repository.Repository, client *http.Client) *FloodService {
    if client == nil {
        client = &http.Client{
            Timeout: FloodAPITimeout,
        }
    }

    return &FloodService{
        repo:       repo,
        httpClient: client,
    }
}

func (s *FloodService) AssessFloodRisk(ctx context.Context, req FloodRiskRequest) (*FloodRisk, error) {
    if err := validation.ValidateCoordinates(req.Latitude, req.Longitude); err != nil {
        return nil, err
    }

    apiData, err := s.fetchFloodData(ctx, req.Latitude, req.Longitude)
    if err != nil {
        return nil, apperrors.NewInternalServerError("Failed to fetch flood data", err)
    }

    riskLevel := s.assessRiskLevel(apiData.LatestDischarge)

    floodRisk := &repository.FloodRisk{
        Latitude:     req.Latitude,
        Longitude:    req.Longitude,
        RiskLevel:    riskLevel,
        DischargeM3S: apiData.LatestDischarge,
        AssessedAt:   time.Now(),
    }

    if err := s.repo.SaveRisk(ctx, floodRisk); err != nil {
        return nil, apperrors.NewInternalServerError("Failed to save flood risk assessment", err)
    }

    return (*FloodRisk)(floodRisk), nil
}

func (s *FloodService) GetAllFloodRisks(ctx context.Context) ([]FloodRisk, error) {
    risks, err := s.repo.FindAllRisks(ctx)
    if err != nil {
        return nil, apperrors.NewInternalServerError("Failed to retrieve flood risks", err)
    }

    floodRisks := make([]FloodRisk, 0, len(risks))
    for _, risk := range risks {
        floodRisks = append(floodRisks, FloodRisk(risk))
    }

    return floodRisks, nil
}

func (s *FloodService) GetFloodRiskByID(ctx context.Context, id int) (*FloodRisk, error) {
    if err := validation.ValidateID(id); err != nil {
        return nil, err
    }

    floodRisk, err := s.repo.FindRiskByID(ctx, id)
    if err != nil {
        return nil, apperrors.NewNotFoundError("Flood risk assessment not found", err)
    }

    return (*FloodRisk)(&floodRisk), nil
}

func (s *FloodService) AddFloodReport(ctx context.Context, report FloodReport) error {
    if err := validation.ValidateFloodReport(report.AffectedArea, report.Actions); err != nil {
        return err
    }

    report.ReportedAt = time.Now()

    if err := s.repo.SaveReport(ctx, (*repository.FloodReport)(&report)); err != nil {
        return apperrors.NewInternalServerError("Failed to save flood report", err)
    }

    return nil
}

// fetchFloodData retrieves flood data from external API
func (s *FloodService) fetchFloodData(ctx context.Context, lat, lng float64) (*FloodAPIData, error) {
    url := fmt.Sprintf(
        "%s?latitude=%f&longitude=%f&daily=river_discharge",
        FloodAPIBaseURL, lat, lng,
    )

    req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
    if err != nil {
        return nil, fmt.Errorf("failed to create request: %w", err)
    }

    resp, err := s.httpClient.Do(req)
    if err != nil {
        return nil, fmt.Errorf("failed to execute request: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
    }

    var apiResponse struct {
        Daily struct {
            Time           []string  `json:"time"`
            RiverDischarge []float64 `json:"river_discharge"`
        } `json:"daily"`
    }

    if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
        return nil, fmt.Errorf("failed to decode API response: %w", err)
    }

    if len(apiResponse.Daily.RiverDischarge) == 0 {
        return nil, apperrors.ErrNoDischargeData
    }

    return &FloodAPIData{
        LatestDischarge: apiResponse.Daily.RiverDischarge[len(apiResponse.Daily.RiverDischarge)-1],
        Timestamps:      apiResponse.Daily.Time,
    }, nil
}

// assessRiskLevel determines risk level based on discharge data
func (s *FloodService) assessRiskLevel(discharge float64) string {
    switch {
    case discharge > 1000:
        return "High"
    case discharge > 500:
        return "Medium"
    default:
        return "Low"
    }
}
