package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var choices [3]string = [3]string{"ROCK", "PAPER", "SCISSORS"}

func getComputerChoice() string {

	nowUnix := time.Now().UnixNano()
	seed := rand.NewSource(nowUnix)
	newRand := rand.New(seed)
	computerChoiceIndex := newRand.Intn(3)
	computerChoice := choices[computerChoiceIndex]
	return computerChoice
}

func getUserChoice() string {

	var userChoice string
	fmt.Print("Choose a move from [R]ock, [P]aper, [S]cissors: ")
	fmt.Scanf("%s", &userChoice)

	// fmt.Printf("userChoice: %v\n", userChoice)
	userChoice = strings.ToUpper(userChoice)
	// fmt.Printf("userChoice: %v\n", userChoice)
	switch userChoice {
	case "R":
		userChoice = choices[0]
	case "P":
		userChoice = choices[1]
	case "S":
		userChoice = choices[2]
	default:
		userChoice = ""
		fmt.Println("Invalid move.")
	}
	return userChoice
}

func main() {
	userChoice := getUserChoice()

	if userChoice == "" {
		return
	}

	computerChoice := getComputerChoice()

	fmt.Printf("you: %v vs computer: %v\n", userChoice, computerChoice)
	if userChoice == computerChoice {
		fmt.Println("Draw.")
	}

	// TODO
	// Add logic to response if the player wins or lose or draw

	// HINT
	// when win ?
	// 1) player rock && computer scissors
	// 2)        paper && computer rock
	// 3)        scissors && computer paper
	// if (userChoice == choices[0] && computerChoice == choices[2]) || case2 || case3

}
