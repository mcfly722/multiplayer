package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/ByteArena/box2d"
)

// Point - point coordinates
type Point struct {
	X float64
	Y float64
}

// TileObject - static Object on Scene Layer
type TileObject struct {
	ID      int
	Name    string
	Type    string
	X       float64
	Y       float64
	Width   float64
	Height  float64
	Polygon []Point
}

// Layer - static Scene Layer
type Layer struct {
	ID      int
	Name    string
	Objects []TileObject
}

// Scene - static scene
type Scene struct {
	Layers []Layer
}

// ApplySceneToBox2dWorld - function applies scent to box2d world
func (scene *Scene) ApplySceneToBox2dWorld(world *box2d.B2World) {
	for _, layer := range scene.Layers {
		for _, object := range layer.Objects {
			if len(object.Polygon) > 0 {

				vs := make([]box2d.B2Vec2, len(object.Polygon)+1)
				for i, p := range object.Polygon {
					vs[i].Set((object.X+p.X)/PixelsPerMeter, (object.Y+p.Y)/PixelsPerMeter)
				}
				vs[len(object.Polygon)].Set(object.X/PixelsPerMeter, object.Y/PixelsPerMeter) // last vertice is not closed as loop, thats why this final vertice

				bd := box2d.MakeB2BodyDef()
				bd.Type = box2d.B2BodyType.B2_staticBody
				body := world.CreateBody(&bd)

				shape := box2d.MakeB2ChainShape()
				shape.CreateChain(vs, len(vs))

				fd := box2d.MakeB2FixtureDef()
				fd.Shape = &shape
				fd.Density = 1
				fd.Restitution = 0
				fd.Friction = 0
				fd.Filter = box2d.MakeB2Filter()

				body.CreateFixtureFromDef(&fd)
			}
		}
	}
}

// NewScene constructor
func NewScene(fileName string, world *box2d.B2World) (*Scene, error) {

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

	scene.ApplySceneToBox2dWorld(world)

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
