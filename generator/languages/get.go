package languages

func Get(name string) Language {
	if (languageMap[name] != Language{}) {
		return languageMap[name]
	}

	return defaultLanguage
}
