package generator

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func Generate(path string) {
	f, _ := os.Create(path)
	png.Encode(f, newImage())
}

func newImage() *image.RGBA {
	width := 800
	height := 300

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	statsBar := StatsBar{
		Width:  width,
		Height: height,
		Img:    img,
		Palette: &Palette{
			Bg: color.RGBA{0x22, 0x22, 0x22, 0xff},
			Fg: color.RGBA{0xff, 0xff, 0x44, 0xff},
		},
	}

	statsBar.Fill()
	statsBar.DrawBar()
	statsBar.DrawStats()
	return statsBar.Img
}
