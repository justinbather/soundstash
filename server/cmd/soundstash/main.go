package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/justinbather/soundstash/server/internal/adapter/downloader"
	"github.com/justinbather/soundstash/server/internal/adapter/handler"
	"github.com/justinbather/soundstash/server/logger"
	"go.uber.org/zap"
)

func main() {
	log := logger.NewLogger(false)
	defer log.Sync()

	r := chi.NewRouter()

	r.Use(loggingMiddleware(log))

	downloader := downloader.New()

	handler := handler.New(downloader)

	r.Get("/health", handler.HealthCheck)
	r.Get("/download", handler.Download)

	log.Info("Server starting", zap.Int("port", 8080))

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Error("Error starting server", zap.Error(err))
	}

}

func loggingMiddleware(logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// Wrap the ResponseWriter to capture status code
			ww := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

			next.ServeHTTP(ww, r)

			duration := time.Since(start)

			logger.Info("incoming request",
				zap.String("method", r.Method),
				zap.String("path", r.URL.Path),
				zap.Int("status", ww.statusCode),
				zap.Duration("duration", duration),
				zap.String("remote", r.RemoteAddr),
				zap.String("user_agent", r.UserAgent()),
			)
		})
	}
}

// responseWriter captures status codes
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
