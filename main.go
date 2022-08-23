package main

import (
	"fmt"
	"modules/modsrc"
)

func main() {

	table := [3][3]string{
		{"-", "-", "-"},
		{"-", "-", "-"},
		{"-", "-", "-"},
	}

	modsrc.Gui(table, 0)

	for i := 0; i < 9; i++ {
		if i == 9 {
			fmt.Println("Game Over! Tie.")
		}
	}
}
