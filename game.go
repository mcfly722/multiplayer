package main

import (
	"encoding/json"
	"math/rand"
	"sync"
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

// SafeWorld - protected with mutex world
type SafeWorld struct {
	Scene      string
	Players    map[int]Player
	playersMux sync.Mutex
}

// Players map
var safeWorld = SafeWorld{Players: make(map[int]Player)}

// NewPlayer constructor
func NewPlayer() int {

	safeWorld.playersMux.Lock()

	safeWorld.Players[lastPlayerID] = Player{
		SpriteSetNum: rand.Intn(2),
		X:            (float32)(rand.Intn(200) - 100),
		Y:            (float32)(rand.Intn(200) - 100),
		SpeedX:       0,
		SpeedY:       0}

	lastPlayerID++

	safeWorld.playersMux.Unlock()

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

	safeWorld.playersMux.Lock()

	if player, ok := safeWorld.Players[playerID]; ok {
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
		safeWorld.Players[playerID] = player
	}

	safeWorld.playersMux.Unlock()

}

// PlayGame - update scene
func PlayGame() {
	for {

		safeWorld.playersMux.Lock()

		for id, player := range safeWorld.Players {
			player.X += player.SpeedX
			player.Y += player.SpeedY
			safeWorld.Players[id] = player
		}

		safeWorld.playersMux.Unlock()

		time.Sleep(1000 / 60 * time.Millisecond)
	}
}

//GetSerializedWorld - return serialized players
func GetSerializedWorld() ([]byte, error) {
	safeWorld.Scene = "scene1.json"
	safeWorld.playersMux.Lock()
	result, err := json.Marshal(&safeWorld)
	safeWorld.playersMux.Unlock()
	return result, err
}
