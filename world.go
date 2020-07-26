package main

import (
	"encoding/json"
	"math/rand"
	"sync"
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

// World - protected with mutex world
type World struct {
	SceneName    string
	Scene        Scene `json:"-"`
	Players      map[int]Player
	playersMux   sync.Mutex
	lastPlayerID int
}

// Movement - pressed buttons on client
type Movement struct {
	ArrowUp    bool
	ArrowDown  bool
	ArrowLeft  bool
	ArrowRight bool
	Space      bool
}

// NewWorld constructor
func NewWorld(sceneFileName string) (*World, error) {

	var safeWorld = World{Players: make(map[int]Player)}
	safeWorld.SceneName = sceneFileName
	scene, err := NewScene(sceneFileName)
	if err != nil {
		return nil, err
	}
	safeWorld.Scene = *scene
	safeWorld.lastPlayerID = 0

	return &safeWorld, nil
}

// JoinNewPlayer - join new player to world and return it Id
func (world *World) JoinNewPlayer() int {

	world.playersMux.Lock()

	world.Players[world.lastPlayerID] = Player{
		SpriteSetNum: rand.Intn(2),
		X:            (float32)(rand.Intn(200) - 100),
		Y:            (float32)(rand.Intn(200) - 100),
		SpeedX:       0,
		SpeedY:       0}

	world.lastPlayerID++

	world.playersMux.Unlock()

	return world.lastPlayerID - 1
}

// ApplyPlayerMovement apply player movement
func (world *World) ApplyPlayerMovement(playerID int, movement Movement) {
	world.playersMux.Lock()

	if player, ok := world.Players[playerID]; ok {
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
		world.Players[playerID] = player
	}

	world.playersMux.Unlock()

}

// Play - update world
func (world *World) Play() {
	world.playersMux.Lock()

	for id, player := range world.Players {
		player.X += player.SpeedX
		player.Y += player.SpeedY
		world.Players[id] = player
	}

	world.playersMux.Unlock()
}

// Marshal - return json serialization of the world
func (world *World) Marshal() ([]byte, error) {
	world.playersMux.Lock()
	result, err := json.Marshal(&world)
	world.playersMux.Unlock()
	return result, err
}
