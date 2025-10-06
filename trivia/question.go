package trivia

import (
	"fmt"
	"math/rand"
	"time"
)

// Question represents a single trivia question
type Question struct {
	Text            string   // The actual question text
	CorrectAnswer   string   // The correct answer
	IncorrectAnswers []string // List of wrong answers
	ShuffledChoices []string // All answers combined + shuffled
	CorrectIndex    int      // Position of the correct answer after shuffle
}

// Prepare mixes up answers and assigns labels A, B, C, ...
func (q *Question) Prepare() {
	// Combine correct + incorrect answers
	allAnswers := append([]string{q.CorrectAnswer}, q.IncorrectAnswers...)

	// Shuffle answers
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(allAnswers), func(i, j int) {
		allAnswers[i], allAnswers[j] = allAnswers[j], allAnswers[i]
	})

	// Store shuffled list
	q.ShuffledChoices = allAnswers

	// Find index of the correct answer
	for i, choice := range allAnswers {
		if choice == q.CorrectAnswer {
			q.CorrectIndex = i
			break
		}
	}
}

// Display prints the question and its choices to the console
func (q *Question) Display() {
	fmt.Println("\nQuestion:")
	fmt.Println(q.Text)
	fmt.Println("Choices:")

	labels := []string{"A", "B", "C", "D"}
	for i, choice := range q.ShuffledChoices {
		fmt.Printf("%s) %s\n", labels[i], choice)
	}
}

// CheckAnswer validates if the given label is correct
func (q *Question) CheckAnswer(label string) bool {
	labels := []string{"A", "B", "C", "D"}

	// Convert label to index
	for i, l := range labels {
		if l == label {
			return i == q.CorrectIndex
		}
	}
	return false
}
