package config

import (
	"os"
	"strings"
)

// Get from variables inviroment
var connectDb = "DATABASE_URL"
var PathApi = "/api/v1"
var PathU = "/users"
var PathID = "/:id"
var PathP = "/products"
var PathO = "/orders"

var Username = "USER_NAME"
// var Password = "PASSWORD"

func GetDBurl() string {
	return os.Getenv(connectDb) // set DB connect
}

func ConcatenateStrings(strs ...string) string {
	return strings.Join(strs, "")
}

func GetUserName() string {
	return os.Getenv(Username)
}

// это не рабоатет, пароль лежит в auth
// func GetPassword() string {
// 	return os.Getenv(Password)
// }



// func GetPath(...string) string {
// 	a := []string{}
// 	b := ""
// 	for i := 0; i < len(a); i++ {
// 		b += a[i]
// 	}
// 	return b
// }
