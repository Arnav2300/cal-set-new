package utils

import (
	"regexp"
	"unicode"
)

func IsValidEmail(email string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}

func IsValidUsername(username string) bool {
	regex := `^[a-zA-Z0-9]{3,20}$` // 3-20 alphanumeric characters
	re := regexp.MustCompile(regex)
	return re.MatchString(username)
}

func IsValidPassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	var hasUpper, hasLower, hasNumber, hasSpecial bool
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasUpper && hasLower && hasNumber && hasSpecial
}
