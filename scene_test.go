package main

import (
	"fmt"
	"testing"

	"github.com/ByteArena/box2d"
)

func TestSceneJsonSerializationDeserialization(t *testing.T) {
	world := box2d.MakeB2World(box2d.MakeB2Vec2(0, 0))
	scene, err := NewScene("scene1.json", &world)
	if err != nil {
		t.Fatal(err)
	}

	serializedScene, err := scene.Marshal()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("serialized scene:%v", serializedScene)
}
