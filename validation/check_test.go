package validation

import (
	"testing"
)

func TestIsValidEmail(t *testing.T) {
	emails := map[string]bool{
		// ✅ Valid Email Addresses
		"user@example.com":             true,
		"john.doe@mail.co.uk":          true,
		"alice+test@sub.domain.org":    true,
		"my-email@company.net":         true,
		"first_last@edu.edu":           true,
		"a123@host.io":                 true,
		"user@sub.subdomain.com":       true,
		"valid123@domain.travel":       true,
		"email@hyphen-domain.com":      true,
		"normal.email@big.biz":         true,
		"email@multi.level.domain":     true,
		"firstname.lastname@work.pro":  true,
		"special_chars@a-b.org":        true,
		"valid_email@one.two.org":      true,
		"1234567890@numbers.com":       true,
		"name@localserver.local":       true,
		"user@[123.45.67.89]":          true, // IP-based email
		"test@xn--example.com":         true, // IDN (Punycode)
		"test@国际化域名.cn":                true, // Unicode domain
		"test@company.museum":          true,
		"username@domain.info":         true,
		"me+regex@test.email":          true,
		"someone@long-tld.photography": true,
		"contact@legal.lawyer":         true,
		"example@sub.example.travel":   true,

		// ❌ Invalid Email Addresses
		".user@example.com":      false, // Starts with a dot
		"user.@example.com":      false, // Ends with a dot
		"user@.example.com":      false, // Domain starts with a dot
		"user@example..com":      false, // Consecutive dots in domain
		"user@-example.com":      false, // Domain starts with a hyphen
		"user@example-.com":      false, // Domain ends with a hyphen
		"user@example,com":       false, // Comma instead of dot
		"user@domain..com":       false, // Double dots in domain
		"user@domain":            false, // Missing TLD
		"user@.com":              false, // Invalid domain
		"@missinglocal.com":      false, // No local part
		"user@missingdotcom":     false, // No dot in domain
		"user@invalid_char$.com": false, // Special character `$` in domain
		"user space@example.com": false, // Whitespace not allowed
		"user@ex!ample.com":      false, // `!` not allowed in domain
		"invalid@com.":           false, // Ends with a dot
		"user@sub_domain.com":    false, // Underscore `_` in domain
		"a@b":                    false, // Too short, missing valid TLD
		"user@localhost":         false, // No dot in domain
		"user@domain.toolongtldoooooooooooooooooooooooooooooooooooooooooooooooooooxyz": false, // TLD too long (max 63 chars)
		" @example.com":       false, // Empty local part
		"user@exam..ple.com":  false, // Double dots in domain
		"user@domain.123":     false, // Numeric-only TLD not allowed
		"username@.1234":      false, // Starts with a dot, numeric TLD
		"user@domain..dotcom": false, // Double dots in domain
	}

	for e, i := range emails {
		if IsValidEmail(e) != i {
			t.Logf("%s is not covered.", e)
		}
	}
}
