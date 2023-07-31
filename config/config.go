package config

import "os"

func GetDBurl() string {
	return os.Getenv("DataBase url: ") // set DB connect
}
