package main

import (
	"fmt"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}


func main() {
	//fmt.Println(Reverse("برای یک نفر تمامیت میا")[0])
	s := rune('ا')
	//sb := []rune{1576, 32, 1576, 1575}
	fmt.Println(s)
	//fmt.Println(string(sb))
}
