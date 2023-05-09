package pg

import (
	"database/sql"
	"fmt"
	"go-training/task1/log"
	"os"
	"sync"

	_ "github.com/lib/pq"
)

var (
	pg           *sql.DB
	DBClientOnce sync.Once
)

type tables struct {
	myTables string `postgress:"nameTable"`
}
type insertData struct {
	myString string `postgress:"myString" json:"myString"`
	myInt    int    `postgress:"myInt" json:"myInt"`
}

func InsertIntoTable(db *sql.DB, data insertData) {
	log.Info("insert Operation")
	_, err := db.Exec(`INSERT INTO tableName (myString, myInt) VALUES ($1, $2)`)
	checkErrors(err, "error inserting data into the DB")
}

func SelectTables(db *sql.DB) map[string]interface{} {
	log.Info("running command")

	rows, err := db.Query("SELECT table_name FROM information_schema.tables WHERE table_schema='public'")
	checkErrors(err, "invalid command to pg")
	tables := make(map[string]interface{})
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			panic(err)
		}
		tables[tableName] = nil
	}
	return tables

}

func connect(command string) *sql.DB {

	var (
		dbHost   = os.Getenv("dbHost")
		port     = os.Getenv("dbHost")
		user     = os.Getenv("email")
		password = os.Getenv("dbPort")
		dbName   = os.Getenv("dbName")
	)
	log.Info("connecting to the PGDB")
	pslConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmod=disable", dbHost, port, user, password, dbName)
	db, err := sql.Open("postgres", pslConn)
	checkErrors(err, "connection error to pg")
	connectionError := db.Ping()
	checkErrors(connectionError, "connection error to pg")
	log.Info("connection success")
	defer db.Close()
	return db
}

func checkErrors(err error, message string) {
	if err != nil {
		log.Err(message)
		os.Exit(1)
	}
}
