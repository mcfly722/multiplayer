package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// TileSet - static TileSet
type TileSet struct {
	Columns    int
	Image      string
	TileHeight int
	TileWidth  int
}

// TileObject - static Object on Scene Layer
type TileObject struct {
	ID     int
	Name   string
	Type   string
	x      int
	y      int
	width  int
	height int
}

// Layer - static Scene Layer
type Layer struct {
	ID          int
	Name        string
	Compression string
	Encoding    string
	Data        string
	objects     []TileObject
}

// Scene - static scene
type Scene struct {
	Layers   []Layer
	Tilesets []TileSet
}

// NewScene constructor
func NewScene(fileName string) (*Scene, error) {

	var scene Scene

	jsonFile, err := os.Open("./static/" + fileName)
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(byteValue, &scene)
	if err != nil {
		return nil, err
	}

	return &scene, nil
}

// Marshal - serialize scene to JSON
func (scene Scene) Marshal() (string, error) {
	jsonContent, err := json.Marshal(scene)
	if err != nil {
		return "", err
	}
	return string(jsonContent), nil
}
