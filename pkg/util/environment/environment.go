package environment

import "os"

// Application Environment name
const (
	Development = "development"
	Test        = "test"
)

// IsDev returns APP_ENV in development mode
func IsDev() bool {
	println("os.Getenv(APP_ENV): ", os.Getenv("APP_ENV"))
	return os.Getenv("APP_ENV") == Development
}

// IsTest returns APP_ENV in test mode
func IsTest() bool {
	return os.Getenv("APP_ENV") == Test
}
