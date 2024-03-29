package training

import "fmt"

func BasicTypes() {
	// in this function will use syntax for basic types in golang
	// int
	// float64
	// bool
	// string
	// first declaration
	var x int = 34
	var y float64 = 34
	var z string = "hello world"
	var c bool = true

	fmt.Println("Start Basic Types")
	fmt.Printf("%d \n", x)
	fmt.Printf("%f \n", y)
	fmt.Printf("%s \n", z)
	fmt.Printf("%t \n", c)

	// second declaration
	x1 := 10
	y1 := 1.3
	z1 := "hello world"
	c1 := false
	fmt.Printf("%d \n", x1)
	fmt.Printf("%f \n", y1)
	fmt.Printf("%s \n", z1)
	fmt.Printf("%t \n", c1)
	fmt.Println("Finish Basic Types \n")
}

func AggregateTypes() {
	// array
	// slice
	// struct
	// map
	var a [3]int
	var s []int
	var t struct {
		x int
		y int
	}

	var g = make(map[string]string)

	a[0] = 1
	a[1] = 1
	a[2] = 1

	s = append(s, 1)
	s = append(s, 3)

	t.x = 10
	t.y = 17

	g["test"] = "hello world"
	fmt.Println("Start Aggregate Types")
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", s)
	fmt.Printf("%v \n", t)
	fmt.Printf("%v \n", g["test"])

	fmt.Println("Finish Aggregate Types \n")
}

func ReferenceTypes() {
	// pointer
	var x int = 10
	var opt *int = &x
	fmt.Printf("%v \n", *opt)

}
