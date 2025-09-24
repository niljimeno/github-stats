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

	innerWidth := width * 3 / 4
	innerHeight := height * 2 / 3

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	statsBar := StatsBar{
		Width:  width,
		Height: height,

		InnerWidth:  innerWidth,
		InnerHeight: innerHeight,

		InnerMarginX: (width - innerWidth) / 2,
		InnerMarginY: (height - innerHeight) / 2,

		Img: img,
		Palette: &Palette{
			Bg: color.RGBA{0x22, 0x22, 0x22, 0xff},
			Fg: color.RGBA{0xff, 0xff, 0xff, 0xff},
		},
	}

	statsBar.Fill()
	statsBar.DrawBar()
	statsBar.DrawStats()
	return statsBar.Img
}
