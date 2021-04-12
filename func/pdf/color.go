package main

// import (
// 	"io/ioutil"
// 	"log"
// 	"os"

// 	"github.com/signintech/gopdf"
// )

// func main() {
// 	pdf := gopdf.GoPdf{}
// 	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4
// 	pdf.AddPage()
// 	err := pdf.AddTTFFont("HDZB_5", "./func/pdf/tty/wts11.ttf")
// 	if err != nil {
// 		log.Print(err.Error())
// 		return
// 	}
// 	err = pdf.SetFont("HDZB_5", "", 14)
// 	if err != nil {
// 		log.Print(err.Error())
// 		return
// 	}
// 	// 添加文字 带颜色
// 	pdf.SetGrayFill(0.5)
// 	pdf.Cell(nil, "gray")

// 	pdf.SetTextColor(255, 0, 0)
// 	pdf.Br(40)
// 	pdf.Cell(nil, "red")

// 	pdf.SetTextColor(0, 0, 0)
// 	pdf.Br(20)
// 	pdf.Cell(nil, "gray")

// 	pdf.SetGrayFill(0)
// 	pdf.Br(40)
// 	pdf.Cell(nil, "black")

// 	pdf.SetTextColor(0, 0, 255)
// 	pdf.Br(20)
// 	pdf.Cell(nil, "blue")

// 	// 添加文字
// 	pdf.SetGrayFill(0.5)
// 	pdf.Cell(nil, "您好")

// 	pdf.SetGrayFill(0.5)
// 	pdf.Cell(nil, "您好")

// 	//添加图片
// 	//use path
// 	pdf.Image("./doc/001.png", 200, 50, nil)

// 	//use image holder by []byte
// 	imgH1, err := gopdf.ImageHolderByBytes(getImageBytes())
// 	if err != nil {
// 		log.Print(err.Error())
// 		return
// 	}
// 	pdf.ImageByHolder(imgH1, 200, 250, nil)

// 	//use image holder by io.Reader
// 	file, err := os.Open("./doc/001.png")
// 	if err != nil {
// 		log.Print(err.Error())
// 		return
// 	}
// 	imgH2, err := gopdf.ImageHolderByReader(file)
// 	if err != nil {
// 		log.Print(err.Error())
// 		return
// 	}
// 	pdf.ImageByHolder(imgH2, 200, 450, nil)

// 	pdf.SetX(250)
// 	pdf.SetY(200)
// 	pdf.Cell(nil, "gopher and gopher")

// 	pdf.SetLineWidth(2)
// 	pdf.SetLineType("dashed")
// 	//pdf.SetLineType("dotted")
// 	pdf.Line(10, 100, 585, 100)

// 	pdf.WritePdf("./doc/color3.pdf")
// }

// func getImageBytes() []byte {
// 	b, err := ioutil.ReadFile("./doc/001.png")
// 	if err != nil {
// 		panic(err)
// 	}
// 	return b
// }
