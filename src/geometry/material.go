package geometry

import (
	"github.com/moorad/raytracing/src/structs"
)

type Material interface {
	Scatter(ray *structs.Ray, rec *HitRecord, attenuation *structs.Color, scattered *structs.Ray) bool
}

type Lambertain struct {
	Albedo structs.Color
}

func (l Lambertain) Scatter(ray *structs.Ray, rec *HitRecord, attenuation *structs.Color, scattered *structs.Ray) bool {
	scatterDirection := structs.VecMult(rec.Normal, structs.RandomUnitVector())

	if scatterDirection.NearZero() {
		scatterDirection = rec.Normal
	}

	*scattered = structs.Ray{
		Origin:    rec.Position,
		Direction: scatterDirection,
	}
	*attenuation = l.Albedo
	return true
}

type Metal struct {
	Albedo structs.Color
	Fuzz   float64
}

func (m Metal) Scatter(ray *structs.Ray, rec *HitRecord, attenuation *structs.Color, scattered *structs.Ray) bool {
	reflected := structs.Reflect(ray.Direction, rec.Normal)
	reflected = structs.VecAdd(
		structs.UnitVector(reflected),
		structs.VecMultScaler(structs.RandomUnitVector(), m.Fuzz),
	)
	*scattered = structs.Ray{
		Origin:    rec.Position,
		Direction: reflected,
	}
	*attenuation = m.Albedo
	return structs.VecDot(scattered.Direction, rec.Normal) > 0
}
