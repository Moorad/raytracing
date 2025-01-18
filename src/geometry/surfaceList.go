package geometry

import "github.com/moorad/raytracing/src/structs"

type SurfaceList struct {
	Surfaces []Surface
}

func (sl *SurfaceList) Add(surface Surface) {
	sl.Surfaces = append(sl.Surfaces, surface)
}

func (sl *SurfaceList) Clear() {
	sl.Surfaces = nil
}

func (sl SurfaceList) Hit(ray *structs.Ray, rayT structs.Interval, rec *HitRecord) bool {
	var tempRec HitRecord
	isAnythingHit := false
	closest := rayT.Max

	for _, surface := range sl.Surfaces {
		if surface.Hit(ray, structs.Interval{Min: rayT.Min, Max: closest}, &tempRec) {
			isAnythingHit = true
			closest = tempRec.T
			*rec = tempRec
		}
	}

	return isAnythingHit
}
