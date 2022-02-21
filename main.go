package main

import (
	"fmt"
	"github.com/aicam/PersianSpellCheck/pkg"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	s := "بر"
	fmt.Println(Reverse(string(pkg.Correct([]rune(s), 0, nil))))
}
