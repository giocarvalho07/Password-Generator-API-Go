package handlers

import (
	"net/http"
	"password-generator/internal/interfaces"
	"password-generator/internal/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

type PasswordHandler struct {
	generateUC *usecase.GeneratePasswordUseCase
	validateUC *usecase.ValidatePasswordUseCase
}

func NewPasswordHandler(
	generateUC *usecase.GeneratePasswordUseCase,
	validateUC *usecase.ValidatePasswordUseCase,
) *PasswordHandler {
	return &PasswordHandler{
		generateUC: generateUC,
		validateUC: validateUC,
	}
}

// GenerateRequest represents the request body for password generation
type GenerateRequest struct {
	Length       int  `json:"length" example:"16"`
	UseUppercase bool `json:"use_uppercase" example:"true"`
	UseLowercase bool `json:"use_lowercase" example:"true"`
	UseNumbers   bool `json:"use_numbers" example:"true"`
	UseSymbols   bool `json:"use_symbols" example:"true"`
}

// GenerateResponse represents the response for password generation
type GenerateResponse struct {
	Password   string              `json:"password" example:"xH7@kL9#mP2&"`
	Validation *ValidationResponse `json:"validation"`
	Generated  time.Time           `json:"generated" example:"2026-07-11T14:30:00Z"`
}

// ValidateRequest represents the request body for password validation
type ValidateRequest struct {
	Password string `json:"password" binding:"required" example:"xH7@kL9#mP2&"`
}

// ValidationResponse represents the response for password validation
type ValidationResponse struct {
	IsValid  bool     `json:"is_valid" example:"true"`
	Errors   []string `json:"errors"`
	Strength string   `json:"strength" example:"strong"`
	Entropy  float64  `json:"entropy" example:"72.54"`
}

// HealthResponse represents the health check response
type HealthResponse struct {
	Status    string    `json:"status" example:"healthy"`
	Timestamp time.Time `json:"timestamp" example:"2026-07-11T14:30:00Z"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error" example:"invalid request"`
}

// GeneratePassword godoc
// @Summary Generate a new password
// @Description Generate a secure password with configurable options
// @Tags password
// @Accept json
// @Produce json
// @Param request body GenerateRequest true "Password generation options"
// @Success 200 {object} GenerateResponse
// @Failure 400 {object} ErrorResponse
// @Router /api/v1/password/generate [post]
func (h *PasswordHandler) GeneratePassword(c *gin.Context) {
	var req GenerateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	config := interfaces.GenerateConfig{
		Length:       req.Length,
		UseUppercase: req.UseUppercase,
		UseLowercase: req.UseLowercase,
		UseNumbers:   req.UseNumbers,
		UseSymbols:   req.UseSymbols,
	}

	password, validation, err := h.generateUC.Execute(config)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, GenerateResponse{
		Password: password.Value,
		Validation: &ValidationResponse{
			IsValid:  validation.IsValid,
			Errors:   validation.Errors,
			Strength: validation.Strength,
			Entropy:  validation.Entropy,
		},
		Generated: password.Generated,
	})
}

// ValidatePassword godoc
// @Summary Validate a password
// @Description Validate password strength and complexity
// @Tags password
// @Accept json
// @Produce json
// @Param request body ValidateRequest true "Password to validate"
// @Success 200 {object} ValidationResponse
// @Failure 400 {object} ErrorResponse
// @Failure 422 {object} ValidationResponse
// @Router /api/v1/password/validate [post]
func (h *PasswordHandler) ValidatePassword(c *gin.Context) {
	var req ValidateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	result, err := h.validateUC.Execute(req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	if !result.IsValid {
		c.JSON(http.StatusUnprocessableEntity, ValidationResponse{
			IsValid:  result.IsValid,
			Errors:   result.Errors,
			Strength: result.Strength,
			Entropy:  result.Entropy,
		})
		return
	}

	c.JSON(http.StatusOK, ValidationResponse{
		IsValid:  result.IsValid,
		Errors:   result.Errors,
		Strength: result.Strength,
		Entropy:  result.Entropy,
	})
}

// HealthCheck godoc
// @Summary Health check
// @Description Check if the API is running
// @Tags health
// @Produce json
// @Success 200 {object} HealthResponse
// @Router /health [get]
func (h *PasswordHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now().UTC(),
	})
}