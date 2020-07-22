package main

import (
	"math/rand"
	"time"
)

// RPG Maker sprites
// https://www.deviantart.com/rpg-maker-artists/gallery/25208345/RPG-Maker-Sprites

// Player struct
type Player struct {
	SpriteSetNum int     `json:"SpriteSetNum"`
	X            float32 `json:"X"`
	Y            float32 `json:"Y"`
	SpeedX       float32 `json:"SpeedX"`
	SpeedY       float32 `json:"SpeedY"`
}

var lastPlayerID int

// Players map
var Players = map[int]Player{}

// NewPlayer constructor
func NewPlayer() int {

	Players[lastPlayerID] = Player{
		SpriteSetNum: rand.Intn(3),
		X:            (float32)(rand.Intn(200) - 100),
		Y:            (float32)(rand.Intn(200) - 100),
		SpeedX:       0,
		SpeedY:       0}

	lastPlayerID++
	return lastPlayerID - 1
}

// Movement - pressed buttons on client
type Movement struct {
	ArrowUp    bool
	ArrowDown  bool
	ArrowLeft  bool
	ArrowRight bool
	Space      bool
}

// ApplyPlayerMovement - change player speed
func ApplyPlayerMovement(playerID int, movement Movement) {
	if player, ok := Players[playerID]; ok {
		if movement.ArrowUp {
			player.SpeedY = -1
		}
		if movement.ArrowDown {
			player.SpeedY = 1
		}
		if movement.ArrowLeft {
			player.SpeedX = -1
		}
		if movement.ArrowRight {
			player.SpeedX = 1
		}

		if (!movement.ArrowUp) && (!movement.ArrowDown) {
			player.SpeedY = 0
		}
		if (!movement.ArrowLeft) && (!movement.ArrowRight) {
			player.SpeedX = 0
		}
		Players[playerID] = player
	}
}

// PlayGame - update scene
func PlayGame() {
	for {
		for id, player := range Players {
			player.X += player.SpeedX
			player.Y += player.SpeedY
			Players[id] = player
		}
		time.Sleep(1000 / 60 * time.Millisecond)
	}

}
