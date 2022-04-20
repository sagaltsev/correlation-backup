package correlation

import (
	"net/http"

	"github.com/google/uuid"
)

const (
	// ID header key
	ID = "X-Correlation-Id"
	// UserID header key
	UserID = "X-User-Correlation-Id"
)

// DecorateRequest decorates http request with provided correlation ID and user correlation ID
func DecorateRequest(r http.Request, correlationID string, userCorrelationID string) http.Request {
	if correlationID == "" {
		correlationID = GenerateID()
	}

	if userCorrelationID == "" {
		userCorrelationID = GenerateID()
	}

	r.Header.Set(ID, correlationID)
	r.Header.Set(UserID, userCorrelationID)

	return r
}

// Middleware adds correlation ID and user correlation ID to handler
func Middleware(next http.Handler) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		req := DecorateRequest(*r, r.Header.Get(ID), r.Header.Get(UserID))

		rw.Header().Set(ID, req.Header.Get(ID))
		rw.Header().Set(UserID, req.Header.Get(UserID))

		next.ServeHTTP(rw, &req)
	}
}

// GenerateID generates random ID
func GenerateID() string {
	return uuid.NewString()
}
