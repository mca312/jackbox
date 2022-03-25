package app

import "strings"

func IsValidEmail(email string) bool {
	if len(email) == 0 || !strings.Contains(email, "@") || (strings.Index(email, "@") != strings.LastIndex(email, "@")) {
		return false
	}

	return true
}
