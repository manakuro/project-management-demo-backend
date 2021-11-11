package utils

import "os"

// IsDev returns APP_ENV in development mode
func IsDev() bool {
	println("os.Getenv(APP_ENV): ", os.Getenv("APP_ENV"))
	return os.Getenv("APP_ENV") == "development"
}
