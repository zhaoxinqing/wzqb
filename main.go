package main

// import (
// 	"Kilroy/htmltopdf"
// 	"fmt"
// 	"log"
// )

// func main() {
// 	var (
// 		pdf = htmltopdf.NewPdf()
// 		// str = "https://cn.bing.com/?scope=web&FORM=BEHPTB&ensearch=1"
// 		//str = "https://www.hao123.com/"
// 		//str = "http://invest.sheitc.sh.gov.cn/"
// 		str = "http://192.168.2.132:8082/5.html"
// 		loc = "./doc/008.pdf"
// 	)
// 	url, err := pdf.OutFile(str, loc)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	fmt.Println(url)
// }

import (
	"fmt"
	"log"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func main() {
	// Create new PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}
	url := "http://google.com/"

	pdfg.AddPage(wkhtmltopdf.NewPage(url))

	// PDFط³
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	// ³
	err = pdfg.WriteFile("./google.pdf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("tada!")
}
