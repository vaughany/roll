package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Get the command-line arguments.
	rolls := os.Args[1:]

	if len(rolls) == 0 {
		fmt.Println("Run this program with one or more arguments in the format 'XdY', where X is the number of dice to roll, and Y is the number of faces, e.g. 2d20 1d4")
		os.Exit(0)
	}

	// Seed the random number generator.
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Variable for the total of all rolls.
	var totalScore int

	// For each argument, run this loop.
	for rollIndex, roll := range rolls {
		// Space out results after the first one.
		if rollIndex != 0 {
			fmt.Println()
		}

		// Parse the arguments.
		command := strings.Split(strings.ToLower(roll), "d")

		// Convert the pre-d and post-d numbers into integers, and skip if errors.
		diceCount, err := strconv.Atoi(command[0])
		if err != nil {
			fmt.Printf("err: %v\n", err)
			continue
		}
		diceFaces, _ := strconv.Atoi(command[1])
		fmt.Printf("Rolling %dd%d:\n", diceCount, diceFaces)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			continue
		}

		// Sanity-check the number of dice, and the number of faces. There's currently no upper limit for either.
		if diceCount < 1 {
			fmt.Println("err: can't have less than 1 dice")
			continue
		}
		if diceFaces < 2 {
			fmt.Println("err: can't have dice with less than 2 faces")
			continue
		}

		// Variable for the score for this set of rolls.
		var score int

		for j := 1; j <= diceCount; j++ {
			// Variable for this roll.
			var rollScore int

			// Roll that dice.
			rollScore += r.Intn(diceFaces) + 1

			// If we're rolling more than one dice, present the information differently.
			if diceCount > 1 {
				// Build the output as a string.
				var out string

				// Each dice gets it's own line.
				out += fmt.Sprintf("> Dice %d rolled a %d.", j, rollScore)

				// If we're rolling a d20, display extra text for a 1 or a 20.
				if diceFaces == 20 {
					switch rollScore {
					case 20:
						out += " - CRIT!"
					case 1:
						out += " - NAT 1!"
					}
				}

				// Print to screen.
				fmt.Println(out)
			}

			// Add the rolled score to the score.
			score += rollScore
		}

		// Display output for this set of rolls.
		fmt.Printf("The result of your %dd%d roll is %d.\n", diceCount, diceFaces, score)

		// Add that set of scores to the total score.
		totalScore += score
	}

	// Final output.
	fmt.Printf("\nThe total result of your rolls is %d.\n", totalScore)
}
