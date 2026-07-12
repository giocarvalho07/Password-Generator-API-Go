package usecase

import (
	"errors"
	"password-generator/internal/domain"
	"password-generator/internal/interfaces"
)

type ValidatePasswordUseCase struct {
	validator interfaces.PasswordValidator
}

func NewValidatePasswordUseCase(validator interfaces.PasswordValidator) *ValidatePasswordUseCase {
	return &ValidatePasswordUseCase{
		validator: validator,
	}
}

func (uc *ValidatePasswordUseCase) Execute(password string) (*domain.ValidationResult, error) {
	if password == "" {
		return nil, errors.New("password cannot be empty")
	}

	result, err := uc.validator.Validate(password)
	if err != nil {
		return nil, err
	}

	return result, nil
}