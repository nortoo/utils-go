package validation

import (
	"regexp"
	"strings"
)

// IsValidEmail returns whether the given email address is legal.
// note: This regex cannot cover:
//  1. IP-Based domain;
//  2. Unicode-Based domain.
func IsValidEmail(email string) bool {
	strings.ToLower(email)

	pattern := `^[a-zA-Z0-9][a-zA-Z0-9._%+\-]*[a-zA-Z0-9]+@([a-zA-Z0-9]+([a-zA-Z0-9-]*[a-zA-Z0-9]+)?\.)+([a-zA-Z]{2,63})$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}
