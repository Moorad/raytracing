package structs

import "math"

type Interval struct {
	Min float64
	Max float64
}

func (i *Interval) Size() float64 {
	return i.Max - i.Min
}

func (i *Interval) Contains(value float64) bool {
	return i.Min <= value && value <= i.Max
}

func (i *Interval) Surrounds(value float64) bool {
	return i.Min < value && value < i.Max
}

func (i *Interval) Clamp(value float64) float64 {
	if value < i.Min {
		return i.Min
	}

	if value > i.Max {
		return i.Max
	}

	return value
}

func NewInterval() Interval {
	return Interval{
		Min: math.Inf(1),
		Max: math.Inf(-1),
	}
}

func NewUniverseInterval() Interval {
	return Interval{
		Min: math.Inf(-1),
		Max: math.Inf(1),
	}
}
