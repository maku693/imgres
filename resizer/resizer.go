package resizer

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func Resize(in string, out string, width int, height int, fit string) error {
	if width <= 0 && height <= 0 {
		return errors.New("resizer: either width or height must be provided and non-zero")
	}

	inFile, err := InFile(in)
	if err != nil {
		return err
	}

	inBuf := &bytes.Buffer{}
	_, err = inBuf.ReadFrom(inFile)
	if err != nil {
		return err
	}
	inBytes := inBuf.Bytes()

	inReader := bytes.NewReader(inBytes)

	cfg, format, err := image.DecodeConfig(inReader)
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

	inReader.Reset(inBytes)
	outBuf := &bytes.Buffer{}

	switch format {
	case "gif":
		err = ScaleGIF(inReader, outBuf, bounds)
		if err != nil {
			return err
		}
	case "jpeg":
		err = ScaleJPEG(inReader, outBuf, bounds)
		if err != nil {
			return err
		}
	case "png":
		err = ScalePNG(inReader, outBuf, bounds)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("resizer: invalid format: %s", format)
	}

	outFile, err := OutFile(out)
	if err != nil {
		return err
	}

	_, err = outBuf.WriteTo(outFile)
	if err != nil {
		return err
	}

	return nil
}
