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
	return math.Min(float64(tgt.X)/float64(src.X), float64(tgt.Y)/float64(src.Y))
}

func coverer(src, tgt image.Point) float64 {
	return math.Max(float64(tgt.X)/float64(src.X), float64(tgt.Y)/float64(src.Y))
}

func FitRect(r image.Rectangle, size image.Point, fit Fit) (image.Rectangle, error) {
	var f fitter
	switch fit {
	case Contain:
		f = container
	case Cover:
		f = coverer
	default:
		return image.Rectangle{}, fmt.Errorf("invalid fit: %s", fit)
	}

	rbounds := r.Bounds()
	rsize := r.Size()

	scale := f(rsize, size)

	return image.Rect(
		int(float64(rbounds.Min.X)*scale),
		int(float64(rbounds.Min.Y)*scale),
		int(float64(rbounds.Max.X)*scale),
		int(float64(rbounds.Max.Y)*scale),
	), nil
}
