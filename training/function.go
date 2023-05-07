package training

import (
	"fmt"
	"os"
)

// variadic
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// type function
type calculate func(int, int) int

func add(a int, b int) int {
	return a + b
}
func sub(a int, b int) int {
	return a - b
}

func execute() {
	var c calculate
	c = add
	fmt.Println(c(1, 2))
	c = sub
	fmt.Println(c(1, 2))
}

// Named return function
func divide(a, b float64) (result float64, err error) {
	if b == 0 {
		err = fmt.Errorf("divide by zero")
		return
	}
	result = a / b
	return
}

// defer function 

func read() {
	f, err := os.Open("test.txt")
	if err != nil {
		fmt.Printf("error opening test")
	}
	defer f.Close()
}

// clouser
func IntSeq() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

// composition function

type operations func(int) int

func square(a int) int {
	return a * a
}
func double(a int) int {
	return a * 2
}

func compose(fn1, fn2 operations) operations {
	return func(x int) int {
		return fn1(fn2(x))
	}
}
// pointer function

func addOp(a int, b int) int {
	return a + b
}

type operation func(int, int) int

func apply(op operation, a int, b int) int {
	return op(a, b)
}

func executeOp() {
	f := add
	fmt.Printf("%d Functions \n", f(3, 3))
	fmt.Printf("%d Functions \n", apply(f, 3, 3))
}








