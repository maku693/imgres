package main

import (
	"flag"
	"image"
	"log"
)

var height, width int
var fit, in, out string

func init() {
	flag.StringVar(&fit, "fit", string(Contain), "fitting of scaled image")
	flag.StringVar(&in, "in", "", "input file (optional)")
	flag.StringVar(&out, "out", "", "out file (optional)")
	flag.IntVar(&height, "height", 0, "max height of out file")
	flag.IntVar(&width, "width", 0, "max width of out file")

	log.SetFlags(0)
}

func main() {
	flag.Parse()

	inFile, err := InFile(in)
	if err != nil {
		FatalError(err)
	}

	outFile, err := OutFile(out)
	if err != nil {
		FatalError(err)
	}

	src, format, err := image.Decode(inFile)
	if err != nil {
		FatalError(err)
	}

	dest, err := Scale(src, image.Pt(width, height), Fit(fit))
	if err != nil {
		FatalError(err)
	}

	if err := Encode(outFile, dest, format); err != nil {
		FatalError(err)
	}
}

func FatalError(err error) {
	log.Fatalln("fatal: " + err.Error())
}
