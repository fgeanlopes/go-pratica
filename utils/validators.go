package utils

import (
	"errors"
	"regexp"
	"strings"
)

// ValidateZipCode valida o formato do CEP (apenas números e opcionalmente um traço)
func ValidateZipCode(zipcode string) (string, error) {

	// Remove o traço, se presente, para retornar apenas os números
	clean := strings.NewReplacer("-", "", ".", "").Replace(zipcode)

	// Verifica se o CEP tem exatamente 8 dígitos após a limpeza
	if len(clean) != 8 {
		return "", errors.New("CEP deve conter exatamente 8 dígitos")
	}

	// Verifica se o CEP contém apenas números
	if matched, _ := regexp.MatchString(`^\d{8}$`, clean); !matched {
		return "", errors.New("CEP deve conter apenas números")
	}

	return clean, nil
}

func ValidateCpf(cpf string) (string, error) {
	// Remove os pontos e traços, se presentes, para retornar apenas os números
	clean := strings.NewReplacer("-", "", ".", "").Replace(cpf)

	// Verifica se o CPF tem exatamente 11 dígitos após a limpeza
	if len(clean) != 11 {
		return "", errors.New("CPF deve conter exatamente 11 dígitos")
	}

	// Verifica se o CPF contém apenas números
	if matched, _ := regexp.MatchString(`^\d{11}$`, clean); !matched {
		return "", errors.New("CPF deve conter apenas números")
	}

	return clean, nil
}
