package game

import (
	"os"
	"strings"
	"testing"
)

func Test_NewGame(t *testing.T) {
	game := NewGame()
	if game == nil {
		t.Error("Expected new game instance, got nil")
	}
}

func Test_AskUserForInput(t *testing.T) {
	expectedInput := "42\n"
	game := NewGame()

	// Redirect stdin
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()
	r, w, _ := os.Pipe()
	w.WriteString(expectedInput)
	w.Close()
	os.Stdin = r

	input := game.AskUserForInput()
	if strings.TrimSpace(input) != strings.TrimSpace(expectedInput) {
		t.Errorf("Expected user input %q, got %q", expectedInput, input)
	}
}

func Test_CheckInput(t *testing.T) {
	game := NewGame()
	game.numberToGuess = 42

	// Test with valid input
	if !game.checkInput("42\n") {
		t.Error("Expected true, got false")
	}

	// Test with invalid input
	if game.checkInput("invalid\n") {
		t.Error("Expected false, got true")
	}

	// Test with wrong guess
	if game.checkInput("41\n") {
		t.Error("Expected false, got true")
	}
}

func Test_GuessTemperatures(t *testing.T) {
	var tests = []struct {
		guess         int
		numberToGuess int
		expected      string
		status        bool
	}{
		{5, 5, "Congratulations! You guessed the number in 1 tries.\n", true},
		{5, 10, "You're cold, try again!", false},
		{7, 10, "You're warm, try again!", false},
		{9, 10, "You're hot, try again!", false},
	}

	for _, test := range tests {
		game := NewGame()
		game.tryCount = 1
		message, status := computeGuess(&test.guess, &test.numberToGuess, &game.tryCount)
		if message != test.expected && status != test.status {
			t.Errorf("Expected %q, got %q", test.expected, message)
		}
	}
}

func Test_GenerateRandomNumber(t *testing.T) {
	number := generateRandomNumber()
	if number <= 0 || number > 10 {
		t.Errorf("Expected number between 1 and 10, got %d", number)
	}
}
