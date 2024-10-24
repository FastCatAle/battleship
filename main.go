package main

import "fmt"

func main() {
	fmt.Println("\n-----------------------------")
	fmt.Println("    Welcome to Battleship    ")
	fmt.Println("-----------------------------")

	var PlayerNumberOne Player
	PlayerNumberOne.createPlayersName(1)
	PlayerNumberOne.createShips()

	var PlayerNumberTwo Player
	PlayerNumberTwo.createPlayersName(2)
	PlayerNumberTwo.createShips()

	fmt.Printf("Player 1: %s\n", PlayerNumberOne.Name)
	fmt.Printf("Ship size %v\n", PlayerNumberOne.Ships)
	fmt.Printf("Player 2: %s\n", PlayerNumberTwo.Name)
	fmt.Printf("Ship size %v\n", PlayerNumberTwo.Ships)

	PlayerNumberOne.createBoard()
	fmt.Println("Your board:")
	PlayerNumberOne.printOwnBoard()

	fmt.Println("Enemy's board:")
	PlayerNumberOne.printEnemyBoard()
	//fmt.Printf("%q\n", PlayerNumberOne.OwnBoard[0][2])
	//fmt.Printf("%q\n", PlayerNumberOne.OwnBoard[2][0])
}
