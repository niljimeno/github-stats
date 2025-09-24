package generator

import (
	"github.com/niljimeno/github-stats/generator/languages"
	"github.com/niljimeno/github-stats/generator/stats"
)

func (s *StatsBar) DrawInnerStats() error {
	stats := []stats.Stat{
		{
			Language:   languages.Get("go"),
			Percentage: 60,
		},
		{
			Language:   languages.Get("haskell"),
			Percentage: 30,
		},
		{
			Language: languages.Get("erlang"),
			Amount:   10,
		},
		{
			Language: languages.Get("gleam"),
			Amount:   10,
		},
	}

	cellSize := 48
	textSize := 24

	maxCols := 3
	colStep := s.InnerWidth / (maxCols)

	paddingY := 40
	marginY := s.InnerMarginY + paddingY

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
