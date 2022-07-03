package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//num := 34
	var guess int
	fmt.Print("The hidden number is in 0-99\n")
	nowUnix := time.Now().UnixNano()
	seed := rand.NewSource(nowUnix)
	newRand := rand.New(seed)

	//fmt.Printf("nowUnix: %v\n", nowUnix)
	// randNum := rand.Int()
	// randNum := newRand.Int()
	randNum := newRand.Intn(100)
	fmt.Println(randNum)

	for guess != randNum {
		fmt.Print("Type your guess: ")
		fmt.Scan(&guess)
		fmt.Printf("Your guess is %v\n", guess)
		if guess == randNum {
			fmt.Println("You are correct!")
		} else {
			fmt.Println("You are not correct yet.")
			if guess < 0 {
				fmt.Println("Guess must not be less than 0.")
			} else if guess > 99 {
				fmt.Println("Guess must not be greater than 99.")
			} else if guess < randNum {
				fmt.Println("guess is less than num.")
			} else {
				fmt.Println("guess is greater than num.")
			}
			fmt.Println("Don't give up! The correct number is still 0-99")
		}
	}
}
