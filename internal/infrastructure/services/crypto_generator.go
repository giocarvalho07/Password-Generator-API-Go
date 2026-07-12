package services

import (
	"crypto/rand"
	"errors"
	"math/big"
	"password-generator/internal/interfaces"
)

type CryptoGenerator struct{}

func NewCryptoGenerator() *CryptoGenerator {
	return &CryptoGenerator{}
}

func (g *CryptoGenerator) Generate(config interfaces.GenerateConfig) (string, error) {
	charset := g.buildCharset(config)
	if len(charset) == 0 {
		return "", errors.New("no characters available for generation")
	}

	result := make([]byte, config.Length)
	for i := range result {
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result[i] = charset[idx.Int64()]
	}

	return string(result), nil
}

func (g *CryptoGenerator) buildCharset(config interfaces.GenerateConfig) []byte {
	var charset []byte

	if config.UseUppercase {
		for c := byte('A'); c <= 'Z'; c++ {
			charset = append(charset, c)
		}
	}
	if config.UseLowercase {
		for c := byte('a'); c <= 'z'; c++ {
			charset = append(charset, c)
		}
	}
	if config.UseNumbers {
		for c := byte('0'); c <= '9'; c++ {
			charset = append(charset, c)
		}
	}
	if config.UseSymbols {
		symbols := "!@#$%^&*()_+-=[]{}|;:',.<>?/`~"
		for i := 0; i < len(symbols); i++ {
			charset = append(charset, symbols[i])
		}
	}

	return charset
}