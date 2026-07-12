package main

import (
	"fmt"
	"log"
	"os"
	"password-generator/internal/infrastructure/config"
	"password-generator/internal/infrastructure/handlers"
	"password-generator/internal/infrastructure/router"
	"password-generator/internal/infrastructure/services"
	"password-generator/internal/usecase"

	_ "password-generator/docs"

	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// @title Password Generator API
// @version 1.0
// @description API for generating and validating secure passwords
// @host localhost:8080
// @BasePath /
func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load config: %v\n", err)
		os.Exit(1)
	}

	generator := services.NewCryptoGenerator()
	validator := services.NewRuleValidator()

	generateUC := usecase.NewGeneratePasswordUseCase(generator, validator)
	validateUC := usecase.NewValidatePasswordUseCase(validator)

	handler := handlers.NewPasswordHandler(generateUC, validateUC)

	r := router.SetupRouter(handler)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Printf("Server starting on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to start server: %v\n", err)
		os.Exit(1)
	}
}