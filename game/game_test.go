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
