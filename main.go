package main

import (
	"fmt"

	_ "github.com/osama-yousef1993/go-training/training"
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
	// fmt.Println(sum(1,2,3,4))

	// func(name string) {
	// 	fmt.Printf("%s Functions", name)
	// 	fmt.Println(sum(1,2,3,4))
	// }("anonymous")

	// var c calculate
	// c = add
	// fmt.Println(c(1,2))
	// c = sub
	// fmt.Println(c(1,2))

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

	f := add
	fmt.Printf("%d Functions \n", f(3, 3))
	fmt.Printf("%d Functions \n", apply(f, 3, 3))
	errStr := fmt.Sprintf("%s-%d", "test", 10)
	panic(errStr)

}

func add(a int, b int) int {
	return a + b
}

type operation func(int, int) int

func apply(op operation, a int, b int) int {
	return op(a, b)
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

