package main

import (
	"fmt"
	"github.com/joho/godotenv"
	auth "github.com/osama-yousef1993/go-training/auth"
	log "github.com/osama-yousef1993/go-training/log"
	"github.com/osama-yousef1993/go-training/store"
	_ "github.com/osama-yousef1993/go-training/training"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.LogErr("error loading environment file " + err.Error())
	}

	tables, err := store.PgInfoSchema()
	if err != nil {
		log.LogErr(err.Error())
	}
	fmt.Println(tables)

	values, err := store.BqInfoSchema()
	if err != nil {
		log.LogErr(err.Error())
	}
	fmt.Println(values)

	token, err := auth.GenerateToken("name@something.com")
	if err != nil {
		errStr := fmt.Sprintf("error generating token - %s", err)
		log.LogErr(errStr)
	}

	err = auth.CheckToken(token)
	if err != nil {
		log.LogErr(err.Error())
	}
	log.LogInfo("Token is valid")
}

// task
// packages
// -- go get github.com/lib/pq
// -- go get github.com/golang-jwt/jwt
// -- cloud.google.com/go/bigquery
// auth
// -- jwt
// -- generate token // string
// -- check token // @hotmail// @gmail // string
// .env file
// postgres
// -- connections
// -- struct nameTable
// -- map (array of table) to table name from postgres
// -- SELECT table_name FROM information_schema.tables WHERE table_schema='public'
// bigquery
// -- connections
// -- struct nameTable
// -- map (array of table) to table name from postgres
// log
//--
