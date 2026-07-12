package usecase

import (
	"errors"
	"password-generator/internal/domain"
	"password-generator/internal/interfaces"
)

type GeneratePasswordUseCase struct {
	generator interfaces.PasswordGenerator
	validator interfaces.PasswordValidator
}

func NewGeneratePasswordUseCase(
	generator interfaces.PasswordGenerator,
	validator interfaces.PasswordValidator,
) *GeneratePasswordUseCase {
	return &GeneratePasswordUseCase{
		generator: generator,
		validator: validator,
	}
}

func (uc *GeneratePasswordUseCase) Execute(config interfaces.GenerateConfig) (*domain.Password, *domain.ValidationResult, error) {
	if config.Length < 8 {
		return nil, nil, errors.New("length must be at least 8")
	}
	if config.Length > 64 {
		return nil, nil, errors.New("length must be at most 64")
	}
	if !config.UseUppercase && !config.UseLowercase && !config.UseNumbers && !config.UseSymbols {
		return nil, nil, errors.New("at least one character type must be enabled")
	}

	value, err := uc.generator.Generate(config)
	if err != nil {
		return nil, nil, err
	}

	password := domain.NewPassword(value)

	validation, err := uc.validator.Validate(value)
	if err != nil {
		return nil, nil, err
	}

	return password, validation, nil
}