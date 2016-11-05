package main

import (
	"fmt"
	"os"

	kata "github.com/egnwd/kata/rover"
)

func main() {
	r := kata.Rover{X: 0, Y: 0, D: kata.North}

	r.Command(os.Args[1])
	fmt.Println(&r)
}
