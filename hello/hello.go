package main

import "fmt"

const helloPrefixEnglish = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "world"
	}
	return helloPrefixEnglish + name + "!"
}

func main() {
	fmt.Println(Hello("Chris"))
}
