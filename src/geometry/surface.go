package geometry

import (
	"github.com/moorad/raytracing/src/structs"
)

type HitRecord struct {
	Position  structs.Vec3
	Normal    structs.Vec3
	T         float64
	FrontFace bool
	Material  Material
}

func (rec *HitRecord) SetFrontNormal(ray *structs.Ray, outwardNormal *structs.Vec3) {
	rec.FrontFace = structs.VecDot(ray.Direction, *outwardNormal) < 0

	if rec.FrontFace {
		rec.Normal = *outwardNormal
	} else {
		rec.Normal = outwardNormal.Negate()

	}
}

type Surface interface {
	Hit(ray *structs.Ray, rayT structs.Interval, rec *HitRecord) bool
}
