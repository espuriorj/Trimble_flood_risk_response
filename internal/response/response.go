package response

import (
    "encoding/json"
    "errors"
    "log"
    "net/http"

    apperrors "floodRiskManagementService/internal/errors"
)

type Response struct {
    Success bool        `json:"success"`
    Data    interface{} `json:"data,omitempty"`
    Error   string      `json:"error,omitempty"`
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)

    response := Response{
        Success: statusCode < 400,
        Data:    data,
    }

    if err := json.NewEncoder(w).Encode(response); err != nil {
        log.Printf("Failed to encode JSON response: %v", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    }
}

func Error(w http.ResponseWriter, err error) {
    var appErr *apperrors.AppError
    if errors.As(err, &appErr) {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(appErr.Code)

        response := Response{
            Success: false,
            Error:   appErr.Message,
        }

        if encodeErr := json.NewEncoder(w).Encode(response); encodeErr != nil {
            log.Printf("Failed to encode error response: %v", encodeErr)
        }
        return
    }

    JSON(w, http.StatusInternalServerError, nil)
}
