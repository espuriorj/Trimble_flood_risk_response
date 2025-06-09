package main

import (
    "fmt"
    "log"

    "floodRiskManagementService/db"
    "floodRiskManagementService/internal"
)

func main() {
    db.Init()

    server, err := internal.NewServer()
    if err != nil {
        log.Fatalf("Failed to initialize server: %v", err)
    }

    fmt.Println("Server running on http://localhost:8080")
    if err := server.Start(":8080"); err != nil {
        log.Fatalf("Server failed: %v", err)
    }
}
