package database

import "fmt"

var connection string

func init() {
	connection = "MySQL"
	fmt.Println("init")
}

func GetDatabase() string {
	return connection
}
