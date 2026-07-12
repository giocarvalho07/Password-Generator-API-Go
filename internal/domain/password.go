package domain

import (
	"strings"
	"time"
)

type Password struct {
	Value     string
	Length    int
	Generated time.Time
}

func NewPassword(value string) *Password {
	return &Password{
		Value:     value,
		Length:    len(value),
		Generated: time.Now().UTC(),
	}
}

func (p *Password) String() string {
	if p.Length <= 2 {
		return strings.Repeat("*", p.Length)
	}
	return string(p.Value[0]) + strings.Repeat("*", p.Length-2) + string(p.Value[p.Length-1])
}

func (p *Password) HasUppercase() bool {
	for _, c := range p.Value {
		if c >= 'A' && c <= 'Z' {
			return true
		}
	}
	return false
}

func (p *Password) HasLowercase() bool {
	for _, c := range p.Value {
		if c >= 'a' && c <= 'z' {
			return true
		}
	}
	return false
}

func (p *Password) HasNumbers() bool {
	for _, c := range p.Value {
		if c >= '0' && c <= '9' {
			return true
		}
	}
	return false
}

func (p *Password) HasSpecialChars() bool {
	specialChars := "!@#$%^&*()_+-=[]{}|;:',.<>?/`~"
	for _, c := range p.Value {
		if strings.ContainsRune(specialChars, c) {
			return true
		}
	}
	return false
}

func (p *Password) AlphabetSize() int {
	size := 0
	if p.HasUppercase() {
		size += 26
	}
	if p.HasLowercase() {
		size += 26
	}
	if p.HasNumbers() {
		size += 10
	}
	if p.HasSpecialChars() {
		size += 32
	}
	return size
}