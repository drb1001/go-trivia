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
	fmt.Println("Welcome to Go Trivia!")

	// Load highscore from file
	highScore := trivia.LoadHighScore("data/highscore.json")
	fmt.Printf("Current High Score: %d\n\n", highScore)

	// Fetch 5 questions
	questions, err := trivia.FetchQuestions(5)
	if err != nil {
		log.Fatalf("Error fetching questions: %v", err)
	}

	score := 0
	reader := bufio.NewReader(os.Stdin)

	// Ask questions one by one
	for _, q := range questions {
		q.Display()

		// Get user input
		fmt.Print("\nYour answer (A/B/C/D): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToUpper(input))

		// Check if correct
		if q.CheckAnswer(input) {
			fmt.Println("✅ Correct!")
			score++
			fmt.Printf("Current Streak: %d\n\n", score)
		} else {
			fmt.Printf("❌ Wrong! The correct answer was: %s\n", q.CorrectAnswer)
			break
		}
	}

	fmt.Printf("\nGame over! Your final score: %d\n", score)

	// Update highscore if beaten
	if score > highScore {
		fmt.Println("🎉 New High Score!")
		trivia.SaveHighScore("data/highscore.json", score)
	} else {
		fmt.Printf("High Score remains: %d\n", highScore)
	}
}
