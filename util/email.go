package util

import (
	"regexp"
)

// 判断邮箱格式是否合法
func IsValidEmail(email string) bool {
	// Regular expression for email validation
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the regular expression
	re := regexp.MustCompile(emailRegex)

	// Check if the email matches the regular expression
	return re.MatchString(email)
}
