package server

import (
	"Kilroy/app/common"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetWeb(c *gin.Context) {
	url := c.Query("url")
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	// 自定义Header
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36 Edg/92.0.902.78")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("http get error", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read error", err)
		return
	}
	path := "./doc/web/a.html"
	file, _ := os.Create(path)
	file.Write(body)
	file.Close()

	common.ResSuccess(c, string(body))
}

func writeIn(path string, info []byte) {
	// os.MkdirAll("", 0777)

	// 创建文件
	os.Create("./doc/web/file.html")

	// 打开文件,得到一个 *File 对象, 用于后续的写入
	file, _ := os.OpenFile("./doc/web/file.html", 2, 0666)

	// file.WriteString("add string")
	file.Write(info)
}
