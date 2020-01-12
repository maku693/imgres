package resizer

import "image"

func Resize(height int, width int, fit string, in string, out string) error {
	inFile, err := InFile(in)
	if err != nil {
		return err
	}

	outFile, err := OutFile(out)
	if err != nil {
		return err
	}

	img, format, err := image.Decode(inFile)
	if err != nil {
		return err
	}

	scaled, err := Scale(img, image.Pt(width, height), Fit(fit))
	if err != nil {
		return err
	}

	if err := Encode(outFile, scaled, format); err != nil {
		return err
	}
	return nil
}
