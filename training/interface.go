package training

import "fmt"

// all function now for interface types example
type Shape interface {
	Area() float64
	Perimeter() float64
}
type Rectangle struct {
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * r.Width * r.Height
}

func InterfaceTypes() {
	var s Shape
	var r Rectangle

	r.Height = 20.5
	r.Width = 32.5
	fmt.Println("Start Interface Types")
	fmt.Printf("%f \n", r.Area())
	fmt.Printf("%f \n", r.Perimeter())

	s = r

	fmt.Printf("%f \n", s.Area())
	fmt.Printf("%f \n", s.Perimeter())
	fmt.Println("Finish Interface Types")
}
