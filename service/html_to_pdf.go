package service

import (
	"fmt"
	"time"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func HTMLtoPDF(htmlPath string) (pdfPath string, err error) {
	var (
		prefixPath = "docs/" + fmt.Sprintf("%d", time.Now().Unix())
	)
	pdfPath = prefixPath + "report.pdf"
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return
	}

	pdfg.PageSize.Set("A4")           // pdf纸张大小
	pdfg.Orientation.Set("Landscape") // pdf 纸张方向，横竖
	pdfg.AddPage(wkhtmltopdf.NewPage(htmlPath))
	err = pdfg.Create()
	if err != nil {
		return
	}
	err = pdfg.WriteFile(pdfPath)
	if err != nil {
		return
	}
	return
}
