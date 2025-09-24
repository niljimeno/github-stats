package stats

import "github.com/niljimeno/github-stats/generator/languages"

type Stat struct {
	Language   languages.Language
	Size       int
	Percentage float64
}
