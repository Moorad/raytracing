package structs

import (
	"fmt"
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

func ToColor(v Vec3) Color {
	return Color{
		R: v.X,
		G: v.Y,
		B: v.Z,
	}
}

func (c *Color) Print() {
	interval := Interval{
		Min: 0,
		Max: 0.999,
	}

	scaledR := int32(interval.Clamp(c.R) * 255)
	scaledG := int32(interval.Clamp(c.G) * 255)
	scaledB := int32(interval.Clamp(c.B) * 255)

	fmt.Printf("%v %v %v\n", scaledR, scaledG, scaledB)

}
