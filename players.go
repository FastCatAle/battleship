package main

import (
	"fmt"
)

type Ships struct {
	Size        [4]int
	Coordinates [4]string
	Position    [4]rune
}

type Player struct {
	Name       string
	OwnBoard   [10][10]string
	EnemyBoard [10][10]string
	Ships      Ships
}

func (p *Player) createPlayersName(playernumber uint8) {
	if playernumber == 1 {
		fmt.Printf("Enter the player #%v name: ", playernumber)
		fmt.Scanln(&p.Name)
	} else {
		fmt.Printf("Enter the player #%v name: ", playernumber)
		fmt.Scanln(&p.Name)
	}
}

func (p *Player) createShips() {
	ships := &p.Ships
	for i := 0; i < 4; i++ {
		if i == 0 {
			ships.Size[i] = i + 2
			fmt.Printf("Enter the coordinates for ship #%v of size %v: ", i+1, ships.Size[i])
			fmt.Scanln(&p.Ships.Coordinates[i])
		} else {
			ships.Size[i] = i + 1
			fmt.Printf("Enter the coordinates for ship #%v of size %v: ", i+1, ships.Size[i])
			fmt.Scanln(&p.Ships.Coordinates[i])
		}
	}
}

func (p *Player) createBoard() {
	own := &p.OwnBoard
	// Fill with water and add coordinates
	for x := 1; x < 10; x++ {
		xCoord := x - 1
		own[x][0] = string(65 + xCoord)
		for y := 1; y < 10; y++ {
			yCoord := y - 1
			own[0][y] = string(49 + yCoord)
			own[x][y] = "~"
		}
	}
	//p.placeShips()

	// Same logic for enemy's board
	enemy := &p.EnemyBoard
	for x := 1; x < 10; x++ {
		xCoord := x - 1
		enemy[x][0] = string(65 + xCoord)
		for y := 1; y < 10; y++ {
			yCoord := y - 1
			enemy[0][y] = string(49 + yCoord)
			enemy[x][y] = "~"
		}
	}
}
