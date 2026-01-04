package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// banner ("Go", 6)
// returns
// 	 Go
// ------

func main() {
	banner("Go", 6)
	banner("Go🥤", 6)
	println(len("world🌎"))

}

func banner(text string, with int) {
	padding := (with - utf8.RuneCountInString((text))) / 2

	fmt.Println(" ")
	fmt.Print(strings.Repeat(" ", padding))
	fmt.Println(text)
	fmt.Println(strings.Repeat("_", with))
}
