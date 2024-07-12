package main

import "fmt"

const (
	indonesia = "Indonesia"
	china = "China"
	
	englishHelloPrefix = "Hello, "
	indonesiaHelloPrefix = "Hi, "
	chinaHelloPrefix = "Hayya, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case indonesia:
		prefix = indonesiaHelloPrefix
	case china:
		prefix= chinaHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("namous", "Indonesia"))
}

