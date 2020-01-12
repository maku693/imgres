package resizer

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
)

func Resize(in string, out string, height int, width int, fit string) error {
	inFile, err := InFile(in)
	if err != nil {
		return err
	}

	outFile, err := OutFile(out)
	if err != nil {
		return err
	}

	cfg, format, err := image.DecodeConfig(inFile)
	if err != nil {
		return err
	}

	_, err = inFile.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	size, err := FitSize(
		image.Pt(cfg.Width, cfg.Height),
		image.Pt(width, height),
		Fit(fit),
	)
	if err != nil {
		return err
	}
	bounds := image.Rectangle{Max: size}

	switch format {
	case "gif":
		return ScaleGIF(inFile, outFile, bounds)
	case "jpeg":
		return ScaleJPEG(inFile, outFile, bounds)
	case "png":
		return ScalePNG(inFile, outFile, bounds)
	default:
		return fmt.Errorf("invalid format: %s", format)
	}
}
