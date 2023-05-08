package pg

import (
	"database/sql"
	"fmt"
	"go-training/task1/log"
	"os"

	_ "github.com/golang-jwt/jwt"
)

var (
	dbHost   = os.Getenv("dbHost")
	port     = 5400
	user     = "postgres"
	password = "password"
	dbName   = "DB_1"
)

func connect(command string) map[int]string {
	log.Info("connecting to the PGDB")
	pslConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmod=disable", dbHost, port, user, password, dbName)
	db, err := sql.Open("postgres", pslConn)
	checkErrors(err, "connection error to pg")
	log.Info("connection success")
	log.Info("running command")
	rows, err := db.Query(command)
	defer db.Close()
	checkErrors(err, "invalid command to pg")
	tables := make(map[int]string)
	i := 0
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			panic(err)
		}
		tables[i] = tableName
		i++
	}
	return tables
}

func checkErrors(err error, message string) {
	if err != nil {
		log.Err(message)
		os.Exit(1)
	}
}
