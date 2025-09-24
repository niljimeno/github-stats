package languages

import "image/color"

var defaultLanguage = Language{
	Name: "go",
	Color: color.RGBA{
		R: 0x40,
		G: 0xc0,
		B: 0xff,
		A: 0xff,
	},
}

var languageMap = make(map[string]Language)

func SetUp() {
	languageMap["go"] = Language{
		Name: "go",
		Color: color.RGBA{
			R: 0x40,
			G: 0xc0,
			B: 0xff,
			A: 0xff,
		},
	}

	languageMap["haskell"] = Language{
		Name: "haskell",
		Color: color.RGBA{
			R: 0xbb,
			G: 0x77,
			B: 0xff,
			A: 0xff,
		},
	}

	languageMap["erlang"] = Language{
		Name: "erlang",
		Color: color.RGBA{
			R: 0xdd,
			G: 0x20,
			B: 0x44,
			A: 0xff,
		},
	}

	languageMap["gleam"] = Language{
		Name: "gleam",
		Color: color.RGBA{
			R: 0xff,
			G: 0xaf,
			B: 0xf3,
			A: 0xff,
		},
	}
}
