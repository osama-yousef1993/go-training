package store

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/osama-yousef1993/go-training/log"
)

type Table struct {
	tableName string
}

func PgInfoSchema() ([]Table, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.LogErr("error loading environment file " + err.Error())
	}
	log.LogInfo("Connecting to test database")
	db, err := sql.Open("postgres", os.Getenv("PGURI"))
	if err != nil {
		log.LogErr("Problem connecting to database " + err.Error())
		return nil, err
	}
	defer db.Close()
	log.LogInfo("Successfully connecting to db test")
	log.LogInfo("Running information_schema query")
	rows, err := db.Query("SELECT table_name FROM information_schema.tables WHERE table_schema='public'")
	if err != nil {
		log.LogErr("Query faild " + err.Error())
		return nil, err
	}
	defer rows.Close()
	var t Table
	var tables []Table
	for rows.Next() {
		if err := rows.Scan(&t.tableName); err != nil {
			log.LogErr("error scanning rows: " + err.Error())
			return nil, err
		}
		tables = append(tables, t)
	}
	log.LogInfo("Info query was a success")
	return tables, nil
}
