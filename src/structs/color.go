package structs

import (
	"fmt"
	"math"
)

type Color struct {
	R, G, B float64
}

func (c *Color) ToVec3() Vec3 {
	return Vec3{
		X: c.R,
		Y: c.G,
		Z: c.B,
	}
}

func (c *Color) Print() {
	interval := Interval{
		Min: 0,
		Max: 0.999,
	}

	r := linearToGamma(c.R)
	g := linearToGamma(c.G)
	b := linearToGamma(c.B)

	scaledR := int32(interval.Clamp(r) * 255)
	scaledG := int32(interval.Clamp(g) * 255)
	scaledB := int32(interval.Clamp(b) * 255)

	fmt.Printf("%v %v %v\n", scaledR, scaledG, scaledB)

}

func ToColor(v Vec3) Color {
	return Color{
		R: v.X,
		G: v.Y,
		B: v.Z,
	}
}

func linearToGamma(linearValue float64) float64 {
	if linearValue > 0 {
		return math.Sqrt(linearValue)
	}

	return 0
}
