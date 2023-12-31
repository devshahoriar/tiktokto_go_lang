package main

import (
	"fmt"
	"os"
	"shuvo/inc"
)

func main() {
	var input string
	var playFirstUser bool
	game := make([]string, 9)
	for i := 0; i < len(game); i++ {
		game[i] = "-"
	}
	// PCIN, USERIN := "o", "×"
	inc.Clear()
	fmt.Printf("__________________________\n")
	fmt.Printf("** welcome to ticktokto ** \n")
	fmt.Printf("--------------------------\n")
	fmt.Printf("playing instraction :\n |1-2-3|\n |4-5-6|\n |7-8-9|\nYou just enter this number to play!\n\n")
	fmt.Printf("Do you want to play first? (y/n) \n")
	fmt.Scanln(&input)
	if inc.InputIsOk(&input, "y", "n") {
		fmt.Printf("wrong input! Try again😑")
		os.Exit(0)
	}
	if input == "y" {
		fmt.Println("You want to play first!")
		playFirstUser = true
	} else {
		fmt.Println("You want pc play first!")
		playFirstUser = false
	}
	inc.Clear()
	fmt.Println("Play tiktokto -")
	inc.PrintGame(game)
	inc.StartGame(&game, playFirstUser)
}
