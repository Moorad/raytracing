package geometry

import (
	"github.com/moorad/raytracing/src/structs"
)

type Material interface {
	Scatter(ray *structs.Ray, rec *HitRecord, attenuation *structs.Color, scattered *structs.Ray) bool
}
