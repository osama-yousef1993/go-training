package store

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	pg           *sql.DB
	DBClientOnce sync.Once
)

type Tables struct {
	Table string `json:"table" postgres:"table"`
}

func PGConnect() *sql.DB {
	if pg == nil {
		DBClientOnce.Do(func() {
			err := godotenv.Load(".env")
			if err != nil {
				fmt.Printf("Error :: %s", err.Error())
			}
			port, _ := strconv.ParseInt(os.Getenv("DB_PORT"), 10, 64)
			connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", os.Getenv("DB_HOST"), port, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_SSLMODE"))
			pg, err = sql.Open("postgres", connectionString)
			if err != nil {
				fmt.Printf("%s", err)
				return
			}
			connectionError := pg.Ping()

			if connectionError != nil {
				fmt.Printf("%s", connectionError)
				return
			}
		})
	}
	return pg

}

func PGClose() {
	pg.Close()
}

func ReadTables() (*[]Tables, error) {
	pg := PGConnect()

	query := `SELECT table_name as table FROM information_schema.tables WHERE table_schema='public'`

	res, err := pg.Query(query)
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}
	defer res.Close()

	var tables []Tables

	for res.Next() {
		var tableName Tables
		err := res.Scan(&tableName.Table)
		if err != nil {
			log.Printf("%s", err)
			return nil, err
		}
		tables = append(tables, tableName)

	}
	return &tables, err

}
