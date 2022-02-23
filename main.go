package main

import (
	"github.com/aicam/PersianSpellCheck/pkg"
	"log"
)

func main() {

	s1 := "سلام بر ۱۲۳۴۵۶ تا کیرخر"
	log.Println(pkg.Correct(s1))
	//s1 := "امتناع: خود داری - پرهیز - خودداری - سر باز زدن- است - ام - اس"
	//s2 := "سبزی پلو قالبی با تهدیگ سیب زمینی "
	//fmt.Println(s1)
	//fmt.Println(pkg.Correct(s1))
	//fmt.Println("----------")
	//fmt.Println(pkg.Correct(s2))

}
