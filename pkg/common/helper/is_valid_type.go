package helper

import (
	"regexp"
	"strings"
)

func IsValidImageType(contentType string) bool {
	// List of allowed image MIME types (you can add more if needed)
	allowedTypes := []string{
		"jpeg",
		"png",
		"gif",
		// Add more allowed types here...
	}

	contentType = strings.ToLower(contentType)
	for _, allowedType := range allowedTypes {
		// Use regular expression to check if the content type matches the allowed types
		if ok, _ := regexp.MatchString("^"+allowedType, contentType); ok {
			return true
		}
	}
	return false
}
