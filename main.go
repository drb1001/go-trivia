package main

import (
	"bufio"
	"fmt"
	"go-trivia/trivia"
	"log"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Colored welcome message
	fmt.Printf("%s%sWelcome to Go Trivia!%s 🎯\n", trivia.ColorBold, trivia.ColorCyan, trivia.ColorReset)
	fmt.Println("──────────────────────────────────────────────")

	for {
		// Load high score from file
		highScore := trivia.LoadHighScore("data/highscore.json")
		fmt.Printf("%sCurrent High Score:%s %s%d%s\n\n",
			trivia.ColorYellow, trivia.ColorReset,
			trivia.ColorGreen, highScore, trivia.ColorReset)

		// Fetch questions
		fmt.Println(trivia.ColorCyan + "Fetching questions..." + trivia.ColorReset)
		questions, err := trivia.FetchQuestions(5)
		if err != nil {
			log.Fatalf("Error fetching questions: %v", err)
		}

		score := 0

		// Game loop
		for _, q := range questions {
			q.Display()

			// Get user input
			fmt.Print("\nYour answer (A/B/C/D): ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(strings.ToUpper(input))

			// Check if correct
			if q.CheckAnswer(input) {
				fmt.Println(trivia.ColorGreen + "✅ Correct!" + trivia.ColorReset)
				score++
				fmt.Printf("Current Streak: %d\n\n", score)
			} else {
				fmt.Println(trivia.ColorRed + "❌ Wrong!" + trivia.ColorReset)
				fmt.Printf("The correct answer was: %s\n", q.CorrectAnswer)
				break
			}
		}

		fmt.Printf("\n%sGame over!%s Your final score: %s%d%s\n",
			trivia.ColorBold, trivia.ColorReset, trivia.ColorGreen, score, trivia.ColorReset)

		// Update high score if beaten
		if score > highScore {
			fmt.Println(trivia.ColorYellow + "🎉 New High Score! 🎉" + trivia.ColorReset)
			trivia.SaveHighScore("data/highscore.json", score)
		} else {
			fmt.Printf("High Score remains: %d\n", highScore)
		}

		// Ask to play again
		fmt.Print("\nPlay again? (y/n): ")
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(strings.ToLower(answer))

		if answer != "y" && answer != "yes" {
			fmt.Println("\n" + trivia.ColorCyan + "Thanks for playing Go Trivia! 👋" + trivia.ColorReset)
			break
		}

		fmt.Println("\n" + trivia.ColorYellow + "🔁 Starting a new game..." + trivia.ColorReset)
		fmt.Println("──────────────────────────────────────────────")
	}
}
