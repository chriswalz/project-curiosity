package validate

import "regexp"

func IsEmailAddress(email string) bool {
	m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, email)
	return m
}

// Minimum 8 characters at least 1 Alphabet and 1 Number:
func IsPassword(password string) bool {
	m, _ := regexp.MatchString(`^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,}$`, password)
	return m
}
