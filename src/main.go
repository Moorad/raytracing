package main

import (
	"github.com/moorad/raytracing/src/geometry"
	"github.com/moorad/raytracing/src/structs"
	"github.com/moorad/raytracing/src/world"
)

func main() {

	// World
	var objects geometry.SurfaceList

	ground := geometry.Lambertain{Albedo: structs.Color{R: 0.8, G: 0.8, B: 0.0}}
	center := geometry.Lambertain{Albedo: structs.Color{R: 0.1, G: 0.2, B: 0.5}}
	matLeft := geometry.Metal{Albedo: structs.Color{R: 0.8, G: 0.8, B: 0.8}, Fuzz: 0.3}
	matRight := geometry.Metal{Albedo: structs.Color{R: 0.8, G: 0.6, B: 0.2}, Fuzz: 1.0}

	objects.Add(geometry.NewSphere(structs.Vec3{X: 0, Y: -100.5, Z: -1}, 100, ground))
	objects.Add(geometry.NewSphere(structs.Vec3{X: 0, Y: 0, Z: -1.2}, 0.5, center))
	objects.Add(geometry.NewSphere(structs.Vec3{X: -1.0, Y: 0, Z: -1.0}, 0.5, matLeft))
	objects.Add(geometry.NewSphere(structs.Vec3{X: 1.0, Y: 0, Z: -1.0}, 0.5, matRight))
	// logger.Printf("Generation info:\n\twidth: %vpx\n\theight: %vpx\n\tviewport width: %v\n\tviewport height: %vvocal length: %v", imageWidth, imageHeight, viewportWidth, viewportHeight, vocalLength)

	camera := world.Camera{}

	camera.Render(objects)

}
