package config

import (
	"os"
	"strings"
)

// Get from variables inviroment
var connectDb = "DATABASE_URL"
var PathApi = "/api/v1"
var PathU = "users"

func GetDBurl() string {
	return os.Getenv(connectDb) // set DB connect
}

func ConcatenateStrings(strs ...string) string {
	return strings.Join(strs, "/")
}

// func GetPath(...string) string {
// 	a := []string{}
// 	b := ""
// 	for i := 0; i < len(a); i++ {
// 		b += a[i]
// 	}
// 	return b
// }
