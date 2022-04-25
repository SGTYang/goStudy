package main

import (
	"fmt"

	goTest "github.com/SGTYang/goDicom/goStudy/module/greeting"
)

func main() {
	message := goTest.PrintHello("Test")
	fmt.Println(message)
}
