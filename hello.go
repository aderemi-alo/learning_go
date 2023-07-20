package main

import (
	"fmt"
	"learning_go/exercises"
)

func main() {
	true := true
	exercises.Negate(&true)
	fmt.Println(true)
	false := false
	exercises.Negate(&false)
	fmt.Println(false)
}
