package main

import (
	// "learning_go/exercises"
	"fmt"
	"reflect"
)

func main() {
	// exercises.GuessGame()
	hello := "hello"
	fmt.Println(&hello)
	fmt.Println(reflect.TypeOf(&hello))
	fmt.Println(1 % 2)
}
