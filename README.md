# tictactoe
[![Build Status](https://travis-ci.org/crossi36/tictactoe.svg)](https://travis-ci.org/crossi36/tictactoe)

## What it is
tictactoe implements the basic Tic-Tac-Toe game. tictactoe does not implement the IO. Communicating with the user and
handling his inputs is your responsibility. This enables you to build any user interface around it. Command-line?
Go ahead. GUI? Cool!

## Usage
The first step is to obtain a pointer to a `Game` object. This should be done with the `NewGame()` function.
```
func main() {
	game := tictactoe.NewGame()
```

Next, you will go into a loop, which you will exit as soon as `game.Over()` returns `true`. In this loop you handle the
inputs of your user, which will give you the x and y coordinates of the field the user wants to mark. You can get the
player ID (`0` or `1`) of the player who has to make the next move with the `CurrentPlayer()` method of your `Game`
object.
```
var x, y int
for !game.Over() {
	fmt.Println("Player", game.CurrentPlayer(), "has to play")
	// handle input of the user which will result in x, y being filled
```

You are now able to tell the game of the move the user wants to make. You do that by invoking the `Play()` method of your
`Game` object. This will return an error if the coordinates are out-of-bounds, the game has already ended (someone won or
all fields are marked), or the field is already marked by one of the players.
```
	err := game.Play(x,y)
	if err != nil {
		fmt.Println(err)
	}
}
```

The game assumes the following coordinates for the fields of your board:

![Imgur](http://i.imgur.com/gszGTMo.png)

As soon as you exit the loop it is time to tell your user who has won. You can retrieve the winner with the `Winner()`
method of your `Game` object. `Winner()` either returns the player ID (`Player1` or `Player2`) or `Nobody` if the game resulted in a draw.
```
switch game.Winner() {
case tictactoe.Nobody:
	fmt.Println("The game resulted in a draw!")
default:
	fmt.Println("Player", game.Winner(), "won!")
}
```

A basic implementation of this usage can be found in the example directory.

## License

Distributed under the MIT license. See the [LICENSE](https://github.com/crossi36/tictactoe/blob/master/LICENSE.md) file for details.
