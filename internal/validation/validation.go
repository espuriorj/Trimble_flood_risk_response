package validation

import (
    "fmt"

    apperrors "floodRiskManagementService/internal/errors"
)

func ValidateCoordinates(lat, lng float64) error {
    if lat < -90 || lat > 90 {
        return apperrors.NewBadRequestError("Invalid latitude", apperrors.ErrInvalidLatitude)
    }
    if lng < -180 || lng > 180 {
        return apperrors.NewBadRequestError("Invalid longitude", apperrors.ErrInvalidLongitude)
    }
    return nil
}

func ValidateFloodReport(affectedArea, actions string) error {
    if affectedArea == "" {
        return apperrors.NewBadRequestError("Affected area is required", apperrors.ErrMissingAffectedArea)
    }
    if actions == "" {
        return apperrors.NewBadRequestError("Actions are required", apperrors.ErrMissingActions)
    }
    return nil
}

func ValidateID(id int) error {
    if id <= 0 {
        return apperrors.NewBadRequestError("Invalid ID", fmt.Errorf("ID must be positive, got %d", id))
    }
    return nil
}
