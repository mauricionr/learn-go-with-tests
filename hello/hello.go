package main

import "fmt"

const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "


func Hello(name string, languate string) string {

		if languate == spanish {
			return spanishHelloPrefix + name
		}

		if name == "" {
			name = "World"
		}

    return englishHelloPrefix + name
}

func main() {
    fmt.Println(Hello("World", ""))
}
