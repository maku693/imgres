package resizer

import (
	"fmt"
	"image"
	"math"
)

type Fit string

const (
	Contain Fit = "contain"
	Cover   Fit = "cover"
)

type fitter func(src, tgt image.Point) float64

func container(src, tgt image.Point) float64 {
	x := float64(tgt.X) / float64(src.X)
	y := float64(tgt.Y) / float64(src.Y)
	if x == 0 {
		return y
	}
	if y == 0 {
		return x
	}
	return math.Min(x, y)
}

func coverer(src, tgt image.Point) float64 {
	return math.Max(float64(tgt.X)/float64(src.X), float64(tgt.Y)/float64(src.Y))
}

func FitSize(src image.Point, tgt image.Point, fit Fit) (image.Point, error) {
	var f fitter
	switch fit {
	case Contain:
		f = container
	case Cover:
		f = coverer
	default:
		return image.Point{}, fmt.Errorf("invalid fit: %s", fit)
	}

	scale := f(src, tgt)

	return image.Pt(
		int(float64(src.X)*scale),
		int(float64(src.Y)*scale),
	), nil
}
