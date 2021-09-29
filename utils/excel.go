package utils

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"golang.org/x/text/encoding/simplifiedchinese"
)

// ReadFileToRecords ...
func ReadFileToRecords(fileType, savePath string, limit int) (records [][]string, err error) {
	switch {
	case strings.EqualFold(fileType, "csv"):
		var (
			recordFile   *os.File
			firstRowRead = false
		)
		recordFile, err = os.Open(savePath)
		if err != nil {
			return
		}
		defer recordFile.Close()
		gbkDecoder := simplifiedchinese.GBK.NewDecoder()
		reader := csv.NewReader(recordFile)
		for {
			record, readErr := reader.Read()
			if readErr == io.EOF {
				break
			}
			if readErr != nil {
				err = readErr
				return
			}
			encodeRecord := make([]string, 0, len(record))
			for _, value := range record {
				var encodeValue string
				coding := GetStringCoding(value)
				switch coding {
				case "UTF8":
					encodeValue = value
				case GBK:
					encodingString, _ := gbkDecoder.Bytes([]byte(value))
					encodeValue = string(encodingString)
				default:
					if encodingString, decodeError := gbkDecoder.Bytes([]byte(value)); decodeError == nil {
						encodeValue = string(encodingString)
					} else {
						err = errors.New("file encoding is not supported")
						return
					}
				}
				if !firstRowRead {
					encodeValue = strings.TrimSpace(encodeValue)
					// 移除特殊字符，比如 &#65279;
					encodeValue = strings.ReplaceAll(encodeValue, "\uFEFF", "")
				}
				encodeValue = strings.TrimSpace(encodeValue)
				encodeRecord = append(encodeRecord, encodeValue)
			}
			records = append(records, encodeRecord)
			firstRowRead = true
			if limit > 0 && len(records) > limit {
				break
			}
		}
	case strings.EqualFold(fileType, "xlsx"):
		recordFile, err := excelize.OpenFile(savePath)
		if err != nil {
			return records, err
		}
		sheetName := recordFile.GetSheetName(0)
		if len(sheetName) > 0 {
			rows, _ := recordFile.Rows(sheetName)
			var (
				rowIndex = 1
				// format         string
				// found          bool
				// styleNumFmtMap = make(map[int]string)
			)
			for rows != nil && rows.Next() {
				record, _ := rows.Columns()
				for i := range record {
					// cellName, _ := excelize.CoordinatesToCellName(i+1, rowIndex)
					// styleID, _ := recordFile.GetCellStyle(sheetName, cellName)
					// if styleID > 0 {
					// 	if format, found = styleNumFmtMap[styleID]; !found {
					// 		format = getFormatByStyleID(styleID, recordFile)
					// 		if len(format) > 0 {
					// 			styleNumFmtMap[styleID] = format
					// 		}
					// 	}
					// 	if len(format) > 0 {
					// 		record[i] = formatCellValue(record[i], format)
					// 	}
					// }
					record[i] = strings.TrimSpace(record[i])
				}
				records = append(records, record)
				if limit > 0 && len(records) > limit {
					break
				}
				rowIndex += 1
			}
		}
	default:
		err = errors.New("file types not supported")
		return
	}
	return
}

// func getFormatByStyleID(styleID int, file *excelize.File) (format string) {
// 	numFmtID := *file.Styles.CellXfs.Xf[styleID].NumFmtID
// 	if numFmtID == 14 {
// 		format = "mm-dd-yy"
// 		return
// 	}
// 	if file.Styles == nil || file.Styles.NumFmts == nil {
// 		return
// 	}
// 	for _, xlsxFmt := range file.Styles.NumFmts.NumFmt {
// 		if xlsxFmt.NumFmtID == numFmtID {
// 			format = strings.ToLower(xlsxFmt.FormatCode)
// 			return
// 		}
// 	}
// 	return
// }

// var excelEpoch = time.Date(1899, time.December, 30, 0, 0, 0, 0, time.UTC)

// func formatCellValue(value string, format string) string {
// 	if len(format) == 0 {
// 		return value
// 	}
// 	if isDateFormat(format) {
// 		// 根据格式转换日期内容
// 		if days, err := strconv.Atoi(value); err == nil {
// 			// 数字转日期
// 			theDate := excelEpoch.Add(time.Hour * time.Duration(days*24))
// 			value = theDate.Format(`2006-01-02`)
// 		} else {
// 			// mm-dd-yy 格式
// 			var formatLayout = strings.ReplaceAll(format, "mm", "01")
// 			formatLayout = strings.ReplaceAll(formatLayout, "dd", "02")
// 			if strings.Contains(formatLayout, "yyyy") {
// 				formatLayout = strings.ReplaceAll(formatLayout, "yyyy", "2006")
// 			} else {
// 				formatLayout = strings.ReplaceAll(formatLayout, "yy", "06")
// 			}
// 			if recordTime, err := time.Parse(formatLayout, value); err == nil {
// 				value = recordTime.Format(`2006-01-02`)
// 			}
// 		}
// 	}
// 	return value
// }
// func isDateFormat(format string) bool {
// 	if strings.Contains(format, "yy") || strings.Contains(format, "mm") ||
// 		strings.Contains(format, "dd") {
// 		return true
// 	}
// 	return false
// }

func ReadExcel(filePath string) (result interface{}, err error) {
	clientsFile, err := os.Open(filePath)
	reader := csv.NewReader(clientsFile)
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if line[0] != "name" && line[1] != "address" { // 过滤标题行
			err = errors.New("这货不正宗")
		}
	}
	return nil, err
}

// GetStringCoding 获取字符串编码
// 需要说明的是，isGBK()是通过双字节是否落在gbk的编码范围内实现的，
// 而utf-8编码格式的每个字节都是落在gbk的编码范围内，
// 所以只有先调用isUtf8()先判断不是utf-8编码，再调用isGBK()才有意义
func GetStringCoding(dataStr string) string {
	// 过滤特殊字符
	dataStr = strings.ReplaceAll(dataStr, "·", "")
	data := []byte(dataStr)
	if IsUTF8(data) {
		return UTF8
	} else if IsGBK(data) {
		return GBK
	} else {
		return UNKNOWN
	}
}

const (
	// GBK ...
	GBK string = "GBK"
	// UTF8 ...
	UTF8 string = "UTF8"
	// UNKNOWN ...
	UNKNOWN string = "UNKNOWN"
)

// IsGBK 判断是否是GBK编码
func IsGBK(data []byte) bool {
	length := len(data)
	var i = 0
	for i < length-1 {
		if data[i] <= 0x7f {
			//编码0~127,只有一个字节的编码，兼容ASCII码
			i++
			continue
		} else {
			//大于127的使用双字节编码，落在gbk编码范围内的字符
			if data[i] >= 0x81 &&
				data[i] <= 0xfe &&
				data[i+1] >= 0x40 &&
				data[i+1] <= 0xfe &&
				data[i+1] != 0x7f {
				i += 2
				continue
			} else {
				return false
			}
		}
	}
	return true
}

// IsUTF8 判断是否是UTF8编码
func IsUTF8(data []byte) bool {
	i := 0
	for i < len(data) {
		if (data[i] & 0x80) == 0x00 {
			// 0XXX_XXXX
			i++
			continue
		} else if num := preNUm(data[i]); num > 2 {
			// 110X_XXXX 10XX_XXXX
			// 1110_XXXX 10XX_XXXX 10XX_XXXX
			// 1111_0XXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
			// 1111_10XX 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
			// 1111_110X 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
			// preNUm() 返回首个字节的8个bits中首个0bit前面1bit的个数，该数量也是该字符所使用的字节数
			i++
			for j := 0; j < num-1; j++ {
				//判断后面的 num - 1 个字节是不是都是10开头
				if (data[i] & 0xc0) != 0x80 {
					return false
				}
				i++
			}
		} else {
			//其他情况说明不是utf-8
			return false
		}
	}
	return true
}

func preNUm(data byte) int {
	var mask byte = 0x80
	var num int = 0
	//8bit中首个0bit前有多少个1bits
	for i := 0; i < 8; i++ {
		if (data & mask) == mask {
			num++
			mask = mask >> 1
		} else {
			break
		}
	}
	return num
}
