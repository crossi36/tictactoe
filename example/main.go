package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/crossi36/tictactoe"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	game := tictactoe.NewGame()

	fmt.Println("Tic-Tac-Toe")
	for !game.Over() {
		fmt.Println(game)
		fmt.Println("It's the turn of player", game.CurrentPlayer())
		fmt.Println("Enter the coordinates like so: x,y")

		input, _ := reader.ReadString('\n')
		fmt.Println()
		parts := strings.Split(strings.TrimSpace(input), ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])

		err := game.Play(x, y)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println(game)
	switch game.Winner() {
	case tictactoe.Nobody:
		fmt.Println("The game resulted in a draw!")
	default:
		fmt.Println("Player", game.Winner(), "won!")
	}
}
