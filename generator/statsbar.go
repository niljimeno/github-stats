package generator

import (
	"image"
	"image/color"
	"log"

	"github.com/niljimeno/github-stats/stats"
)

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

func (s *StatsBar) DrawStats() {
	stats := []stats.Stat{
		{
			Language:   "go",
			Percentage: 60,
		},
		{
			Language:   "haskell",
			Percentage: 30,
		},
		{
			Language: "erlang",
			Amount:   10,
		},
	}

	for i, v := range stats {
		cellSize := 48
		textSize := 24

		row := i / 3
		col := i % 3

		x := col * 160
		y := row * (cellSize + 20)

		err := DrawIcon(DrawIconOps{
			X:      x,
			Y:      y,
			Screen: s.Img,
			Size:   cellSize,
			Name:   v.Language,
		})

		if err != nil {
			log.Panic(err)
		}

		DrawText(DrawTextOps{
			X: x + cellSize,
			Y: y + textSize,

			Screen: s.Img,
			Size:   textSize,
			Text:   v.Language,
		})
	}
}
