package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
)

type Game struct {
	numberToGuess int
	tryCount      int
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Start() {
	g.numberToGuess = generateRandomNumber()
	fmt.Println("I've picked a number between 1 and 10. Try to guess it!")
	g.tryCount = 0

	for {
		input := g.AskUserForInput()
		g.tryCount++
		if g.checkInput(input) {
			break
		}
	}
}

func (g *Game) AskUserForInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a number: ")
	text, _ := reader.ReadString('\n')
	return text
}

func (g *Game) checkInput(input string) bool {
	matched, _ := regexp.MatchString("^[0-9]+$", strings.Trim(input, "\n"))
	if !matched {
		fmt.Println("Please enter a valid number.")
		return false
	}
	var guess int
	fmt.Sscanf(input, "%d", &guess)
	message, status := computeGuess(&guess, &g.numberToGuess, &g.tryCount)
	fmt.Println(message)
	return status
}

func generateRandomNumber() int {
	return rand.Intn(10)
}

func computeGuess(guess, numberToGuess, tryCount *int) (string, bool) {
  switch {
    case *guess == *numberToGuess: return fmt.Sprintf("Congratulations! You guessed the number in %d tries.\n", *tryCount), true
    case (*guess - *numberToGuess) >= 5: return "You're cold, try again!", false
    case (*guess - *numberToGuess) >= 3: return "You're warm, try again!", false
    case (*guess - *numberToGuess) >= 1: return "You're hot, try again!", false
    default: return "Wrong nuess", false  
  }
}
