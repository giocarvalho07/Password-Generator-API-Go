package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port            string `mapstructure:"PORT"`
	MinLength       int    `mapstructure:"MIN_LENGTH"`
	MaxLength       int    `mapstructure:"MAX_LENGTH"`
	MinSpecialChars int    `mapstructure:"MIN_SPECIAL_CHARS"`
	MinEntropy      int    `mapstructure:"MIN_ENTROPY"`
	MaxConsecutive  int    `mapstructure:"MAX_CONSECUTIVE"`
}

func LoadConfig() (*Config, error) {
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("MIN_LENGTH", 8)
	viper.SetDefault("MAX_LENGTH", 64)
	viper.SetDefault("MIN_SPECIAL_CHARS", 1)
	viper.SetDefault("MIN_ENTROPY", 40)
	viper.SetDefault("MAX_CONSECUTIVE", 2)

	viper.AutomaticEnv()

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}