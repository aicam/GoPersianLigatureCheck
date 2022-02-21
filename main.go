package main

import (
	"fmt"
	"github.com/aicam/PersianSpellCheck/pkg"
)

func main() {
	s1 := "سلام به همگیه دوستان"
	s2 := "من در تمامی ادوار کون میدهم"
	s3 := "کس ننه چوچاغ"
	fmt.Println(pkg.Correct(s1))
	fmt.Println(pkg.Correct(s2))
	fmt.Println(pkg.Correct(s3))
}
