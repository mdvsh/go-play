package hello

const constantPrefix = "hi "
const hindiPrefix = "namaste "
const frenchPrefix = "bonjour "

func Hello(name, lang string) string {
	if name == "" {
		name = "world"
	}
	return getPrefix(lang) + name
}

func getPrefix(lang string) (prefix string) {
	// named return value: 0 value of type
	switch lang {
	case "hindi":
		prefix = hindiPrefix
	case "french":
		prefix = frenchPrefix
	default:
		prefix = constantPrefix
	}
	return
}
