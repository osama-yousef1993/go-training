package training

import "fmt"



func IfElse() {
	fmt.Println("Start If Else Statement \n")
	var age int

	fmt.Println("Enter Your Age:")
	fmt.Scan(&age)

	if age >= 18 {
		fmt.Println("Your are an Adult")
	} else {
		fmt.Println("Your are a child")
	}
	fmt.Println("Finish If Else Statement \n")
}

func ForLoop() {
	fmt.Println("Start For Loop Statement \n")

	for i := 0; i <=5; i++{
		fmt.Printf("%d \n", i)
	}

	// while loop
	 
	j := 1

	for j <=5 {
		fmt.Printf("%d \n", j)
		j++
	}


	// do while

	x := 1

	for {
		fmt.Printf("%d \n", x)
		x++

		if x >= 5 {
			break
		}
	}

	fmt.Println("Finish For Loop Statement \n")

}

func SwitchStatement() {

	fmt.Println("Start Switch Statement \n")
	day := "Sunday"

	switch day {
		case "Sunday":
			fmt.Println("Today is Sunday")
		case "Monday":
			fmt.Println("Today is Monday")
		case "Tuesday":
			fmt.Println("Today is Tuesday")
		case "Wednesday":
			fmt.Println("Today is Wednesday")
		case "Thursday":
			fmt.Println("Today is Thursday")
		default:
			fmt.Println("Today is not Sunday")
	}
	fmt.Println("Finish Switch Statement \n")

}