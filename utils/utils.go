package utils

// IsDev returns APP_ENV in development mode
func IsDev() bool {
	return true
	//println("os.Getenv(APP_ENV): ", os.Getenv("APP_ENV"))
	//return os.Getenv("APP_ENV") == "development"
}
