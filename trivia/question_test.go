package trivia

import (
	"strings"
	"testing"
)

// helper function to create a sample question
func sampleQuestion() *Question {
	return &Question{
		Text:            "What is 2 + 2?",
		CorrectAnswer:   "4",
		IncorrectAnswers: []string{"3", "5", "22"},
		Category:        "Math",
		Difficulty:      "Easy",
	}
}

func TestPrepare(t *testing.T) {
	q := sampleQuestion()
	q.Prepare()

	if len(q.ShuffledChoices) != 4 {
		t.Fatalf("Expected 4 choices, got %d", len(q.ShuffledChoices))
	}

	foundCorrect := false
	for i, choice := range q.ShuffledChoices {
		if choice == q.CorrectAnswer {
			if i != q.CorrectIndex {
				t.Errorf("CorrectIndex %d does not match correct answer position %d", q.CorrectIndex, i)
			}
			foundCorrect = true
		}
	}
	if !foundCorrect {
		t.Error("Correct answer not found in ShuffledChoices")
	}
}

func TestCheckAnswer(t *testing.T) {
	q := sampleQuestion()
	q.Prepare()

	labels := []string{"A", "B", "C", "D"}
	correctLabel := labels[q.CorrectIndex]

	if !q.CheckAnswer(correctLabel) {
		t.Errorf("CheckAnswer failed for correct label %s", correctLabel)
	}

	for i, label := range labels {
		if i != q.CorrectIndex {
			if q.CheckAnswer(label) {
				t.Errorf("CheckAnswer returned true for wrong label %s", label)
			}
		}
	}

	// Test lowercase input (should fail in current implementation)
	if q.CheckAnswer(strings.ToLower(correctLabel)) {
		t.Errorf("CheckAnswer should be case-sensitive, lowercase label should fail")
	}

	// Test invalid label
	if q.CheckAnswer("Z") {
		t.Errorf("CheckAnswer returned true for invalid label 'Z'")
	}
}

// TestDisplay just ensures it runs without panicking
func TestDisplay(t *testing.T) {
	q := sampleQuestion()
	q.Prepare()

	// We don't capture stdout here, just verify it doesn't panic
	q.Display()
}
