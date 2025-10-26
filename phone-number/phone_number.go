package phonenumber

import (
	"fmt"
)

func Number(phone string) (string, error) {
	var digits []byte

	for i := 0; i < len(phone) && len(digits) <= 11; i++ {
		c := phone[i]
		if c >= '0' && c <= '9' {
			digits = append(digits, c)
		}
	}

	if err := validate(digits); err != nil {
		return "", err
	}

	return string(digits[len(digits)-10:]), nil
}

func validate(digits []byte) error {
	switch len(digits) {
	case 11:
		if digits[0] != '1' {
			return fmt.Errorf("expected country code '1', got %q", digits[0])
		}
		digits = digits[1:] // remove country code for next checks
		fallthrough
	case 10:
		if digits[0] < '2' || digits[0] > '9' {
			return fmt.Errorf("area code must start with '2'-'9', got %q", digits[0])
		}
		if digits[3] < '2' || digits[3] > '9' {
			return fmt.Errorf("exchange code must start with '2'-'9', got %q", digits[3])
		}
	default:
		return fmt.Errorf("invalid phone number: %s", string(digits))
	}
	return nil
}

// AreaCode returns the 3-digit area code.
func AreaCode(phone string) (string, error) {
	num, err := Number(phone)
	if err != nil {
		return "", err
	}
	return num[:3], nil
}

// Format returns the phone number in (XXX) XXX-XXXX format.
func Format(phone string) (string, error) {
	num, err := Number(phone)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", num[:3], num[3:6], num[6:]), nil
}
