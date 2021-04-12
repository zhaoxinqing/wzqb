package wkhtmltopdf

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
