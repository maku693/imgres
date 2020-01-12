package main

import (
	"flag"
	"image"
	"log"
	"os"
)

var in, out string
var width, height int

func init() {
	flag.StringVar(&in, "in", "", "input file (optional)")
	flag.StringVar(&out, "out", "", "out file (optional)")
	flag.IntVar(&width, "width", 0, "max width of out file")
	flag.IntVar(&height, "height", 0, "max height of out file")
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

	dest, err := Scale(src, width, height)
	if err != nil {
		FatalError(err)
	}

	if err := Encode(outFile, dest, format); err != nil {
		FatalError(err)
	}
}

var logger = log.New(os.Stderr, "", 0)

func FatalError(err error) {
	logger.Fatalln("fatal: " + err.Error())
}
