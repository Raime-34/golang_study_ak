package main

import "regexp"

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`.+@example\.com$`)
	return re.Match([]byte(email))
}
