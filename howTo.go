package main

import (
	"github.com/aicam/GoPersianLigatureCheck/pkg"
	"log"
)

func main() {

	s1 := "میز ۰۰۳"
	log.Println(pkg.Correct(s1))
	//s1 := "امتناع: خود داری - پرهیز - خودداری - سر باز زدن- است - ام - اس"
	//s2 := "سبزی پلو قالبی با تهدیگ سیب زمینی "
	//fmt.Println(s1)
	//fmt.Println(pkg.Correct(s1))
	//fmt.Println("----------")
	//fmt.Println(pkg.Correct(s2))

}
