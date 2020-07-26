package main

import (
	"log"
	"math/rand"
	"testing"
)

func TestPlayerRaceCondition(t *testing.T) {
	world, err := NewWorld("scene1.json")

	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			world.Play()
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			world.Marshal()
			if err != nil {
				log.Fatal(err)
			}
		}
	}()

	for i := 0; i < 10; i++ {
		var id = world.JoinNewPlayer()

		go func() {

			for i := 0; i < 10; i++ {
				world.ApplyPlayerMovement(id, Movement{
					ArrowUp:    rand.Intn(2) == 0,
					ArrowDown:  rand.Intn(2) == 0,
					ArrowLeft:  rand.Intn(2) == 0,
					ArrowRight: rand.Intn(2) == 0,
					Space:      rand.Intn(2) == 0})
			}
		}()

	}
}
