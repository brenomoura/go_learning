package main

import (
	"fmt"
)

const portuguese = "Portuguese"
const french = "French"
const helloPrefixEnglish = "Hello, "
const helloPrefixPortuguese = "Olá, "
const helloPrefixFrench = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name + "!"
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = helloPrefixFrench
	case portuguese:
		prefix = helloPrefixPortuguese
	default:
		prefix = helloPrefixEnglish
	}

	return
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
