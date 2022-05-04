# About Project
This project is developed to handle Persian characters ligature. By default, 
Golang does not support Persian characters and does not put correct form of alphabets
in strings.

# Features
- Correct alphabets combination
- Does not change Persian and English numbers order
- Compatible with PDF generator libraries (refer testPDF)

# How to use
There is a pkg module which can be renamed to anything easier for yourself. 
The only function needs to be called is "Correct". It will check string
and replace correct form of each alphabet. <br>
First pull the repository to your project <br>
`go pull https://github.com/aicam/GoPersianLigatureCheck.git` <br>
Next you can use library as below
```go
    import (
    PersianCorrection "github.com/aicam/GoPersianLigatureCheck/pkg"
    "log"
    )
    s1 := "میز ۰۰۳"
    log.Println(PersianCorrection.Correct(s1))
```
