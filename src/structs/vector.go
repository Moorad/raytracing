package structs

import (
	"math"
	"math/rand"
)

type Vec3 struct {
	X, Y, Z float64
}

func (v *Vec3) Negate() Vec3 {
	return Vec3{
		X: -v.X,
		Y: -v.Y,
		Z: -v.Z,
	}
}

func (v *Vec3) Add(value float64) *Vec3 {
	v.Z += value
	v.Y += value
	v.Z += value
	return v
}

func (v *Vec3) Subtract(value float64) *Vec3 {
	v.X -= value
	v.Y -= value
	v.Z -= value
	return v
}

func (v *Vec3) Multiply(value float64) *Vec3 {
	v.X *= value
	v.Y *= value
	v.Z *= value
	return v
}

func (v *Vec3) Divide(value float64) *Vec3 {
	v.Multiply(1 / value)
	return v
}

func (v *Vec3) Length() float64 {
	return math.Sqrt(float64(v.LengthSqaured()))
}

func (v *Vec3) LengthSqaured() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func VecAdd(v1 Vec3, v2 Vec3) Vec3 {
	return Vec3{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
		Z: v1.Z + v2.Z,
	}
}

func VecSub(v1 Vec3, v2 Vec3) Vec3 {
	return Vec3{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
		Z: v1.Z - v2.Z,
	}
}

func VecMult(v1 Vec3, v2 Vec3) Vec3 {
	return Vec3{
		X: v1.X * v2.X,
		Y: v1.Y * v2.Y,
		Z: v1.Z * v2.Z,
	}
}

func VecDiv(v1 Vec3, v2 Vec3) Vec3 {
	return Vec3{
		X: v1.X / v2.X,
		Y: v1.Y / v2.Y,
		Z: v1.Z / v2.Z,
	}
}

func VecMultScaler(v Vec3, value float64) Vec3 {
	return Vec3{
		X: v.X * value,
		Y: v.Y * value,
		Z: v.Z * value,
	}
}

func VecDivScaler(v Vec3, value float64) Vec3 {
	return VecMultScaler(v, 1/value)
}

func VecDot(v1 Vec3, v2 Vec3) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

func VecCross(v1 Vec3, v2 Vec3) Vec3 {
	return Vec3{
		X: v1.Y*v2.Z - v1.Z*v2.Y,
		Y: v1.Z*v2.X - v1.X*v2.Z,
		Z: v1.X*v2.Y - v1.Y*v2.Z,
	}
}

func UnitVector(v Vec3) Vec3 {
	return VecDivScaler(v, v.Length())
}

func RandomVector(min float64, max float64) Vec3 {
	return Vec3{
		X: min + (max-min)*rand.Float64(),
		Y: min + (max-min)*rand.Float64(),
		Z: min + (max-min)*rand.Float64(),
	}
}

func RandomUnitVector() Vec3 {
	for {
		p := RandomVector(-1, 1)
		lenSq := p.LengthSqaured()
		if 1e-160 < lenSq && lenSq <= 1 {
			return VecDivScaler(p, math.Sqrt(lenSq))
		}
	}
}

func RandomOnHemisphere(normal Vec3) Vec3 {
	unitSphere := RandomUnitVector()

	if VecDot(unitSphere, normal) > 0 {
		return unitSphere
	}

	return unitSphere.Negate()
}
