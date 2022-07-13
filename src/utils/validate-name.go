package utils

import "regexp"

func ValidateName(name string) bool {
	var regexName = regexp.MustCompile(`^[a-zA-Z .\-']+$`)
	return regexName.MatchString(name)
}
