package trivia

import (
	"fmt"
	"math/rand"
	"time"
)

// ANSI color codes (work on most terminals)
const (
	ColorReset  = "\033[0m"
	ColorBold   = "\033[1m"
	ColorGreen  = "\033[32m"
	ColorRed    = "\033[31m"
	ColorYellow = "\033[33m"
	ColorCyan   = "\033[36m"
)

// Question represents a single trivia question
type Question struct {
	Text            string
	CorrectAnswer   string
	IncorrectAnswers []string
	ShuffledChoices []string
	CorrectIndex    int
	Category        string
	Difficulty      string
}

// Prepare mixes up answers and assigns labels
func (q *Question) Prepare() {
	allAnswers := append([]string{q.CorrectAnswer}, q.IncorrectAnswers...)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(allAnswers), func(i, j int) {
		allAnswers[i], allAnswers[j] = allAnswers[j], allAnswers[i]
	})

	q.ShuffledChoices = allAnswers
	for i, choice := range allAnswers {
		if choice == q.CorrectAnswer {
			q.CorrectIndex = i
			break
		}
	}
}

// Display prints the question neatly
func (q *Question) Display() {
	labels := []string{"A", "B", "C", "D"}

	fmt.Printf("\n%s───────────────────────────────%s\n", ColorCyan, ColorReset)
	fmt.Printf("%sCategory:%s %s | %sDifficulty:%s %s\n",
		ColorYellow, ColorReset, q.Category, ColorYellow, ColorReset, q.Difficulty)
	fmt.Printf("%sQ:%s %s\n", ColorBold, ColorReset, q.Text)
	fmt.Println("")

	for i, choice := range q.ShuffledChoices {
		fmt.Printf("  %s%s)%s %s\n", ColorCyan, labels[i], ColorReset, choice)
	}
	fmt.Printf("%s───────────────────────────────%s\n", ColorCyan, ColorReset)
}

// CheckAnswer validates if the given label is correct
func (q *Question) CheckAnswer(label string) bool {
	labels := []string{"A", "B", "C", "D"}
	for i, l := range labels {
		if l == label {
			return i == q.CorrectIndex
		}
	}
	return false
}
