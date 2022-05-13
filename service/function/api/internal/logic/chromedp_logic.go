package logic

import (
	"context"
	"fmt"
	"log"
	"time"

	"wzqb/service/function/api/internal/svc"
	"wzqb/service/function/api/internal/types"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/zeromicro/go-zero/core/logx"
)

type ChromedpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChromedpLogic(ctx context.Context, svcCtx *svc.ServiceContext) ChromedpLogic {
	return ChromedpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Chromedp 下载
func (l *ChromedpLogic) Chromedp(req types.ChromedpParam) error {
	var (
		url     = req.URL
		timenow = fmt.Sprintf("%d", time.Now().Unix())
		pdfPath = "docs/report/" + timenow + ".pdf"
	)
	err := Wk_HTML_PDF(url, pdfPath)
	return err
}

func Wk_HTML_PDF(url string, path string) (err error) {
	// Create new PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}
	// 设置
	pdfg.Orientation.Set("Landscape") // 纵向
	// pdfg.MarginBottom.Set(0)
	// pdfg.MarginTop.Set(0)
	// pdfg.MarginLeft.Set(0)
	// pdfg.MarginRight.Set(0)

	pdfg.AddPage(wkhtmltopdf.NewPage(url))

	// PDFط³
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}
	// Write buffer contents to file on disk
	err = pdfg.WriteFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return
}
