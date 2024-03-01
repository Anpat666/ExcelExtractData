package cores

import (
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func MakeNewDataExcel(this [][]string, last [][]string) {

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	index, err := f.NewSheet("Sheet2")
	if err != nil {
		fmt.Println(err)
		return
	}

	f.SetSheetCol("Sheet2", "A1", &this[0])
	f.SetSheetCol("Sheet2", "B1", &this[1])
	f.SetSheetCol("Sheet2", "C1", &this[6])
	f.SetSheetCol("Sheet2", "D1", &last[6])
	f.SetSheetCol("Sheet2", "F1", &this[8])
	f.SetSheetCol("Sheet2", "G1", &this[5])
	f.SetSheetCol("Sheet2", "H1", &this[7])
	f.SetSheetCol("Sheet2", "J1", &this[9])
	f.SetCellDefault("Sheet2", "D1", "上周有效流水")
	f.SetCellDefault("Sheet2", "E1", "投注量成长（与上周对比）")
	f.SetCellDefault("Sheet2", "H1", "实际输赢")
	f.SetCellDefault("Sheet2", "I1", "杀率（未计算返水）")

	var cellformula string

	for i := 2; i < len(this[0])+1; i++ {
		cell := this[5][i-1]
		cell2 := this[6][i-1]
		cellfloat, _ := strconv.ParseFloat(cell, 64)
		cell2float, _ := strconv.ParseFloat(cell2, 64)
		res := cellfloat / cell2float
		res *= 100
		resStr := strconv.FormatFloat(res, 'f', 2, 64)
		newStr := fmt.Sprintf("%s%%", resStr)
		cellformula = fmt.Sprintf("I%v", i)
		f.SetCellValue("Sheet2", cellformula, newStr)

		thisCell := this[6][i-1]
		lastCELL := last[6][i-1]
		thisFloat, _ := strconv.ParseFloat(thisCell, 64)
		lastFloat, _ := strconv.ParseFloat(lastCELL, 64)
		res2 := thisFloat/lastFloat - 1
		res2 *= 100
		res2Str := strconv.FormatFloat(res2, 'f', 2, 64)
		new2Str := fmt.Sprintf("%s%%", res2Str)
		cellformula = fmt.Sprintf("E%v", i)
		f.SetCellValue("Sheet2", cellformula, new2Str)

	}
	toBeRemoved := []int{}

	for k, v := range this[6] {
		if v == "0" {
			toBeRemoved = append(toBeRemoved, k+index)
		}
	}
	for i := len(toBeRemoved) - 1; i >= 0; i-- {
		f.RemoveRow("Sheet2", toBeRemoved[i])
	}

	var BetAmount float64
	var betTotal float64
	var lastBetTotal float64
	var returnWater float64
	var winOrlose float64
	var profit float64
	var betTatalLink float64
	var winDivbet float64

	for i := 0; i < len(this[0]); i++ {
		BetAmountFloat, _ := strconv.ParseFloat(this[1][i], 64)
		BetAmount += BetAmountFloat

		betTotalFloat, _ := strconv.ParseFloat(this[6][i], 64)
		betTotal += betTotalFloat

		lastBetTotalFloat, _ := strconv.ParseFloat(last[6][i], 64)
		lastBetTotal += lastBetTotalFloat

		returnWaterFloat, _ := strconv.ParseFloat(this[8][i], 64)
		returnWater += returnWaterFloat

		winOrloseFloat, _ := strconv.ParseFloat(this[5][i], 64)
		winOrlose += winOrloseFloat

		profitFloat, _ := strconv.ParseFloat(this[7][i], 64)
		profit += profitFloat
	}

	betTatalLink = betTotal/lastBetTotal - 1
	betTatalLink *= 100
	winDivbet = winOrlose / betTotal

	betTatalLinkStr := strconv.FormatFloat(betTatalLink, 'f', 2, 64)
	betTatalLinkStr = fmt.Sprintf("%s%%", betTatalLinkStr)

	winDivbetStr := strconv.FormatFloat(winDivbet, 'f', 2, 64)
	winDivbetStr = fmt.Sprintf("%s%%", winDivbetStr)

	f.InsertRows("Sheet2", 2, 1)
	f.SetCellValue("Sheet2", "A2", "总计")
	f.SetCellValue("Sheet2", "B2", BetAmount)
	f.SetCellValue("Sheet2", "C2", betTotal)
	f.SetCellValue("Sheet2", "D2", lastBetTotal)
	f.SetCellValue("Sheet2", "E2", betTatalLinkStr)
	f.SetCellValue("Sheet2", "F2", returnWater)
	f.SetCellValue("Sheet2", "G2", winOrlose)
	f.SetCellValue("Sheet2", "H2", profit)
	f.SetCellValue("Sheet2", "I2", winDivbetStr)

	// 设置工作簿的默认工作表
	f.SetActiveSheet(index)
	// 根据指定路径保存文件
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
	f.Close()

}

func GamesDataExcelSort() ([][]string, [][]string) {
	Sheet2 := "Sheet2"
	f := OpenExcel("Book1.xlsx")
	rows, err := f.GetRows(Sheet2)
	if err != nil {
		fmt.Println(err)
	}
	newRows := rows[2:]

	for i := 0; i < len(newRows)-1; i++ {
		for k := i + 1; k < len(newRows); k++ {
			numI, _ := strconv.ParseFloat(newRows[i][7], 64)
			numK, _ := strconv.ParseFloat(newRows[k][7], 64)

			if numI < numK {
				newRows[i], newRows[k] = newRows[k], newRows[i]
			}
		}
	}

	f.SetSheetRow("Sheet1", "A1", &rows[0])
	f.SetSheetRow("Sheet1", "A2", &rows[1])

	for i := 0; i < len(newRows); i++ {
		index := fmt.Sprintf("A%v", i+3)
		f.SetSheetRow("Sheet1", index, &newRows[i])
	}
	f.DeleteSheet("Sheet2")

	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
	f.Close()

	return newRows, rows
}

func GameCategory(data [][]string) [][]string {
	dataClass1 := make([]string, len(data[0]))
	dataClass2 := make([]string, len(data[0]))
	dataClass1[0] = "分类"
	dataClass2[0] = "细分"
	baccarat := []string{"MC斗牛01", "MC龙虎斗03", "MC斗牛02", "MC百家乐A01", "MC百家乐A04", "MC龙虎斗01",
		"MC百家乐V01", "MC百家乐V03", "MC百家乐09", "MC百家乐02", "MC真人百家乐", "MC龙虎斗02", "MC百家乐03",
		"MC百家乐A03", "MC百家乐06", "MC百家乐V02", "MC百家乐05", "MC百家乐V04", "MC百家乐A02", "MC百家乐07",
		"MC百家乐04", "MC百家乐08", "MC百家乐01"}

	FastThree := []string{"MC十分快三", "MC五分快三", "MC一分快三"}
	NumFive := []string{"极速时时彩", "MC秒速时时彩", "澳洲幸运5"}
	NumTen := []string{"MC秒速赛车", "幸运赛车", "澳洲幸运10", "幸运飞艇", "极速赛车", "极速飞艇"}
	MarkSix := []string{"台湾大乐透", "澳门六合彩", "MC十分六合彩", "香港六合彩", "MC三分六合彩", "台湾六合彩", "MC五分六合彩"}
	Other := []string{"加拿大PC28"}

	for i := 0; i < len(data[0]); i++ {
		gameName := data[0][i]
		if contains(baccarat, gameName) {
			dataClass1[i] = "1"
			dataClass2[i] = "1"
		}
		if contains(FastThree, gameName) {
			dataClass1[i] = "2"
			dataClass2[i] = "3"
		}
		if contains(NumFive, gameName) {
			dataClass1[i] = "2"
			dataClass2[i] = "5"
		}
		if contains(NumTen, gameName) {
			dataClass1[i] = "2"
			dataClass2[i] = "10"
		}
		if contains(MarkSix, gameName) {
			dataClass1[i] = "2"
			dataClass2[i] = "6"
		}
		if contains(Other, gameName) {
			dataClass1[i] = "2"
			dataClass2[i] = "7"
		}
	}

	data = append(data, dataClass1, dataClass2)
	return data

}

func contains(list []string, target string) bool {
	for _, value := range list {
		if value == target {
			return true
		}
	}
	return false
}
