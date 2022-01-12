package main

import (
	"gitlab.alibaba-inc.com/cos/kdm/pkg/log"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"time"
)

func main() {
	start := time.Now()
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	f, _ := os.Create("image.png")
	png.Encode(f, img) // NOTE: ignoring errors
	log.Infof("%d Milliseconds", time.Now().Sub(start).Milliseconds())
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
