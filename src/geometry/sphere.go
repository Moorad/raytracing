package geometry

import (
	"math"

	"github.com/moorad/raytracing/src/structs"
)

type Sphere struct {
	center structs.Vec3
	radius float64
}

func NewSphere(center structs.Vec3, radius float64) *Sphere {
	return &Sphere{
		radius: math.Max(0, radius),
		center: center,
	}
}

func (s Sphere) Hit(ray *structs.Ray, rayT structs.Interval, rec *HitRecord) bool {
	oc := structs.VecSub(s.center, ray.Origin)
	a := ray.Direction.LengthSqaured()
	h := structs.VecDot(ray.Direction, oc)
	c := oc.LengthSqaured() - s.radius*s.radius
	discriminant := h*h - a*c

	if discriminant < 0 {
		return false
	}

	sqrtD := math.Sqrt(discriminant)

	root := (h - sqrtD) / a

	if !rayT.Surrounds(root) {
		root = (h + sqrtD) / a
		if !rayT.Surrounds(root) {
			return false
		}
	}

	rec.T = root
	rec.Position = ray.At(rec.T)
	outwardNormal := structs.VecDivScaler(structs.VecSub(rec.Position, s.center), s.radius)
	rec.SetFrontNormal(ray, &outwardNormal)

	return true
}
