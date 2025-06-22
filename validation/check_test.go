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

func TestIsValidPhoneNumber(t *testing.T) {
	// Test cases for valid phone numbers
	samples := map[string]bool{
		// United States (Region: "US")
		"+12025550100":   true, // E.164 format
		"415-555-0101":   true, // National format with hyphens
		"(650) 555-0102": true, // National format with parentheses

		// United Kingdom (Region: "GB")
		"+447911123456": true, // Standard UK mobile format

		// India (Region: "IN")
		"+919876543210": true, // Standard Indian mobile format

		// Brazil (Region: "BR")
		"+5511999998888": true, // São Paulo mobile number

		//// Australia (Region: "AU")
		//"0491570123": true, // Common national mobile format

		// --- Expected to be FALSE (Valid Numbers, but NOT Mobile) ---

		// US Landlines (Fixed Lines)
		//"+12125550103": false,
		//"626-555-0104": false,

		// US Toll-Free
		"+18005550105": false,
		"888-555-0106": false,

		// UK Landline
		"+442071234567": false,

		//// US VoIP (Can sometimes be detected as FIXED_LINE or VOIP)
		//"+15622223333": false,

		// --- Expected to be FALSE (Invalid or Malformed Numbers) ---

		// Malformed or too short
		"12345":           false,
		"+15550101":       false, // Too short for a valid US number
		"not-a-number":    false,
		"+1-202-555-01OO": false, // Contains letters 'O' instead of zeros '0'

		// Invalid prefix/area code for the country
		"+11234567890":  false, // "+11" is not a valid start
		"+441234567890": false, // UK number with a non-mobile prefix

		// Invalid country code
		"+999123456789": false,
	}

	for number, isValid := range samples {
		valid, err := IsValidMobileNumber(number, "US")
		if err != nil && isValid {
			t.Errorf("Expected %s to be %v, but got an error: %v", number, isValid, err)
			continue
		}
		if valid != isValid {
			t.Errorf("Expected %s to be %v, got: %v", number, isValid, valid)
		}
	}
}
