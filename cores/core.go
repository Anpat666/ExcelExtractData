package cores

import (
	"fmt"
	"strconv"

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
			return strData
		}

		strData := fmt.Sprintf("亏损%v元", int(dataFloat))
		return strData
	}
}

// 公司盈亏判定格式化输出
func IsCompanyWinOrLose(dataStr string) string {
	data, err := strconv.ParseFloat(dataStr, 64)
	if err != nil {
		fmt.Println("格式转换出错，core.IsCompanyWinOrLose", err)
	}
	if data >= 0 {
		if data >= 100 {
			data /= 10000
			sdata := strconv.FormatFloat(data, 'f', 2, 64)
			strdata := fmt.Sprintf("总盈利%s万（分公司赢钱）", sdata)
			return strdata
		}

		strdata := fmt.Sprintf("总盈利%v元（分公司赢钱）", int(data))
		return strdata
	} else {
		if data > -100 {
			strdata := fmt.Sprintf("总亏损%v元（分公司亏钱）", int(data))
			return strdata
		}
		data /= 10000
		sdata := strconv.FormatFloat(data, 'f', 2, 64)
		strdata := fmt.Sprintf("总亏损%s万（分公司亏钱）", sdata)
		return strdata
	}
}
