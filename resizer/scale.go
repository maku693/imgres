package resizer

import (
	"image"

	"golang.org/x/image/draw"
)

func Scale(src image.Image, size image.Point, fit Fit) (image.Image, error) {
	bounds, err := FitRect(src.Bounds(), size, fit)
	if err != nil {
		return nil, err
	}

	img := image.NewRGBA(bounds)
	draw.CatmullRom.Scale(img, img.Bounds(), src, src.Bounds(), draw.Over, nil)

	return img, nil
}
