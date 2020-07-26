package main

import (
	"fmt"
	"testing"
)

func TestSceneJsonSerializationDeserialization(t *testing.T) {
	scene, err := NewScene("scene1.json")
	if err != nil {
		t.Fatal(err)
	}

	serializedScene, err := scene.Marshal()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("serialized scene:%v", serializedScene)
}
