package main

import (
	"fmt"

	"github.com/johnmphillips/dice-roller/dice"
)

func main() {
	// Roll 3d6
	result, err := dice.Roll("3d6")
	if err != nil {
		panic(err)
	}

	// Print the total
	fmt.Println(result.Total())
}
