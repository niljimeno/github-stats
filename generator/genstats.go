package generator

import (
	"fmt"
	"image/color"

	"github.com/niljimeno/github-stats/generator/languages"
	"github.com/niljimeno/github-stats/generator/stats"
)

func (s *StatsBar) DrawInnerStats() error {
	stats := []stats.Stat{
		{
			Language:   languages.GetLanguage("go"),
			Percentage: 60,
		},
		{
			Language:   languages.GetLanguage("haskell"),
			Percentage: 30,
		},
		{
			Language:   languages.GetLanguage("erlang"),
			Percentage: 5,
		},
		{
			Language:   languages.GetLanguage("gleam"),
			Percentage: 5,
		},
	}

	cellSize := 48
	textSize := 24

	maxCols := 3
	colStep := s.InnerWidth / (maxCols)

	paddingY := 40
	marginY := s.InnerMarginY + paddingY

	percentageCount := float64(0)

	for i, v := range stats {
		row := i / 3
		col := i % maxCols

		x := s.InnerMarginX + col*colStep
		y := marginY + row*(cellSize+paddingY)

		err := DrawIcon(DrawIconOps{
			X:      x,
			Y:      y,
			Screen: s.Img,
			Size:   cellSize,
			Name:   v.Language.Name,
		})

		if err != nil {
			return err
		}

		s.DrawBar(percentageCount, v.Percentage, v.Language.Color)
		percentageCount += v.Percentage

		DrawText(DrawTextOps{
			X: x + cellSize,
			Y: y + textSize,

			Screen: s.Img,
			Size:   textSize,
			Text:   v.Language.Name,
			Color:  v.Language.Color,
		})
	}

	return nil
}

func (s *StatsBar) DrawBar(a, b float64, color color.Color) {
	startPoint := s.InnerMarginX + int((float64(s.InnerWidth)*a)/100)
	endPoint := int((float64(s.InnerWidth) * b) / 100)
	height := 10

	fmt.Println(a, b, startPoint, endPoint, int((float64(s.InnerWidth)*b)/100), (float64(s.InnerWidth)*b)/100)

	for x := range endPoint {
		for y := range height {
			s.Img.Set(startPoint+x, s.InnerMarginY+y, color)
		}
	}
}
