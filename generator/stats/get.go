package stats

import (
	"io"
	"net/http"

	"github.com/niljimeno/github-stats/generator/languages"
)

type repository struct {
	Stat     int
	Language string
}

func GetStats(username string) ([]byte, error) {
	res, err := http.Get("https://api.github.com/users/niljimeno/repos")
	if err != nil {
		return nil, err
	}

	content, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func IsLanguageValid(name string) bool {
	for _, v := range languages.LanguageMap {
		if v.Name == name {
			return true
		}
	}

	return false
}
