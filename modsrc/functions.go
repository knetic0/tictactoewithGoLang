package modsrc

import "fmt"

var control bool = true
var tour string = "X"
var winnerControl bool = false

var timer int = 0

func Gui(array [3][3]string, t int) {
	fmt.Println()
	ControlState(array)
	fmt.Println("  ", 0, 1, 2)
	for k, v := range array {
		fmt.Println(k, v)
	}
	if !winnerControl {
		if t != 9 {
			Update(array)
		} else {
			fmt.Println("Tie!")
			return
		}
	}
}

func ControlState(controlArray [3][3]string) {

	if controlArray[0][0] == "X" && controlArray[0][1] == "X" && controlArray[0][2] == "X" && !winnerControl ||
		controlArray[0][0] == "X" && controlArray[1][1] == "X" && controlArray[2][2] == "X" && !winnerControl ||
		controlArray[0][0] == "X" && controlArray[1][0] == "X" && controlArray[2][0] == "X" && !winnerControl ||
		controlArray[0][1] == "X" && controlArray[1][1] == "X" && controlArray[2][1] == "X" && !winnerControl ||
		controlArray[0][2] == "X" && controlArray[1][2] == "X" && controlArray[2][2] == "X" && !winnerControl ||
		controlArray[0][2] == "X" && controlArray[1][1] == "X" && controlArray[2][0] == "X" && !winnerControl ||
		controlArray[1][0] == "X" && controlArray[1][1] == "X" && controlArray[1][2] == "X" && !winnerControl ||
		controlArray[2][0] == "X" && controlArray[2][1] == "X" && controlArray[2][2] == "X" && !winnerControl {
		fmt.Println("X is a Winner !")
		winnerControl = true
		return
	} else if controlArray[0][0] == "O" && controlArray[0][1] == "O" && controlArray[0][2] == "O" && !winnerControl ||
		controlArray[0][0] == "O" && controlArray[1][1] == "O" && controlArray[2][2] == "O" && !winnerControl ||
		controlArray[0][0] == "O" && controlArray[1][0] == "O" && controlArray[2][0] == "O" && !winnerControl ||
		controlArray[0][1] == "O" && controlArray[1][1] == "O" && controlArray[2][1] == "O" && !winnerControl ||
		controlArray[0][2] == "O" && controlArray[1][2] == "O" && controlArray[2][2] == "O" && !winnerControl ||
		controlArray[0][2] == "O" && controlArray[1][1] == "O" && controlArray[2][0] == "O" && !winnerControl ||
		controlArray[1][0] == "O" && controlArray[1][1] == "O" && controlArray[1][2] == "O" && !winnerControl ||
		controlArray[2][0] == "O" && controlArray[2][1] == "O" && controlArray[2][2] == "O" && !winnerControl {
		fmt.Println("O is a Winner !")
		winnerControl = true
		return
	}
}

func Update(upArray [3][3]string) {
	var rows, columns int

	fmt.Printf("%s Please enter row :", tour)
	fmt.Scan(&rows)
	fmt.Printf("%s Please enter row :", tour)
	fmt.Scan(&columns)

	if control && upArray[rows][columns] == "-" {

		upArray[rows][columns] = tour
		tour = "O"
		control = false

	} else if !control && upArray[rows][columns] == "-" {

		upArray[rows][columns] = tour
		tour = "X"
		control = true

	} else {

		fmt.Println("Please try again.")
		Update(upArray)

	}

	timer += 1

	Gui(upArray, timer)
}
