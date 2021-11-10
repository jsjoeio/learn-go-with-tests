package main

const spanish = "Spanish"
const french = "French"
const italian = "Italian"
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const englishHelloPrefix = "Hello, "
const italianHelloPrefix = "Ciao, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

// Notice in this function we don't define prefix
// in the function body
// but we do create a _named return value_ with `(prefix string)`
// this creates a variable called prefix inside our function
// so cool!!
//
// Another thing to note is private functions start with a lowercase letter
// while public functions start with a capital one.
func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case italian:
		prefix = italianHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
