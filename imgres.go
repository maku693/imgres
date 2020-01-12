package main

import (
	"flag"
	"log"

	"github.com/maku693/imgres/resizer"
)

var height, width int
var fit, in, out string

func init() {
	flag.StringVar(&fit, "fit", string(resizer.Contain), "fitting of scaled image")
	flag.StringVar(&in, "in", "", "input file (optional)")
	flag.StringVar(&out, "out", "", "out file (optional)")
	flag.IntVar(&height, "height", 0, "max height of out file")
	flag.IntVar(&width, "width", 0, "max width of out file")

	log.SetFlags(0)
}

func main() {
	flag.Parse()

	if err := resizer.Resize(in, out, height, width, fit); err != nil {
		FatalError(err)
	}
}

func FatalError(err error) {
	log.Fatalln("fatal: " + err.Error())
}
