package middleware

import (
    "log"
    "net/http"
    "runtime/debug"
    "time"
)

func Logger(next http.Handler) http.Handler {
    return http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {
            start := time.Now()
            log.Printf("Started %s %s", r.Method, r.URL.Path)

            next.ServeHTTP(w, r)

            log.Printf("Completed %s %s in %v", r.Method, r.URL.Path, time.Since(start))
        },
    )
}

func Recovery(next http.Handler) http.Handler {
    return http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {
            defer func() {
                if err := recover(); err != nil {
                    log.Printf("Panic recovered: %v\n%s", err, debug.Stack())
                    http.Error(w, "Internal Server Error", http.StatusInternalServerError)
                }
            }()
            next.ServeHTTP(w, r)
        },
    )
}

func CORS(next http.Handler) http.Handler {
    return http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {
            w.Header().Set("Access-Control-Allow-Origin", "*")
            w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
            w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

            if r.Method == "OPTIONS" {
                w.WriteHeader(http.StatusOK)
                return
            }

            next.ServeHTTP(w, r)
        },
    )
}
