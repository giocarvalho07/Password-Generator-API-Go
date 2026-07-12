package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"password-generator/internal/infrastructure/services"
	"password-generator/internal/usecase"
	"testing"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)

	generator := services.NewCryptoGenerator()
	validator := services.NewRuleValidator()

	generateUC := usecase.NewGeneratePasswordUseCase(generator, validator)
	validateUC := usecase.NewValidatePasswordUseCase(validator)

	handler := NewPasswordHandler(generateUC, validateUC)

	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/health", handler.HealthCheck)

	api := r.Group("/api/v1")
	{
		password := api.Group("/password")
		{
			password.POST("/generate", handler.GeneratePassword)
			password.POST("/validate", handler.ValidatePassword)
		}
	}

	return r
}

func TestHealthCheck(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var response HealthResponse
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if response.Status != "healthy" {
		t.Errorf("expected status 'healthy', got '%s'", response.Status)
	}
	if response.Timestamp.IsZero() {
		t.Error("expected timestamp to be set")
	}
}

func TestGeneratePassword_Success(t *testing.T) {
	router := setupRouter()

	body := `{"length":12,"use_uppercase":true,"use_lowercase":true,"use_numbers":true,"use_symbols":true}`
	req, _ := http.NewRequest("POST", "/api/v1/password/generate", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var response GenerateResponse
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if len(response.Password) != 12 {
		t.Errorf("expected password length 12, got %d", len(response.Password))
	}
	if response.Validation == nil {
		t.Error("expected validation to be present")
	}
}

func TestGeneratePassword_ShortLength(t *testing.T) {
	router := setupRouter()

	body := `{"length":4,"use_uppercase":true}`
	req, _ := http.NewRequest("POST", "/api/v1/password/generate", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", w.Code)
	}
}

func TestGeneratePassword_InvalidJSON(t *testing.T) {
	router := setupRouter()

	body := `invalid json`
	req, _ := http.NewRequest("POST", "/api/v1/password/generate", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", w.Code)
	}
}

func TestValidatePassword_Success(t *testing.T) {
	router := setupRouter()

	body := `{"password":"Abc123!@#"}`
	req, _ := http.NewRequest("POST", "/api/v1/password/validate", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var response ValidationResponse
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if !response.IsValid {
		t.Error("expected password to be valid")
	}
}

func TestValidatePassword_InvalidPassword(t *testing.T) {
	router := setupRouter()

	body := `{"password":"abc"}`
	req, _ := http.NewRequest("POST", "/api/v1/password/validate", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnprocessableEntity {
		t.Errorf("expected status 422, got %d", w.Code)
	}
}

func TestValidatePassword_EmptyPassword(t *testing.T) {
	router := setupRouter()

	body := `{"password":""}`
	req, _ := http.NewRequest("POST", "/api/v1/password/validate", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", w.Code)
	}
}

func TestValidatePassword_MissingPassword(t *testing.T) {
	router := setupRouter()

	body := `{}`
	req, _ := http.NewRequest("POST", "/api/v1/password/validate", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", w.Code)
	}
}