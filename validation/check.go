package validation

import (
	"regexp"
	"strings"

	"github.com/nyaruka/phonenumbers"
	"github.com/pkg/errors"
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

// IsValidMobileNumber checks if a phone number string is a valid mobile number.
// It requires the number string and a defaultRegion (e.g., "US", "IN", "GB").
func IsValidMobileNumber(numberStr, defaultRegion string) (bool, error) {
	// 1. Parse the number with the default region.
	// The region is crucial for interpreting numbers without a country code.
	phoneNumber, err := phonenumbers.Parse(numberStr, defaultRegion)
	if err != nil {
		return false, errors.Errorf("failed to parse number: %v", err)
	}

	// 2. Check if the number is a possible and valid number for the determined region.
	if !phonenumbers.IsValidNumber(phoneNumber) {
		return false, nil // Not a valid number, so definitely not a valid mobile number.
	}

	// 3. Get the type of the number.
	numberType := phonenumbers.GetNumberType(phoneNumber)

	// 4. Check if the type is MOBILE or FIXED_LINE_OR_MOBILE.
	// Some regions have numbers that can be both, so checking for both is robust.
	isMobile := numberType == phonenumbers.MOBILE || numberType == phonenumbers.FIXED_LINE_OR_MOBILE

	return isMobile, nil
}
