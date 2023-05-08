package log

import (
	"fmt"
	"os"
)

func Err(err string) {
	fmt.Println(err)
}
func Info(info string) {
	fmt.Println(info)
}
func Test() {
	fmt.Println(os.Getenv("t1"))
	fmt.Println("fail")
}
