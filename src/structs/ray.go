package structs

type Ray struct {
	Origin    Vec3
	Direction Vec3
}

func (r *Ray) At(position float64) Vec3 {
	return VecAdd(r.Origin, VecMultScaler(r.Direction, position))
}
