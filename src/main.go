package main

import (
	"github.com/moorad/raytracing/src/geometry"
	"github.com/moorad/raytracing/src/structs"
	"github.com/moorad/raytracing/src/world"
)

func main() {

	// World
	var objects geometry.SurfaceList

	objects.Add(geometry.NewSphere(structs.Vec3{X: 0, Y: 0, Z: -1}, 0.5))
	objects.Add(geometry.NewSphere(structs.Vec3{X: 0, Y: -100.5, Z: -1}, 100))
	// logger.Printf("Generation info:\n\twidth: %vpx\n\theight: %vpx\n\tviewport width: %v\n\tviewport height: %vvocal length: %v", imageWidth, imageHeight, viewportWidth, viewportHeight, vocalLength)

	camera := world.Camera{}

	camera.Render(objects)

}
