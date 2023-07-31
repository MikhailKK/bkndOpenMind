package config

import "os"

// Get from variables inviroment
var connect = "DATABASE_URL"

func GetDBurl() string {
	return os.Getenv(connect) // set DB connect
}
