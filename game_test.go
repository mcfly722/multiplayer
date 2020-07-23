package main

import (
	"math/rand"
	"testing"
)

func TestPlayerRaceCondition(t *testing.T) {

	go PlayGame()

	go func() {
		for i := 0; i < 10; i++ {
			GetSerializedPlayers()
		}
	}()

	for i := 0; i < 10; i++ {
		var id = NewPlayer()

		go func() {

			for i := 0; i < 10; i++ {
				ApplyPlayerMovement(id, Movement{
					ArrowUp:    rand.Intn(2) == 0,
					ArrowDown:  rand.Intn(2) == 0,
					ArrowLeft:  rand.Intn(2) == 0,
					ArrowRight: rand.Intn(2) == 0,
					Space:      rand.Intn(2) == 0})
			}
		}()

	}
}
