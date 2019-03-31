package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var connection = CreateConnectionFromEnvironmentVariables()

type Connection struct {
	*sql.DB
}

func CreateConnectionFromEnvironmentVariables() *Connection {
	var (
		connectionName = getEnvOrFanic("CLOUDSQL_CONNECTION_NAME")
		user           = getEnvOrFanic("CLOUDSQL_USER")
		password       = os.Getenv("CLOUDSQL_PASSWORD") // NOTE: password may be empty
	)

	conn, err := sql.Open("mysql", fmt.Sprintf("%s:%s@cloudsql(%s)/", user, password, connectionName))
	if err != nil {
		log.Panic("Error")
	}

	return &Connection{conn}
}

func getEnvOrFanic(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Panicf("%s environment variable not set.", k)
	}
	return v
}