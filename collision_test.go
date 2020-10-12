package main

import (
	"testing"

	"github.com/ByteArena/box2d"
)

func TestBox2dCollisions(t *testing.T) {

	world := box2d.MakeB2World(box2d.MakeB2Vec2(0, -10))

	// ground body
	{
		bd := box2d.MakeB2BodyDef()
		ground := world.CreateBody(&bd)
		shape := box2d.MakeB2EdgeShape()
		shape.Set(box2d.MakeB2Vec2(-20.0, 0.0), box2d.MakeB2Vec2(20.0, 0.0))
		ground.CreateFixture(&shape, 0.0)
	}

	// circle character
	bd := box2d.MakeB2BodyDef()
	bd.Position.Set(0, 10.0)
	bd.Type = box2d.B2BodyType.B2_dynamicBody
	bd.FixedRotation = true
	bd.AllowSleep = false

	body := world.CreateBody(&bd)

	shape := box2d.MakeB2CircleShape()
	shape.M_radius = 0.5

	fd := box2d.MakeB2FixtureDef()
	fd.Shape = &shape
	fd.Density = 200.0
	body.CreateFixtureFromDef(&fd)

	// Prepare for simulation. Typically we use a time step of 1/60 of a
	// second (60Hz) and 10 iterations. This provides a high quality simulation
	// in most game scenarios.
	timeStep := 1.0 / 60.0
	velocityIterations := 8
	positionIterations := 3

	for i := 0; i < 100; i++ {
		// Instruct the world to perform a single step of simulation.
		// It is generally best to keep the time step and iterations fixed.
		//runtime.Breakpoint()
		world.Step(timeStep, velocityIterations, positionIterations)
		//fmt.Printf("circle character coordinates: %v,%v\n", body.GetPosition().X, body.GetPosition().Y)
	}

}
