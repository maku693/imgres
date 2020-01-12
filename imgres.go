package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"math"
	"os"

	"golang.org/x/image/draw"
)

var in, out string
var width, height int

func init() {
	flag.StringVar(&in, "in", "", "input file (optional)")
	flag.StringVar(&out, "out", "", "out file (optional)")
	flag.IntVar(&width, "width", 0, "max width of out file")
	flag.IntVar(&height, "height", 0, "max height of out file")
}

var logger = log.New(os.Stderr, "", 0)

func main() {
	flag.Parse()

	if width == 0 && height == 0 {
		logger.Fatalln(fmt.Errorf("either -width or -height must be specified and non-zero").Error())
	}

	reader := os.Stdin
	if in != "" {
		file, err := os.Open(in)
		if err != nil {
			logger.Fatalln(err.Error())
		}
		reader = file
	}

	writer := os.Stdout
	if out != "" {
		file, err := os.OpenFile(out, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			logger.Fatalln(err.Error())
		}
		writer = file
	}

	src, format, err := image.Decode(reader)
	if err != nil {
		logger.Fatalln(err.Error())
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

	var encode func(io.Writer, image.Image) error
	switch format {
	case "png":
		encode = func(w io.Writer, i image.Image) error { return png.Encode(w, i) }
	case "jpeg":
		encode = func(w io.Writer, i image.Image) error { return jpeg.Encode(w, i, nil) }
	case "gif":
		encode = func(w io.Writer, i image.Image) error { return gif.Encode(w, i, nil) }
	default:
		logger.Fatalln(fmt.Errorf("invalid format: %s", format).Error())
	}

	err = encode(writer, dest)
	if err != nil {
		logger.Fatalln(err.Error())
	}
}
