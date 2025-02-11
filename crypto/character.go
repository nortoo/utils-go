package crypto

import (
	"fmt"
	"strings"

	"github.com/nortoo/utils-go/validation"
)

// Email returns encrypted email address string
// e.g. t***@***.com
func Email(email string) string {
	if !validation.IsValidEmail(email) {
		return ""
	}

	emailSp := strings.Split(email, ".")
	return fmt.Sprintf("%s***@***.%s", email[:1], emailSp[len(emailSp)-1])
}
