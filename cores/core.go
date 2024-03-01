package cores

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

// 提取Excel的值
func GetExcelValue(f *excelize.File, tableName string, element string) string {
	cell, err := f.GetCellValue(tableName, element)
	if err != nil {
		fmt.Println("表格元素获取失败：", element, err)
	}
	return cell
}

// 提取行数据
func GetExcelRows(f *excelize.File, tableName string) [][]string {
	rows, err := f.GetRows(tableName)
	if err != nil {
		fmt.Println("表格获取列元素失败：", tableName, err)
	}
	return rows
}

// 提取列数据
func GetExcelCols(f *excelize.File, tableName string) [][]string {
	rows, err := f.GetCols(tableName)
	if err != nil {
		fmt.Println("表格获取列元素失败：", tableName, err)
	}
	return rows
}

// 转换判定 100以上格式化输出加万 100以下直接格式化输出加元
func TransitionData(data string) string {
	datafloat, _ := strconv.ParseFloat(data, 64)
	if datafloat < 100 && datafloat > -100 {
		strdata := strconv.FormatFloat(datafloat, 'f', -1, 64)
		newdata := fmt.Sprintf("%s元", strdata)
		return newdata
	}
	datafloat /= 10000
	str := strconv.FormatFloat(datafloat, 'f', 2, 64)
	str = fmt.Sprintf("%s万", str)
	return str
}

// 判断游戏投注盈亏并格式化输出
func TransitionWinOrLose(data string) string {
	dataFloat, _ := strconv.ParseFloat(data, 64)
	if dataFloat >= 0 {
		if dataFloat >= 100 {
			dataFloat /= 10000
			sData := strconv.FormatFloat(dataFloat, 'f', 2, 64)
			strData := fmt.Sprintf("盈利%s万", sData)
			return strData
		}

		strData := fmt.Sprintf("盈利%v元", int(dataFloat))
		return strData
	} else {
		if dataFloat < -99 {
			dataFloat /= 10000
			sData := strconv.FormatFloat(dataFloat, 'f', 2, 64)
			strData := fmt.Sprintf("亏损%s万", sData)
			strData = strings.Replace(strData, "-", "", 1)
			return strData
		}

		strData := fmt.Sprintf("亏损%v元", int(dataFloat))
		strData = strings.Replace(strData, "-", "", 1)
		return strData
	}
}

// 公司盈亏判定格式化输出
func IsCompanyWinOrLose(dataStr string, HouseName string) string {
	data, err := strconv.ParseFloat(dataStr, 64)
	name := "房主"
	if strings.Contains(HouseName, "公司") {
		name = "分公司"
	}
	if strings.Contains(HouseName, "招商") {
		name = "招商部"
	}

	if err != nil {
		fmt.Println("格式转换出错，core.IsCompanyWinOrLose", err)
	}
	if data >= 0 {
		if data >= 100 {
			data /= 10000
			sdata := strconv.FormatFloat(data, 'f', 2, 64)
			strdata := fmt.Sprintf("总盈利%s万（%s赢钱）", sdata, name)
			return strdata
		}

		strdata := fmt.Sprintf("总盈利%v元（%s赢钱）", int(data), name)
		return strdata
	} else {
		if data > -100 {
			strdata := fmt.Sprintf("总亏损%v元（%s亏钱）", int(data), name)
			strdata = strings.Replace(strdata, "-", "", 1)
			return strdata
		}
		data /= 10000
		sdata := strconv.FormatFloat(data, 'f', 2, 64)
		strdata := fmt.Sprintf("总亏损%s万（%s亏钱）", sdata, name)
		strdata = strings.Replace(strdata, "-", "", 1)
		return strdata
	}
}

// 游戏输赢冒泡排序
func GamesDataSort(data [][]string, row int) {
	for i := 0; i < len(data); i++ {
		for k := i + 1; k < len(data); k++ {
			numI, _ := strconv.ParseFloat(data[i][row], 64)
			numK, _ := strconv.ParseFloat(data[k][row], 64)

			if numI < numK {
				data[i], data[k] = data[k], data[i]
			}
		}
	}
}

func Slicing(data *[][]string, first int, last int) {
	*data = append((*data)[:first], (*data)[len(*data)-last:]...)
}

func MergeSlice(data1 [][]string, data2 [][]string, data3 [][]string) [][]string {
	rows := len(data1)
	cols := len(data1[0])
	MergeData := make([][]string, rows)
	for i := range MergeData {
		MergeData[i] = make([]string, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i == 0 || i == 9 || j == 0 {
				MergeData[i][j] = data1[i][j]
				continue
			}

			data1Float, _ := strconv.ParseFloat(data1[i][j], 64)
			data2Float, _ := strconv.ParseFloat(data2[i][j], 64)
			data3Float, _ := strconv.ParseFloat(data3[i][j], 64)
			res := data1Float + data2Float + data3Float
			MergeData[i][j] = fmt.Sprintf("%v", int(res))

		}
	}
	return MergeData
}
