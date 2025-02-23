package world

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"

	"github.com/moorad/raytracing/src/geometry"
	"github.com/moorad/raytracing/src/structs"
)

var logger = log.New(os.Stderr, "", 0)

// Camera constants

type Camera struct {
	magicNumber      string
	aspectRatio      float64
	imageWidth       int
	imageHeight      int
	maxColorValue    int
	samplesPerPixel  int
	maxDepth         int
	pixelSampleScale float64
	center           structs.Vec3
	pixelDeltaU      structs.Vec3
	pixelDeltaV      structs.Vec3
	originPixelLoc   structs.Vec3
}

func (camera *Camera) Render(world geometry.Surface) {
	camera.initialize()

	fmt.Printf("%s\n%v %v\n%v\n", camera.magicNumber, camera.imageWidth, camera.imageHeight, camera.maxColorValue)

	for j := 0; j < camera.imageHeight; j++ {
		logger.Printf("Scanlines remaining: %v (%v%%)\n",
			camera.imageHeight-j,
			math.Floor((float64(j)/float64(camera.imageHeight))*100),
		)

		for i := 0; i < camera.imageWidth; i++ {
			pixelColor := structs.Vec3{
				X: 0,
				Y: 0,
				Z: 0,
			}

			for sample := 0; sample < camera.samplesPerPixel; sample++ {
				ray := camera.getRay(i, j)
				rayColor := camera.rayColor(&ray, camera.maxDepth, world)
				pixelColor = structs.VecAdd(pixelColor, rayColor.ToVec3())
			}

			resultColor := structs.ToColor(structs.VecMultScaler(pixelColor, camera.pixelSampleScale))

			resultColor.Print()
		}
	}
	logger.Println("Done")
}

func (camera *Camera) initialize() {
	camera.magicNumber = "P3"
	camera.imageWidth = 400
	camera.aspectRatio = 16.0 / 9.0
	camera.maxColorValue = 255
	camera.samplesPerPixel = 100
	camera.maxDepth = 50

	camera.imageHeight = int(float64(camera.imageWidth) / camera.aspectRatio)

	if camera.imageHeight < 1 {
		camera.imageHeight = 1
	}

	camera.pixelSampleScale = 1 / float64(camera.samplesPerPixel)

	camera.center = structs.Vec3{
		X: 0,
		Y: 0,
		Z: 0,
	}

	// Viewport dimensions
	viewportHeight := 2.0
	viewportWidth := viewportHeight * (float64(camera.imageWidth) / float64(camera.imageHeight))
	focalLength := 1

	viewportU := structs.Vec3{
		X: viewportWidth,
		Y: 0,
		Z: 0,
	}
	viewportV := structs.Vec3{
		X: 0,
		Y: -viewportHeight,
		Z: 0,
	}

	camera.pixelDeltaU = structs.VecDivScaler(viewportU, float64(camera.imageWidth))
	camera.pixelDeltaV = structs.VecDivScaler(viewportV, float64(camera.imageHeight))

	viewportUpperLeft := structs.VecSub(
		structs.VecSub(
			structs.VecSub(
				camera.center,
				structs.Vec3{
					X: 0,
					Y: 0,
					Z: float64(focalLength),
				},
			),
			structs.VecDivScaler(viewportU, 2),
		),
		structs.VecDivScaler(viewportV, 2),
	)

	camera.originPixelLoc = structs.VecAdd(viewportUpperLeft, structs.VecMultScaler(structs.VecAdd(camera.pixelDeltaU, camera.pixelDeltaV), 0.5))
}

func (camera *Camera) getRay(i int, j int) structs.Ray {

	offset := camera.sampleSquare()
	pixelSample := structs.VecAdd(
		camera.originPixelLoc,
		structs.VecAdd(
			structs.VecMultScaler(
				camera.pixelDeltaU,
				float64(i)+offset.X,
			),
			structs.VecMultScaler(
				camera.pixelDeltaV,
				float64(j)+offset.Y,
			),
		),
	)

	rayOrigin := camera.center
	rayDirection := structs.VecSub(pixelSample, rayOrigin)

	return structs.Ray{
		Origin:    rayOrigin,
		Direction: rayDirection,
	}
}

func (camera *Camera) sampleSquare() structs.Vec3 {
	return structs.Vec3{X: rand.Float64() - 0.5, Y: rand.Float64() - 0.5, Z: 0}
}

func (camera *Camera) rayColor(r *structs.Ray, depth int, world geometry.Surface) structs.Color {
	if depth <= 0 {
		return structs.Color{R: 0, G: 0, B: 0}
	}

	var rec geometry.HitRecord

	if world.Hit(r, structs.Interval{Min: 0.001, Max: math.Inf(1)}, &rec) {
		var scattered structs.Ray
		var attenuation structs.Color

		if rec.Material.Scatter(r, &rec, &attenuation, &scattered) {
			color := camera.rayColor(&scattered, depth-1, world)
			return structs.ToColor(
				structs.VecMult(attenuation.ToVec3(), color.ToVec3()),
			)
		}

		return structs.Color{
			R: 0,
			G: 0,
			B: 0,
		}

		// direction := structs.VecAdd(rec.Normal, structs.RandomOnHemisphere(rec.Normal))

		// color := camera.rayColor(&structs.Ray{
		// 	Origin:    rec.Position,
		// 	Direction: direction,
		// }, depth-1, world)

		// return structs.ToColor(
		// 	structs.VecMultScaler(
		// 		color.ToVec3(),
		// 		0.1,
		// 	),
		// )
	}

	unitDirection := structs.UnitVector(r.Direction)
	a := 0.5 * (unitDirection.Y + 1)

	return structs.ToColor(
		structs.VecAdd(
			structs.VecMultScaler(structs.Vec3{
				X: 1.0,
				Y: 1.0,
				Z: 1.0,
			}, 1.0-a),
			structs.VecMultScaler(structs.Vec3{
				X: 0.5,
				Y: 0.7,
				Z: 1.0,
			}, a),
		),
	)
}
