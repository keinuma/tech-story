package library

import "os"

func IsLocal() bool {
	return os.Getenv("APP_ENV") == "local"
}

func IsProduction() bool {
	return os.Getenv("APP_ENV") == "production"
}
