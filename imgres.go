package main

import (
	"flag"
	"log"

	"github.com/maku693/imgres/resizer"
)

var height, width int
var fit, in, out string

func init() {
	flag.StringVar(&fit, "fit", string(resizer.Contain), `fitting of scaled image ("contain" or "cover")`)
	flag.StringVar(&in, "in", "", "input file (optional, default stdin)")
	flag.StringVar(&out, "out", "", "output file (optional, default stdout)")
	flag.IntVar(&height, "height", 0, "max height of output file")
	flag.IntVar(&width, "width", 0, "max width of output file")

	log.SetFlags(0)
}

func main() {
	flag.Parse()

	if err := resizer.Resize(in, out, width, height, fit); err != nil {
		log.Fatalln("fatal: " + err.Error())
	}
}
