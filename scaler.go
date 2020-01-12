package main

import (
	"fmt"
	"image"
	"math"

	"golang.org/x/image/draw"
)

func Scale(src image.Image, width int, height int) (image.Image, error) {
	if width == 0 && height == 0 {
		return nil, fmt.Errorf("either width or height must be specified and non-zero")
	}

	srcBounds := src.Bounds()

	srcSize := srcBounds.Size()
	xScale := float64(1)
	yScale := float64(1)
	if srcSize.X > width {
		xScale = float64(width) / float64(srcSize.X)
	}
	if srcSize.Y > height {
		yScale = float64(height) / float64(srcSize.Y)
	}
	scale := math.Min(math.Max(float64(xScale), float64(yScale)), float64(1))

	destBounds := image.Rect(
		int(float64(srcBounds.Min.X)*scale),
		int(float64(srcBounds.Min.Y)*scale),
		int(float64(srcBounds.Max.X)*scale),
		int(float64(srcBounds.Max.Y)*scale),
	)

	dest := image.NewRGBA(destBounds)
	draw.CatmullRom.Scale(dest, destBounds, src, srcBounds, draw.Over, nil)

	return dest, nil
}
