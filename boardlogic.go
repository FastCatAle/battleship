package main

import "fmt"

func (p *Player) printOwnBoard() {
	board := &p.OwnBoard
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			fmt.Printf("%-2v", board[x][y])
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func (p *Player) printEnemyBoard() {
	board := &p.EnemyBoard
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			fmt.Printf("%-2v", board[x][y])
		}
		fmt.Println("")
	}
	fmt.Println("")
}
