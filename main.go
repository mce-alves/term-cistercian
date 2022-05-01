package main

import (
	"fmt"
	"strconv"
)

func main() {
	var exit bool = false

	for !exit {
		fmt.Println("\nEnter a number from 0 to 9999 to convert to Cistercian representation: ")
		var line string
		fmt.Scanln(&line)

		number, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("\t Not a number. Safely exiting the program...")
			exit = true
		} else {
			if number < 0 || number > 9999 {
				fmt.Println("\t Number outside of valid range: [0, 9999]")
				continue
			}

			cn := Convert(number)
			cn.Draw()
		}
	}

}
