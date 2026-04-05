package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Welcome to the Dice Game!")

	game := newGame()
	fmt.Println("Enter how many players will play: ")
	var playersCount int
	for {
		fmt.Scanln(&playersCount)
		if playersCount < 3 {
			fmt.Println("Please enter a number greater than 2")
			continue
		}
		break
	}
	for i := 0; i < playersCount; i++ {
		fmt.Printf("Enter the name of player %d: ", i+1)
		var name string
		fmt.Scanln(&name)
		game.join(name)
	}
	game.playGame()
}

func newGame() *Game {
	return &Game{}
}

// Join a player to the game
func (g *Game) join(name string) *Player {
	player := Player{
		name:  name,
		token: 3,
	}
	playerCount := len(g.players)
	if playerCount > 0 {
		lastPlayer := g.players[playerCount-1]
		player.left = lastPlayer
		player.rigth = g.players[0]
		lastPlayer.rigth = &player
		g.players[0].left = &player

	}
	g.players = append(g.players, &player)
	return &player
}

// Check if we have only one player with tokens, this should be done after each turn
func (g *Game) checkWinner() *Player {
	playersWithTokens := 0
	var winner *Player
	for _, value := range g.players {
		if value.token > 0 {
			playersWithTokens++
			if playersWithTokens > 1 {
				return nil
			}
			winner = value
		}
	}
	return winner
}

// New game instance
type Game struct {
	players []*Player
}

// Add a player to the game
type Player struct {
	name  string
	token int
	rigth *Player
	left  *Player
	dice  *Dice
}

// Roll the dice

// roll uses rand to return diceface
func (Player *Player) rollDice() (result []string) {
	// find how many dice the player should roll, this is the minimum between the number of tokens and 3
	diceToRoll := Player.token
	if diceToRoll > 2 {
		diceToRoll = 3
	}
	for index := 0; index < diceToRoll; index++ {
		d := Dice{}
		diceResult := d.roll()
		result = append(result, diceResult)
		switch diceResult {
		case "Rigth":
			Player.rigth.token++
			Player.token--
		case "Left":
			Player.left.token++
			Player.token--
		case "Center":
			Player.token--
		}
	}
	return result
}

type Dice struct {
	value int
}

// Roll the dice and return the result as a string, also update the value of the dice
func (d *Dice) roll() string {
	rand.Seed(time.Now().UnixNano())

	r := rand.Intn(6)
	switch r {
	case 0:
		d.value = 0
		return "Rigth"
	case 1:
		d.value = 1
		return "Left"
	case 2:
		d.value = 2
		return "Center"
	default:
		return "Dot"
	}
}

// Display the current game status
func (g *Game) displayGameStatus() {
	fmt.Println("\n========== Current Status ==========")
	for i, player := range g.players {
		fmt.Printf("Player %d: %s - Tokens: %d\n", i+1, player.name, player.token)
	}
	fmt.Println("===================================\n")
}

// Get the next player in the circular linked list
func (g *Game) getNextPlayer(current *Player) *Player {
	if current.rigth == nil {
		return g.players[0]
	}
	return current.rigth
}

// Execute one player's turn
func (g *Game) takeTurn(player *Player) {
	if player.token == 0 {
		fmt.Printf("%s has no tokens, skipping turn.\n", player.name)
		return
	}

	fmt.Printf("\n%s's turn!\n", player.name)
	fmt.Println("Press Enter to roll the dice...")
	fmt.Scanln()

	diceResults := player.rollDice()
	fmt.Printf("%s rolled: %v\n", player.name, diceResults)
	fmt.Printf("%s now has %d tokens\n", player.name, player.token)
}

// Main game loop
func (g *Game) playGame() {
	if len(g.players) < 3 {
		fmt.Println("Not enough players to start the game!")
		return
	}

	fmt.Println("\n========== Game Started! ==========")
	g.displayGameStatus()

	currentPlayer := g.players[0]
	turn := 0

	for {
		turn++
		fmt.Printf("\n--- Turn %d ---\n", turn)

		g.takeTurn(currentPlayer)
		g.displayGameStatus()

		// Check for winner
		winner := g.checkWinner()
		if winner != nil {
			g.announceWinner(winner)
			break
		}

		// Move to next player
		currentPlayer = g.getNextPlayer(currentPlayer)
	}
}

// Announce the winner and end the game
func (g *Game) announceWinner(winner *Player) {
	fmt.Println("\n🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉")
	fmt.Printf("🎉 %s WINS THE GAME! 🎉\n", winner.name)
	fmt.Println("🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉\n")
}
