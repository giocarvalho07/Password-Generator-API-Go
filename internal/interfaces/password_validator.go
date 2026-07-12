package interfaces

import "password-generator/internal/domain"

type PasswordValidator interface {
	Validate(password string) (*domain.ValidationResult, error)
}