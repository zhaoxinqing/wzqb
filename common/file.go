package common

import (
	"archive/zip"
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

var (
	newFile *os.File
	err     error
)

// 创建空文件
func CreateEmptyFile() {
	newFile, err = os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	log.Println(newFile)
	newFile.Close()
}

// 截短文件
func TruncateFile() {
	// 裁剪一个文件到100个字节。
	// 如果文件本来就少于100个字节，则文件中原始内容得以保留，剩余的字节以null字节填充。
	// 如果文件本来超过100个字节，则超过的字节会被抛弃。
	// 这样我们总是得到精确的100个字节的文件。
	// 传入0则会清空文件。
	err := os.Truncate("test.txt", 100)
	if err != nil {
		fmt.Println(err)
	}
}

// ExtractFileType 从文件路径中提取文件类型
func ExtractFileType(filePath string) (fileType string) {
	var (
		dotChar  = "."
		dotIndex = strings.LastIndex(filePath, dotChar)
	)
	if dotIndex >= 0 {
		fileType = filePath[dotIndex+1:]
	} else {
		fileType = ""
	}
	if len(fileType) > 0 {
		fileType = strings.ToLower(fileType)
	}
	return
}

// 根据文件后缀获取文件格式
func GetFileType(a string) (result string) {
	result = path.Ext(a)
	return
}

var (
	fileInfo os.FileInfo
)

func OsOperation() {
	originalPath := "test.txt"
	newPath := "test2.txt"
	_ = os.Remove("test.txt")
	_ = os.Rename(originalPath, newPath)
	// 如果文件不存在，则返回错误
	fileInfo, err = os.Stat("test.txt")
	if err != nil {
		fmt.Println(err)
	}
}

// 打开和关闭文件
func OpenAndCloseFile() {
	// 简单地以只读的方式打开。下面的例子会介绍读写的例子。
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	file.Close()
	// OpenFile提供更多的选项。
	// 最后一个参数是权限模式permission mode
	// 第二个是打开时的属性
	file, err = os.OpenFile("test.txt", os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	file.Close()
	// 下面的属性可以单独使用，也可以组合使用。
	// 组合使用时可以使用 OR 操作设置 OpenFile的第二个参数，例如：
	// os.O_CREATE|os.O_APPEND
	// 或者 os.O_CREATE|os.O_TRUNC|os.O_WRONLY
	// os.O_RDONLY // 只读
	// os.O_WRONLY // 只写
	// os.O_RDWR // 读写
	// os.O_APPEND // 往文件中添建（Append）
	// os.O_CREATE // 如果文件不存在则先创建
	// os.O_TRUNC // 文件打开时裁剪文件
	// os.O_EXCL // 和O_CREATE一起使用，文件不能存在
	// os.O_SYNC // 以同步I/O的方式打开
}

// 检查文件是否存在

func CheckFileStat() {
	// 文件不存在则返回error
	fileInfo, err := os.Stat("test.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exist.")
		}
	}
	log.Println("File does exist. File information:")
	log.Println(fileInfo)
}

// 查看读写权限
func GetFilePermission() {
	// 这个例子测试写权限，如果没有写权限则返回error。
	// 注意文件不存在也会返回error，需要检查error的信息来获取到底是哪个错误导致。
	file, err := os.OpenFile("test.txt", os.O_WRONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			log.Println("Error: Write permission denied.")
		}
	}
	file.Close()
	// 测试读权限
	file, err = os.OpenFile("test.txt", os.O_RDONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			log.Println("Error: Read permission denied.")
		}
	}
	file.Close()
}

// 改变权限、拥有者、时间戳
func ChangeFilePermisssion() {
	// 使用Linux风格改变文件权限
	err := os.Chmod("test.txt", 0777)
	if err != nil {
		log.Println(err)
	}
	// 改变文件所有者
	err = os.Chown("test.txt", os.Getuid(), os.Getgid())
	if err != nil {
		log.Println(err)
	}
	// 改变时间戳
	twoDaysFromNow := time.Now().Add(48 * time.Hour)
	lastAccessTime := twoDaysFromNow
	lastModifyTime := twoDaysFromNow
	err = os.Chtimes("test.txt", lastAccessTime, lastModifyTime)
	if err != nil {
		log.Println(err)
	}
}

// 读写、复制文件
func ReadAndCopyFile() {
	// 打开原始文件
	originalFile, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer originalFile.Close()
	// 创建新的文件作为目标文件
	newFile, err := os.Create("test_copy.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer newFile.Close()
	// 从源中复制字节到目标文件
	bytesWritten, err := io.Copy(newFile, originalFile)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("Copied %d bytes.", bytesWritten)
	// 将文件内容flush到硬盘中
	err = newFile.Sync()
	if err != nil {
		fmt.Println(err)
	}
}

// 写文件
// 可以使用os包写入一个打开的文件。
// 因为Go可执行包是静态链接的可执行文件，你import的每一个包都会增加你的可执行文件的大小。
// 其它的包如io、｀ioutil｀、｀bufio｀提供了一些方法，但是它们不是必须的。
func WriteFile() {
	// 可写方式打开文件
	file, err := os.OpenFile(
		"test.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	// 写字节到文件中
	byteSlice := []byte("Bytes!\n")
	bytesWritten, err := file.Write(byteSlice)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("Wrote %d bytes.\n", bytesWritten)
}

// 快写文件
// ioutil包有一个非常有用的方法WriteFile()可以处理创建／打开文件、写字节slice和关闭文件一系列的操作。
// 如果你需要简洁快速地写字节slice到文件中，你可以使用它。
func QuickWriteFile() {
	err := ioutil.WriteFile("test.txt", []byte("Hi\n"), 0666)
	if err != nil {
		fmt.Println(err)
	}
}

// 使用缓存写
// bufio包提供了带缓存功能的writer，所以你可以在写字节到硬盘前使用内存缓存。
// 当你处理很多的数据很有用，因为它可以节省操作硬盘I/O的时间。
// 在其它一些情况下它也很有用，比如你每次写一个字节，把它们攒在内存缓存中，然后一次写入到硬盘中，减少硬盘的磨损以及提升性能。
func BufferWriteFile() {
	// 打开文件，只写
	file, err := os.OpenFile("test.txt", os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	// 为这个文件创建buffered writer
	bufferedWriter := bufio.NewWriter(file)
	// 写字节到buffer
	bytesWritten, err := bufferedWriter.Write(
		[]byte{65, 66, 67},
	)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("Bytes written: %d\n", bytesWritten)
	// 写字符串到buffer
	// 也可以使用 WriteRune() 和 WriteByte()
	bytesWritten, err = bufferedWriter.WriteString(
		"Buffered string\n",
	)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("Bytes written: %d\n", bytesWritten)
	// 检查缓存中的字节数
	unflushedBufferSize := bufferedWriter.Buffered()
	log.Printf("Bytes buffered: %d\n", unflushedBufferSize)
	// 还有多少字节可用（未使用的缓存大小）
	bytesAvailable := bufferedWriter.Available()
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("Available buffer: %d\n", bytesAvailable)
	// 写内存buffer到硬盘
	bufferedWriter.Flush()
	// 丢弃还没有flush的缓存的内容，清除错误并把它的输出传给参数中的writer
	// 当你想将缓存传给另外一个writer时有用
	bufferedWriter.Reset(bufferedWriter)
	bytesAvailable = bufferedWriter.Available()
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("Available buffer: %d\n", bytesAvailable)
	// 重新设置缓存的大小。
	// 第一个参数是缓存应该输出到哪里，这个例子中我们使用相同的writer。
	// 如果我们设置的新的大小小于第一个参数writer的缓存大小， 比如10，我们不会得到一个10字节大小的缓存，
	// 而是writer的原始大小的缓存，默认是4096。
	// 它的功能主要还是为了扩容。
	bufferedWriter = bufio.NewWriterSize(
		bufferedWriter,
		8000,
	)
	// resize后检查缓存的大小
	bytesAvailable = bufferedWriter.Available()
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("Available buffer: %d\n", bytesAvailable)
}

// 读取最多N个字节
// os.File提供了文件操作的基本功能， 而io、ioutil、bufio提供了额外的辅助函数。
func LimitReadFile() {
	// 打开文件，只读
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	// 从文件中读取len(b)字节的文件。
	// 返回0字节意味着读取到文件尾了
	// 读取到文件会返回io.EOF的error
	byteSlice := make([]byte, 16)
	bytesRead, err := file.Read(byteSlice)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("Number of bytes read: %d\n", bytesRead)
	log.Printf("Data read: %s\n", byteSlice)
}

// 读取正好N个字节
func ControllReadFile() {
	// Open file for reading
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	// file.Read()可以读取一个小文件到大的byte slice中，
	// 但是io.ReadFull()在文件的字节数小于byte slice字节数的时候会返回错误
	byteSlice := make([]byte, 2)
	numBytesRead, err := io.ReadFull(file, byteSlice)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("Number of bytes read: %d\n", numBytesRead)
	log.Printf("Data read: %s\n", byteSlice)
}

// 读取至少N个字节
func LessReadFile() {
	// 打开文件，只读
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	byteSlice := make([]byte, 512)
	minBytes := 8
	// io.ReadAtLeast()在不能得到最小的字节的时候会返回错误，但会把已读的文件保留
	numBytesRead, err := io.ReadAtLeast(file, byteSlice, minBytes)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("Number of bytes read: %d\n", numBytesRead)
	log.Printf("Data read: %s\n", byteSlice)
}

// 读取全部字节
func AllRead() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	// os.File.Read(), io.ReadFull() 和
	// io.ReadAtLeast() 在读取之前都需要一个固定大小的byte slice。
	// 但ioutil.ReadAll()会读取reader(这个例子中是file)的每一个字节，然后把字节slice返回。
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Data as hex: %x\n", data)
	fmt.Printf("Data as string: %s\n", data)
	fmt.Println("Number of bytes read:", len(data))
}

// 快读到内存
func QuickReadFile() {
	// 读取文件到byte slice中
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("Data read: %s\n", data)
}

// 使用缓存读
// 有缓存写也有缓存读。
// 缓存reader会把一些内容缓存在内存中。
// 它会提供比os.File和io.Reader更多的函数,缺省的缓存大小是4096，最小缓存是16。
func BufferReadFile() {
	// 打开文件，创建buffered reader
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	bufferedReader := bufio.NewReader(file)
	// 得到字节，当前指针不变
	byteSlice := make([]byte, 5)
	byteSlice, err = bufferedReader.Peek(5)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Peeked at 5 bytes: %s\n", byteSlice)
	// 读取，指针同时移动
	numBytesRead, err := bufferedReader.Read(byteSlice)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Read %d bytes: %s\n", numBytesRead, byteSlice)
	// 读取一个字节, 如果读取不成功会返回Error
	myByte, err := bufferedReader.ReadByte()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Read 1 byte: %c\n", myByte)
	// 读取到分隔符，包含分隔符，返回byte slice
	dataBytes, err := bufferedReader.ReadBytes('\n')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Read bytes: %s\n", dataBytes)
	// 读取到分隔符，包含分隔符，返回字符串
	dataString, err := bufferedReader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Read string: %s\n", dataString)
	//这个例子读取了很多行，所以test.txt应该包含多行文本才不至于出错
}

// 使用 scanner
// Scanner是bufio包下的类型,在处理文件中以分隔符分隔的文本时很有用。
// 通常我们使用换行符作为分隔符将文件内容分成多行。在CSV文件中，逗号一般作为分隔符。
// os.File文件可以被包装成bufio.Scanner，它就像一个缓存reader。
// 我们会调用Scan()方法去读取下一个分隔符，使用Text()或者Bytes()获取读取的数据。

// 分隔符可以不是一个简单的字节或者字符，有一个特殊的方法可以实现分隔符的功能，以及将指针移动多少，返回什么数据。
// 如果没有定制的SplitFunc提供，缺省的ScanLines会使用newline字符作为分隔符，其它的分隔函数还包括ScanRunes和ScanWords,皆在bufio包中。

// 为一个文件创建了bufio.Scanner，并按照单词逐个读取：
func ScannerFile() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	// 缺省的分隔函数是bufio.ScanLines,我们这里使用ScanWords。
	// 也可以定制一个SplitFunc类型的分隔函数
	scanner.Split(bufio.ScanWords)
	// scan下一个token.
	success := scanner.Scan()
	if success == false {
		// 出现错误或者EOF是返回Error
		err = scanner.Err()
		if err == nil {
			log.Println("Scan completed and reached EOF")
		} else {
			fmt.Println(err)
		}
	}
	// 得到数据，Bytes() 或者 Text()
	fmt.Println("First word found:", scanner.Text())
	// 再次调用scanner.Scan()发现下一个token
}

// 压缩 打包（zip）文件
func ZipFile() {
	// 创建一个打包文件
	outFile, err := os.Create("test.zip")
	if err != nil {
		fmt.Println(err)
	}
	defer outFile.Close()
	// 创建zip writer
	zipWriter := zip.NewWriter(outFile)
	// 往打包文件中写文件。
	// 这里我们使用硬编码的内容，你可以遍历一个文件夹，把文件夹下的文件以及它们的内容写入到这个打包文件中。
	var filesToArchive = []struct {
		Name, Body string
	}{
		{"test.txt", "String contents of file"},
		{"test2.txt", "\x61\x62\x63\n"},
	}
	// 下面将要打包的内容写入到打包文件中，依次写入。
	for _, file := range filesToArchive {
		fileWriter, err := zipWriter.Create(file.Name)
		if err != nil {
			fmt.Println(err)
		}
		_, err = fileWriter.Write([]byte(file.Body))
		if err != nil {
			fmt.Println(err)
		}
	}
	// 清理
	err = zipWriter.Close()
	if err != nil {
		fmt.Println(err)
	}
}

// 抽取(unzip) 文件
func UnzipFile() {
	zipReader, err := zip.OpenReader("test.zip")
	if err != nil {
		fmt.Println(err)
	}
	defer zipReader.Close()
	// 遍历打包文件中的每一文件/文件夹
	for _, file := range zipReader.Reader.File {
		// 打包文件中的文件就像普通的一个文件对象一样
		zippedFile, err := file.Open()
		if err != nil {
			fmt.Println(err)
		}
		defer zippedFile.Close()
		// 指定抽取的文件名。
		// 你可以指定全路径名或者一个前缀，这样可以把它们放在不同的文件夹中。
		// 我们这个例子使用打包文件中相同的文件名。
		targetDir := "./"
		extractedFilePath := filepath.Join(
			targetDir,
			file.Name,
		)
		// 抽取项目或者创建文件夹
		if file.FileInfo().IsDir() {
			// 创建文件夹并设置同样的权限
			log.Println("Creating directory:", extractedFilePath)
			os.MkdirAll(extractedFilePath, file.Mode())
		} else {
			//抽取正常的文件
			log.Println("Extracting file:", file.Name)
			outputFile, err := os.OpenFile(
				extractedFilePath,
				os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
				file.Mode(),
			)
			if err != nil {
				fmt.Println(err)
			}
			defer outputFile.Close()
			// 通过io.Copy简洁地复制文件内容
			_, err = io.Copy(outputFile, zippedFile)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

// 压缩文件
func RarFile() {
	outputFile, err := os.Create("test.txt.gz")
	if err != nil {
		fmt.Println(err)
	}
	gzipWriter := gzip.NewWriter(outputFile)
	defer gzipWriter.Close()
	// 当我们写如到gizp writer数据时，它会依次压缩数据并写入到底层的文件中。
	// 我们不必关心它是如何压缩的，还是像普通的writer一样操作即可。
	_, err = gzipWriter.Write([]byte("Gophers rule!\n"))
	if err != nil {
		fmt.Println(err)
	}
	log.Println("Compressed data written to file.")
}

// 解压缩文件
func UnRarFile() {
	// 打开一个gzip文件。
	// 文件是一个reader,但是我们可以使用各种数据源，比如web服务器返回的gzipped内容，
	// 它的内容不是一个文件，而是一个内存流
	gzipFile, err := os.Open("test.txt.gz")
	if err != nil {
		fmt.Println(err)
	}
	gzipReader, err := gzip.NewReader(gzipFile)
	if err != nil {
		fmt.Println(err)
	}
	defer gzipReader.Close()
	// 解压缩到一个writer,它是一个file writer
	outfileWriter, err := os.Create("unzipped.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer outfileWriter.Close()
	// 复制内容
	_, err = io.Copy(outfileWriter, gzipReader)
	if err != nil {
		fmt.Println(err)
	}
}

// 其它
// 临时文件和目录
// ioutil提供了两个函数: TempDir() 和 TempFile()。
// 使用完毕后，调用者负责删除这些临时文件和文件夹。
// 有一点好处就是当你传递一个空字符串作为文件夹名的时候，它会在操作系统的临时文件夹中创建这些项目（/tmp on Linux）。
// os.TempDir()返回当前操作系统的临时文件夹。
func OtherFileOperation() {
	// 在系统临时文件夹中创建一个临时文件夹
	tempDirPath, err := ioutil.TempDir("", "myTempDir")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Temp dir created:", tempDirPath)
	// 在临时文件夹中创建临时文件
	tempFile, err := ioutil.TempFile(tempDirPath, "myTempFile.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Temp file created:", tempFile.Name())
	// ... 做一些操作 ...
	// 关闭文件
	err = tempFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	// 删除我们创建的资源
	err = os.Remove(tempFile.Name())
	if err != nil {
		fmt.Println(err)
	}
	err = os.Remove(tempDirPath)
	if err != nil {
		fmt.Println(err)
	}
}

// 通过HTTP下载文件
func HTTPDownloadFile(url string) (filePath string, err error) {
	filePath = "/doc/temp/demo.html"
	newFile, err := os.Create(filePath)
	if err != nil {
		return
	}
	defer newFile.Close()
	response, err := http.Get(url)
	if err != nil {
		return
	}
	defer response.Body.Close()
	numBytesWritten, err := io.Copy(newFile, response.Body)
	if err != nil {
		return
	}
	fmt.Println(numBytesWritten)
	return
}
