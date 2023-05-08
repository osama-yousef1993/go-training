package store

import (
	"database/sql"
	"fmt"
	"os"
	"sync"
	_ "github.com/lib/pq"
)

var (
	pg           *sql.DB
	DBClientOnce sync.Once
)

func PGConnect() *sql.DB {
	if pg == nil {
		DBClientOnce.Do(func() {
			connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_SSLMODE"))
			pg, err := sql.Open("postgres", connectionString)
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

func Te() {
	pg := PGConnect()

	defer pg.Close()
}