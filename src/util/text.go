package util

import "regexp"

const (
	usernamePattern = `^[\p{L}\p{N}_.,;:'"-]{4,20}$`
	emailPattern    = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	titlePattern    = `[^a-zA-Z0-9\s-]|^\s+|\s+$`
)

func ValidateUsername(username string) bool {
	re := regexp.MustCompile(usernamePattern)
	return re.MatchString(username)
}

func ValidateEmail(email string) bool {
	re := regexp.MustCompile(emailPattern)
	return !re.MatchString(email)
}

func ValidatePassword(password string) bool {
	reDigit := regexp.MustCompile(`[0-9]`)
	reLetter := regexp.MustCompile(`[A-Za-z]`)
	return reDigit.MatchString(password) && reLetter.MatchString(password) && len(password) >= 8 && len(password) <= 16
}

func ValidateTitle(title string) bool {
	re := regexp.MustCompile(titlePattern)
	return !re.MatchString(title)
}
