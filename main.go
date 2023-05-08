package main

import (
	"fmt"

	"github.com/osama-yousef1993/go-training/auth"
)

func main() {
	// fmt.Println("Data Types Function")
	// training.BasicTypes()
	// training.AggregateTypes()
	// training.InterfaceTypes()

	// fmt.Println("Control Flow Functions")
	// training.IfElse()
	// training.ForLoop()
	// training.SwitchStatement()

	email, err := auth.GenerateToken("test@hotmail.com")

	if err != nil {

		fmt.Printf("%s Error", err.Error())
	}
	fmt.Printf("%s ", email)
	// email := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJlbWFpbCI6InRlc3RAaG90bWFpbC5jb20iLCJleHAiOjE2ODM0OTY1MTl9.Vk3j2uu35RJ_9Py7F6mi4fTT79VH-NuNWGDscO8G718"

	token, err := auth.ValidateToken(email)
	if err != nil {

		fmt.Printf("%s Error", err.Error())
	}
	fmt.Printf("%s ", token)

	// func(name string) {
	// 	fmt.Println(sum(1,2,3,4))
	// }("anonymous")

	// res, err := divide(10, 0)
	// if err != nil {
	// 	fmt.Printf("%s Functions", err.Error())
	// }
	// fmt.Printf("%f Functions", res)

	// x := 0

	// incr := IntSeq()

	// incr()
	// incr()
	// incr()
	// fmt.Printf("%d Functions", x)
	// fn := compose(square, double)
	// fmt.Printf("%d Functions", fn(3))

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