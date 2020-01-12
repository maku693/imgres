package main

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
)

func Encode(w io.Writer, i image.Image, format string) error {
	switch format {
	case "png":
		return png.Encode(w, i)
	case "jpeg":
		return jpeg.Encode(w, i, nil)
	case "gif":
		return gif.Encode(w, i, nil)
	}
	return fmt.Errorf("encode: invalid format: %s", format)
}
