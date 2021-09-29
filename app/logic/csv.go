package logic

import (
	"Moonlight/utils"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func OperateCSV(path, newPath string) {
	var (
		header = []string{"a", "b", "c", "d", "e", "f", "g"}
	)
	// writer
	f, _ := os.OpenFile(newPath, os.O_WRONLY|os.O_APPEND, 0666)
	writer := csv.NewWriter(f)
	writer.Write(header)
	writer.Flush()

	// reader
	clientsFile, _ := os.Open(path)
	reader := csv.NewReader(clientsFile)
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		line[6] = fmt.Sprint(utils.StrToInt(line[6]) + 1)
		writer.Write(line)
		writer.Flush()
	}
}
