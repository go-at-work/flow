package uuid

import "github.com/google/uuid"

func GenerateUUID() string {
	return uuid.NewString()
}

func Validate(value string) bool {
	if _, err := uuid.Parse(value); err != nil {
		return false
	}
	return true
}
