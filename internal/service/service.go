package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// Convert преобразует Морзе в текст, текст в Морзе
func Convert(input string) (string, error) {
	input = strings.TrimSpace(input)

	if input == "" {
		return "", errors.New("пустая строка")
	}

	cleaned := strings.Join(strings.Fields(input), "")

	isMorse := true

	for _, symbol := range cleaned {
		if symbol != '.' && symbol != '-' {
			isMorse = false
			break
		}
	}

	if isMorse {
		return morse.ToText(input), nil
	}

	return morse.ToMorse(input), nil
}
