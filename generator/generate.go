package generator

import (
	"image"
	"image/color"
	"image/png"
	"os"
	// "golang.org/x/image/font"
	// "golang.org/x/image/font/basicfont"
	// "golang.org/x/image/math/fixed"
)

func Generate(path string) {
	f, _ := os.Create(path)
	png.Encode(f, doubleScale(newImage()))
}

type StatsBar struct {
	Width  int
	Height int

	Img     *image.RGBA
	Palette *Palette
}

type Palette struct {
	Bg color.Color
	Fg color.Color
}

func newImage() *image.RGBA {
	width := 500
	height := 250

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
	statsBar.DrawText()
	return statsBar.Img
}

func (s *StatsBar) Fill() {
	width := s.Width
	height := s.Height

	for x := range width {
		for y := range height {
			s.Img.Set(x, y, s.Palette.Bg)
		}
	}
}

func (s *StatsBar) DrawBar() {
	width := s.Width * 2 / 3
	height := 10

	marginX := (s.Width - width) / 2
	marginY := (s.Height - height) / 3

	for x := range width {
		for y := range height {
			s.Img.Set(marginX+x, marginY+y, s.Palette.Fg)
		}
	}
}

func (s *StatsBar) DrawText() {
	width := s.Width * 2 / 3
	height := 10

	marginX := (s.Width - width) / 2
	marginY := (s.Height - height) * 2 / 3

	DrawTextSimplified(marginX, marginY, s.Img)

	// col := color.RGBA{0xff, 0x55, 0, 0xff}
	// point := fixed.Point26_6{
	// 	X: fixed.I(marginX),
	// 	Y: fixed.I(marginY),
	// }

	// d := &font.Drawer{
	// 	Dst:  s.Img,
	// 	Src:  image.NewUniform(col),
	// 	Face: basicfont.Face7x13,
	// 	Dot:  point,
	// }
	// d.DrawString("hello")
}

func doubleScale(source *image.RGBA) *image.RGBA {
	initialWidth := source.Rect.Dx()
	initialHeight := source.Rect.Dy()

	upLeft := image.Point{0, 0}
	lowRight := image.Point{initialWidth * 2, initialHeight * 2}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	for x := range initialWidth {
		for y := range initialHeight {
			col := source.At(x, y)

			img.Set(2*x, 2*y, col)
			img.Set(2*x, 2*y+1, col)
			img.Set(2*x+1, 2*y, col)
			img.Set(2*x+1, 2*y+1, col)
		}
	}

	return img
}
