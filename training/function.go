package training

import (
	"fmt"
	"os"
)

type operation func(int) int

func square(a int) int {
	return a * a
}
func double(a int) int {
	return a * 2
}

func compose(fn1, fn2 operation) operation {
	return func(x int) int {
		return fn1(fn2(x))
	}
}

func IntSeq() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

// variadic
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

type calculate func(int, int) int

func add(a int, b int) int {
	return a + b
}
func sub(a int, b int) int {
	return a - b
}

func divide(a, b float64) (result float64, err error) {
	if b == 0 {
		err = fmt.Errorf("divide by zero")
		return
	}
	result = a / b
	return
}

func read() {
	f, err := os.Open("test.txt")
	if err != nil {
		fmt.Printf("error opening test")
	}
	defer f.Close()
}
