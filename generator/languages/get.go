package languages

func GetLanguage(name string) Language {
	if (LanguageMap[name] != Language{}) {
		return LanguageMap[name]
	}

	return defaultLanguage
}
