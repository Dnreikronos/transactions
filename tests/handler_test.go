package handlers_tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Dnreikronos/transactions/handlers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction_Validation(t *testing.T) {
	router := gin.Default()
	router.POST("/transactions", handlers.CreateTransaction)

	tests := []struct {
		name       string
		body       string
		statusCode int
	}{
		{
			name:       "Valid input",
			body:       `{"description": "Test transaction", "value": 50.99, "date": "2025-01-02"}`,
			statusCode: http.StatusAccepted,
		},
		{
			name:       "Invalid value",
			body:       `{"description": "Test transaction", "value": -10.00, "date": "2025-01-02"}`,
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "Missing description",
			body:       `{"value": 50.99, "date": "2025-01-02"}`,
			statusCode: http.StatusBadRequest,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/transactions", bytes.NewBufferString(tc.body))
			req.Header.Set("Content-Type", "application/json")
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)
			assert.Equal(t, tc.statusCode, resp.Code)
		})
	}
}
