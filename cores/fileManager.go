package cores

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/xuri/excelize/v2"
)

// 打开Excel表格
func OpenExcel(tableAddress string) *excelize.File {
	f, errf := excelize.OpenFile(tableAddress)
	if errf != nil {
		fmt.Println(errf)
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	return f
}

// 创建文档//写入数据
func UpDataReport(data string, path string) {
	writefile, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("创建或打开TXT失败", err)
	}
	_, err2 := writefile.Write([]byte(data))
	if err2 != nil {
		fmt.Println("数据写入TXT失败", err2)
	}
	defer writefile.Close()
}

// 清空文档
func ClearDocument(path string) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("创建或打开TXT失败", err)
	}

	file.Truncate(0)
	defer file.Close()
}

// 给文档替换
func ReplaceDocument(old string, new string, TxtPath string, startLint int) {
	file, err := os.OpenFile(TxtPath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("打开修改TXT失败", err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	var newText string
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()

		if lineNumber >= startLint {
			newLine := strings.Replace(line, old, new, -1)
			newText += newLine + "\n"
		} else {
			newText += line + "\n"
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("读取文件失败:", err)
		return
	}

	if err := file.Truncate(0); err != nil {
		fmt.Println("清空文件内容失败:", err)
		return
	}

	if _, err := file.Seek(0, 0); err != nil {
		fmt.Println("移动文件指针失败:", err)
		return
	}

	if _, err := file.WriteString(newText); err != nil {
		fmt.Println("写入新内容失败:", err)
		return
	}

}
