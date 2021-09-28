package helloworld

import "fmt"

const (
	spanish            = "spanish"
	france             = "france"
	englishHelloPrefix = "hello, "
	spanishHelloPrefix = "hola, "
	franceHelloPrefix  = "bonjour, "
)

func main() {
	fmt.Println(Hello("world", ""))
}

func Hello(name, language string) string {
	if name == "" {
		return englishHelloPrefix + "world"
	}
	prefix := getGreetingPrefix(language)
	return prefix + name
}

func getGreetingPrefix(language string) string {
	var prefix string

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case france:
		prefix = franceHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return prefix
}
