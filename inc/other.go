package inc

import (
	"fmt"
	"math/rand"
	"strconv"
)

const PC_IN, USERIN string = "o", "×"

func Clear() {
	fmt.Print("\033[H\033[2J")
}

func InputIsOk(input *string, accept ...string) (exsist bool) {
	exsist = true
	for _, a := range accept {
		if *input == a {
			exsist = false
			return
		}
	}
	return

}

// arr = [] {-,-,-,-,-,-,-,,-} len =9

func PrintGame(game []string) {
	fmt.Print("-----------\n|")
	for i := 0; i < len(game); i++ {
		if game[i] == "-" {
			fmt.Printf(" █ ")
		} else {
			fmt.Printf(" " + game[i] + " ")
		}
		if (i+1)%3 == 0 {
			fmt.Print("|\n")
			if i+1 != 9 {
				fmt.Print("|")
			}
		}
	}
	fmt.Print("-----------")
}

func StartGame(game *[]string, playFirstUser bool) {
	fmt.Printf("\n")
	for {
		var x int
		fmt.Println("\nYour Turn: enter(1-9)")
		fmt.Scan(&x)
		(*game)[x-1] = USERIN
		(*game)[getRandomForPc(*game)-1] = PC_IN
		Clear()
		PrintGame(*game)
		bb := safeTurn(*game, x)
		fmt.Println(strconv.FormatBool(bb))
	}
}

func getRandomForPc(game []string) int {
	x := rand.Intn(9-1) + 1
	return x
}

func safeTurn(game []string, input int) (x bool) {
	x = false
	if game[input] == PC_IN || game[input] == USERIN {
		x = true
	}
	return
}
