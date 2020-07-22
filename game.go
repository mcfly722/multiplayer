package main

import "math/rand"

// Movement - pressed buttons on client
type Movement struct {
	ArrowUp    bool
	ArrowDown  bool
	ArrowLeft  bool
	ArrowRight bool
	Space      bool
}

// RPG Maker sprites
// https://www.deviantart.com/rpg-maker-artists/gallery/25208345/RPG-Maker-Sprites

// Player struct
type Player struct {
	spriteSetNum int
	x            int
	y            int
}

var lastPlayerID float64

// Players map
var Players = map[float64]Player{}

// NewPlayer constructor
func NewPlayer() float64 {

	Players[lastPlayerID] = Player{
		spriteSetNum: rand.Intn(3),
		x:            rand.Intn(200) - 100,
		y:            rand.Intn(200) - 100}

	lastPlayerID++
	return lastPlayerID - 1
}
