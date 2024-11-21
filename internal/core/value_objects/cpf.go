package valueobjects

import (
	"errors"
	"regexp"
	"strconv"
)

type CPF struct {
	value string
}

func NewCPF(value string) (*CPF, error) {
	reg := regexp.MustCompile(`\D`)
	cpf := &CPF{value: reg.ReplaceAllString(value, "")}
	if !cpf.validate() {
		return nil, errors.New("Invalid cpf")
	}
	return cpf, nil
}

func (c *CPF) Value() string {
	return c.value
}


func (c *CPF) validate() bool {
	if c.value == "" { return false }
	if len(c.value) != 11 { return false }
	if c.allDigitsAreTheSame() { return false }
	digit1 := strconv.Itoa(c.calculateDigit(10))
	digit2 := strconv.Itoa(c.calculateDigit(11))
	return digit1 + digit2 == c.value[len(c.value)-2:]
}

func (c *CPF) allDigitsAreTheSame() bool {
	for i := 1; i < len(c.value); i++ {
		if c.value[1] != c.value[0] {
			return false
		}
	}
	return true
}

func (c *CPF) calculateDigit(factor int) int {
	total := 0;
	for i := 0; i < len(c.value); i++ {
		if factor > 1 {
			digit, _ := strconv.Atoi(string(c.value[i]))
			total += digit * factor
			factor--
		} 
	}
	rest := total % 11
	if rest < 2 {
		return 0
	}
	return 11 - rest
}
