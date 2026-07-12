package interfaces

type GenerateConfig struct {
	Length       int  `json:"length"`
	UseUppercase bool `json:"use_uppercase"`
	UseLowercase bool `json:"use_lowercase"`
	UseNumbers   bool `json:"use_numbers"`
	UseSymbols   bool `json:"use_symbols"`
}

type PasswordGenerator interface {
	Generate(config GenerateConfig) (string, error)
}