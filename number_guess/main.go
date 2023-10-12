package main

import (
	"fmt"
	"math/rand"
)

var score uint8 = 0
const upper_lmt int = 10

func getSequence(seq []int) []int {
	for i := 0; i < 3; i++ {
		num := rand.Intn(upper_lmt)
		seq = append(seq, num)
	}
	return seq 
}

func showNums(seq []int) int {
	hidden_idx := rand.Intn(len(seq))
	for i := 0; i < len(seq); i++ {
		if (i == hidden_idx) {
			fmt.Printf(" # ")
		} else {
			fmt.Printf(" %v ", seq[i])
		}
	}
	fmt.Printf("\n")
	// returns the hidden value index
	return hidden_idx
}

func getAnswer(hidden_idx int, seq []int) {
	var guess int;

	fmt.Printf("Enter your guess: ")
	fmt.Scan(&guess)

	for guess != seq[hidden_idx] {
		var try_again string
		fmt.Printf("WRONG GUESS!!! Try Again[y/n]: ")
		fmt.Scan(&try_again)
		if (try_again == "y") || (try_again == "Y") {
			fmt.Printf("Enter your guess: ")
			fmt.Scan(&guess)
		} else {
			// user opts out
			fmt.Printf("Answer was: %d\n", seq[hidden_idx])
			break
		}
	}
	// user entered correct guess
	score = score + 1
}

func showResults() {
	fmt.Printf("Thank you for playing, your score is: %d", score)
}

func playGame(hidden_idx int, seq []int) {
	getAnswer(hidden_idx, seq)

	var prompt string
	fmt.Printf("play again[y/n]: ")
	fmt.Scan(&prompt)

	if (prompt == "y") || (prompt == "Y") {
		play()
	} else {
		showResults()
	}
}

func play () {
	seq := []int{}
	seq = getSequence(seq)
	hidden_idx := showNums(seq)
	playGame(hidden_idx, seq)
}

func main() {
	play()
}
