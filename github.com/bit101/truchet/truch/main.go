package main

import (
	"flag"
	"fmt"
	"math"
	"os"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/blmath"
	"github.com/bit101/go-cairo"
	"github.com/bit101/truchet/truchet"
)

var image *cairo.Surface
var width float64
var height float64

var pattern = truchet.Pattern0

func main() {
	inFile := flag.String("i", "", "The png image file to read from.")
	outFile := flag.String("o", "out.png", "The output png file to write.")
	tileSizeInt := flag.Int("t", 10, "The tile size of each Truchet tile.")
	pattern := flag.String("p", "a", "Which pattern to use.")
	flag.Parse()

	tileSize := float64(*tileSizeInt)

	image, _ = cairo.NewSurfaceFromPNG(*inFile)
	if image == nil {
		fmt.Println("Error loading png")
		os.Exit(1)
	}
	data := image.GetData()

	width = float64(image.GetWidth())
	height = float64(image.GetHeight())
	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)

	for y := 0.0; y < height; y += tileSize {
		for x := 0.0; x < width; x += tileSize {
			brightness := getBrightnessFromPNG(x, y, tileSize, data)
			truchet.Truchet(surface, *pattern, x, y, tileSize, brightness)
		}
	}

	surface.WriteToPNG(*outFile)
}

func getBrightness(x, y float64) float64 {
	dx := x - width/2
	dy := y - height/2
	dist := math.Sqrt(dx*dx + dy*dy)
	n := math.Sin(dist * 0.05)
	return blmath.Map(n, -1, 1, 0, 1)
}

func getBrightnessFromPNG(x, y, t float64, data []byte) float64 {
	val := 0
	for yy := int(y); yy < int(y+t); yy++ {
		for xx := int(x); xx < int(x+t); xx++ {
			xxx := int(math.Min(float64(xx), width-1))
			yyy := int(math.Min(float64(yy), height-1))
			index := (yyy*image.GetWidth() + xxx) * 4
			val += int(data[index])
			val += int(data[index+1])
			val += int(data[index+2])
		}
	}

	return 1.0 - float64(val)/(3*t*t)/255.0

}
