package models

import (
	"errors"
	"fmt"
	"strings"
	"time"
	"unicode"
)

func EmptyMesage(caption string) error {
	return errors.New(fmt.Sprintf("%s is empty", caption))
}

func EmptyString(s string, caption string) error {
	if strings.Trim(s, " ") == "" {
		return errors.New(fmt.Sprintf("%s: could not be empty", caption))
	}

	return nil
}

func DateInvalid(s string, caption string) error {
	_, err := time.Parse("2006-01-02", s)

	if err != nil {
		return errors.New(fmt.Sprintf("%s:  format is invalid (yyyy-mm-dd). %s", caption, err))
	}

	return nil
}

func DateTimeInvalid(s string, caption string) error {
	_, err := time.Parse("2006-01-02 15:04:05", s)

	if err != nil {
		return errors.New(fmt.Sprintf("%s: format is invalid (yyyy-mm-dd hh:ii:ss). %s", caption, err))
	}

	return nil
}

func PasswordInvalid(s string, caption string) error {
	hasUpper := false
	hasLower := false
	hasNumber := false
	hasEspecialCharacter := false
	for _, r := range s {
		if unicode.IsUpper(r) {
			hasUpper = true
		}
		if unicode.IsLower(r) {
			hasLower = true
		}
		if r == '.' || r == '_' {
			hasEspecialCharacter = true
		}
		if unicode.IsNumber(r) {
			hasNumber = true
		}
	}

	if len(s) < 8 || !hasUpper || !hasLower || !hasNumber || !hasEspecialCharacter {
		return errors.New(fmt.Sprintf("%s: %s", caption, "must have at least 8 characters, one uppercase letter, one lowercase letter, one number and one special character . _"))
	}

	return nil
}
