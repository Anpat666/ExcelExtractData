package cores

import (
	"fmt"
	"os"

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
