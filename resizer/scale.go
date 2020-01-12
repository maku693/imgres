package resizer

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"

	"golang.org/x/image/draw"
)

func ScaleGIF(r io.Reader, w io.Writer, rect image.Rectangle) error {
	src, err := gif.DecodeAll(r)
	if err != nil {
		return err
	}

	dst := &gif.GIF{
		Delay:     src.Delay,
		LoopCount: src.LoopCount,
		Disposal:  src.Disposal,
		Config: image.Config{
			ColorModel: src.Config.ColorModel,
			Width:      rect.Dx(),
			Height:     rect.Dy(),
		},
		BackgroundIndex: src.BackgroundIndex,
	}
	for _, s := range src.Image {
		d := image.NewPaletted(rect, s.Palette)
		catmullRom(d, s)
		dst.Image = append(dst.Image, d)
	}

	return gif.EncodeAll(w, dst)
}

func ScaleJPEG(r io.Reader, w io.Writer, rect image.Rectangle) error {
	src, err := jpeg.Decode(r)
	if err != nil {
		return err
	}

	dst := image.NewRGBA(rect)
	catmullRom(dst, src)

	return jpeg.Encode(w, dst, nil)
}

func ScalePNG(r io.Reader, w io.Writer, rect image.Rectangle) error {
	src, err := png.Decode(r)
	if err != nil {
		return err
	}

	dst := image.NewRGBA(rect)
	catmullRom(dst, src)

	return png.Encode(w, dst)
}

func catmullRom(dst draw.Image, src image.Image) {
	draw.CatmullRom.Scale(dst, dst.Bounds(), src, src.Bounds(), draw.Src, nil)
}
