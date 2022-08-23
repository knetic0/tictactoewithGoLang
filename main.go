package main

import (
	"modules/modsrc"
)

func main() {

	table := [3][3]string{
		{"-", "-", "-"},
		{"-", "-", "-"},
		{"-", "-", "-"},
	}

	modsrc.Gui(table, 0)
}
