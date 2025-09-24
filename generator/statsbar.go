package generator

import (
	"image"
	"image/color"
	"log"

	"github.com/niljimeno/github-stats/generator/languages"
)

type StatsBar struct {
	Width  int
	Height int

	InnerMarginX int
	InnerMarginY int

	InnerWidth  int
	InnerHeight int

	Img     *image.RGBA
	Palette *Palette
}

type Palette struct {
	Bg color.Color
	Fg color.Color
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
	width := s.InnerWidth
	height := 10

	for x := range width {
		for y := range height {
			s.Img.Set(s.InnerMarginX+x, s.InnerMarginY+y, s.Palette.Fg)
		}
	}
}

func (s *StatsBar) DrawStats() {
	languages.SetUp()
	s.DrawBar()
	err := s.DrawInnerStats()
	if err != nil {
		log.Panic(err)
	}
}
