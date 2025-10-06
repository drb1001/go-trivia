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
		// Load high score and ensure secret key exists
		highName, highScore, err := trivia.LoadHighScore("data/highscore.json")
		if err != nil {
			// Friendly error and instructions on how to create config
			configPath, _ := func() (string, error) { return trivia.GetConfigPathForUser() }()
			msg := fmt.Sprintf("ERROR: %v\n\n"+
				"Please create a config file with a secret key used for obfuscation of highscores.\n"+
				"Create it at: %s\n"+
				"Example content:  { \"secret_key\": \"ExampleSecretKey\" }\n\n", 
				err, configPath)
			
			log.Fatalf(msg)
		}

		if highScore > 0 {
			fmt.Printf("%sCurrent High Score:%s %s%d%s (by %s)\n\n",
				trivia.ColorYellow, trivia.ColorReset,
				trivia.ColorGreen, highScore, trivia.ColorReset, highName)
		} else {
			fmt.Printf("%sCurrent High Score:%s %s%d%s\n\n",
				trivia.ColorYellow, trivia.ColorReset,
				trivia.ColorGreen, 0, trivia.ColorReset)
		}

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

			var input string
			for {
				fmt.Print("\nYour answer (A/B/C/D): ")
				userInput, _ := reader.ReadString('\n')
				input = strings.TrimSpace(strings.ToUpper(userInput))

				if input == "A" || input == "B" || input == "C" || input == "D" {
					break
				}

				fmt.Println(trivia.ColorYellow + "⚠️  Please enter only A, B, C, or D." + trivia.ColorReset)
			}

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

		// Update highscore if beaten
		if score > highScore {
			fmt.Println(trivia.ColorYellow + "🎉 New High Score! 🎉" + trivia.ColorReset)

			// Prompt for name
			fmt.Print("Enter your name to record the high score: ")
			nameInput, _ := reader.ReadString('\n')
			name := strings.TrimSpace(nameInput)
			if name == "" {
				name = "Anonymous"
			}

			// Save obfuscated highscore
			if err := trivia.SaveHighScore("data/highscore.json", name, score); err != nil {
				fmt.Printf("Error saving high score: %v\n", err)
			} else {
				fmt.Printf("Saved high score: %s - %d\n", name, score)
			}
		} else {
			fmt.Printf("High Score remains: %d (by %s)\n", highScore, highName)
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
