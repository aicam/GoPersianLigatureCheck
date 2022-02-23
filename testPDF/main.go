package main

import (
	persian_ligature "github.com/aicam/PersianSpellCheck/pkg"
	"github.com/jung-kurt/gofpdf"
	"log"
)

func main() {

	pdf := gofpdf.NewCustom(&gofpdf.InitType{
		OrientationStr: "p",
		UnitStr:        "mm",
		SizeStr:        "",
		Size: gofpdf.SizeType{
			Wd: 80,
			Ht: 80,
		},
		FontDirStr: "",
	})
	defer pdf.Close()
	pdf.SetAutoPageBreak(false, 0)
	pdf.AddPageFormat("p", gofpdf.SizeType{
		Wd: 80,
		Ht: 70,
	})
	pdf.SetY(0)
	pdf.SetX(0)
	pdf.AddUTF8Font("iransans", "B", "./testPDF/font/IRANSansWeb.ttf")
	pdf.SetFont("iransans", "B", 3)

	pdf.LTR()

	ggg := "یسشهخدیخشهس یتسخک یتکخستبتخکیبکتشثتبخش ه س"
	g2 := "یبعیهسادبسهع اهع یبساسیدعه ابس سذدیهب سیب <عسهید> باهسعس >>"
	g3 := "گککمو مگ گ یسشخت  (گش) یض ثث صثیثی پثپ ثپثپث پ ث"
	//sss := &ggg
	log.Println("Len gggg: ", len(ggg))
	pdf.CellFormat(80, 23, persian_ligature.Correct(ggg), "", 20, "C", false, 0, "")
	pdf.CellFormat(80, 23, persian_ligature.Correct(g2), "", 20, "C", false, 0, "")
	pdf.CellFormat(80, 23, persian_ligature.Correct(g3), "", 20, "C", false, 0, "")
	err := pdf.OutputFileAndClose("fff.pdf")
	if err != nil {
		log.Println(err)
	}
	//return
}
