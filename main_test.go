package main

import (
	"testing"

	"github.com/niljimeno/github-stats/generator/stats"
)

func TestHelloName(t *testing.T) {
	s, err := stats.GetStats("hi")
	if err != nil {
		t.Fatal(err)
	}

	t.Errorf("%s", s)
}
