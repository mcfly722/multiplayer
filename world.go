package main

import (
	"encoding/json"
	"math/rand"
	"sync"

	"github.com/ByteArena/box2d"
)

// RPG Maker sprites
// https://www.deviantart.com/rpg-maker-artists/gallery/25208345/RPG-Maker-Sprites

// PixelsPerMeter for box2d
const PixelsPerMeter = 20.0
const PlayersSpeed = 6.0

// Player struct
type Player struct {
	SpriteSetNum int
	X            float64
	Y            float64
	SpeedX       float64
	SpeedY       float64
	box2dBody    *box2d.B2Body
}

// World - protected with mutex world
type World struct {
	SceneName    string
	Scene        Scene `json:"-"`
	Players      map[int]Player
	playersMux   sync.Mutex
	lastPlayerID int
	box2dWorld   *box2d.B2World
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

	box2dWorld := box2d.MakeB2World(box2d.MakeB2Vec2(0, 0))

	var safeWorld = World{
		Players:    make(map[int]Player),
		box2dWorld: &box2dWorld}

	safeWorld.SceneName = sceneFileName
	scene, err := NewScene(sceneFileName, safeWorld.box2dWorld)
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

	shape := box2d.MakeB2CircleShape()
	shape.M_radius = 5.0 / PixelsPerMeter

	fd := box2d.MakeB2FixtureDef()
	fd.Shape = &shape
	fd.Density = 1
	fd.Restitution = 0
	fd.Friction = 0
	fd.Filter = box2d.MakeB2Filter()

	bd := box2d.MakeB2BodyDef()
	bd.Position.Set(50*32/PixelsPerMeter, 50*32/PixelsPerMeter)
	bd.Type = box2d.B2BodyType.B2_dynamicBody
	bd.FixedRotation = true
	bd.AllowSleep = false

	body := world.box2dWorld.CreateBody(&bd)
	body.CreateFixtureFromDef(&fd)

	player := Player{
		SpriteSetNum: rand.Intn(2),
		X:            0,
		Y:            0,
		SpeedX:       0,
		SpeedY:       0,
	}
	player.box2dBody = body

	//	player.box2dBody.SetUserData(world.lastPlayerID)

	world.Players[world.lastPlayerID] = player

	world.lastPlayerID++

	world.playersMux.Unlock()

	return world.lastPlayerID - 1
}

// ApplyPlayerMovement apply player movement
func (world *World) ApplyPlayerMovement(playerID int, movement Movement) {
	world.playersMux.Lock()

	if player, ok := world.Players[playerID]; ok {
		if movement.ArrowUp {
			player.SpeedY = -PlayersSpeed
		}
		if movement.ArrowDown {
			player.SpeedY = PlayersSpeed
		}
		if movement.ArrowLeft {
			player.SpeedX = -PlayersSpeed
		}
		if movement.ArrowRight {
			player.SpeedX = PlayersSpeed
		}

		if (!movement.ArrowUp) && (!movement.ArrowDown) {
			player.SpeedY = 0
		}
		if (!movement.ArrowLeft) && (!movement.ArrowRight) {
			player.SpeedX = 0
		}

		player.box2dBody.SetLinearVelocity(box2d.MakeB2Vec2(player.SpeedX, player.SpeedY))
		world.Players[playerID] = player
	}
	world.playersMux.Unlock()
}

// Play - update world
func (world *World) Play() {
	world.playersMux.Lock()

	timeStep := 1.0 / 60.0
	velocityIterations := 10
	positionIterations := 10

	world.box2dWorld.Step(timeStep, velocityIterations, positionIterations)

	for id, player := range world.Players {
		var position = player.box2dBody.GetPosition()
		//		log.Printf("%v -(%v,%v)", id, position.X, position.Y)

		// players inertia
		player.box2dBody.SetLinearVelocity(box2d.MakeB2Vec2(player.SpeedX, player.SpeedY))

		player.X = position.X * PixelsPerMeter
		player.Y = position.Y * PixelsPerMeter
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
