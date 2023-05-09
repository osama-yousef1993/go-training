package main

import (
	"bufio"
	"fmt"
	"go-training/task1/auth"
	"go-training/task1/log"
	"os"
	"strings"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t.Format("2006-01-02 15:04:05"))
	getEnv()
	log.Info("generating token...")
	auth.GenerateToken()
}

func getEnv() {
	file, err := os.Open(".env")
	if err != nil {
		log.Err("error reading .env file")

	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "=")
		key := parts[0]
		value := parts[1]
		os.Setenv(key, value)
	}
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

// task two
//"github.com/gorilla/mux"
// -- endpoint generate token
// -- endpoiint1 write table // auth user
// -- endpoiint2 read table
// function insert table
// -- string
//-- int
// json *********
// select table
//
