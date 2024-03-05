package services

import (
	"data/controllers"
	"data/cores"
	"data/models"
	"fmt"
	"reflect"
	"strconv"

	"github.com/xuri/excelize/v2"
)

type WeekDataService struct {
	WeekController    *controllers.WeekDataController
	TableName         string
	HouseAmount       int
	CompanyAmount     int
	HouseDataStarLin  int
	ThisDataTableName string
	LastDataTableName string

	F *excelize.File
}

func NewWeekDataService(CompanyAmount int, HouseAmount int, HouseDataStarLin int, ExcelPath string, TxtPath string, ThisDataTableName string, LastDataTableName string) *WeekDataService {
	return &WeekDataService{
		WeekController: &controllers.WeekDataController{
			WeekGame:             &models.GameWeek{},
			GamesExcel:           &models.GamesExcel{},
			ThisWeekData:         models.ThisWeekDataMap,
			ThisWeekDataMapValue: &models.ThisWeekDataMapValue,
			LastWeekDataMapValue: &models.LastWeekDataMapValue,
			ThisHouseData:        models.ThisHouseData,
			ThisHouseDataValue:   &models.ThisHouseDataValue,
			LastHouseDataValue:   &models.LastHouseDataValue,
			TotalGameClass:       &models.GameClass{},
			G7GameClass:          &models.GameClass{},
			YYGameClass:          &models.GameClass{},
			BYGameClass:          &models.GameClass{},
			ZSGameClass:          &models.GameClass{},
			TxtPath:              TxtPath,
		},

		TableName:         "Sheet1",
		CompanyAmount:     CompanyAmount,
		HouseAmount:       HouseAmount,
		HouseDataStarLin:  HouseDataStarLin,
		ThisDataTableName: ThisDataTableName,
		LastDataTableName: LastDataTableName,

		F: cores.OpenExcel(ExcelPath),
	}
}

func (w *WeekDataService) WeekDataService() {

	cores.ClearDocument(w.WeekController.TxtPath)
	for i := 6; i <= w.CompanyAmount; i++ {
		for k, v := range w.WeekController.ThisWeekData {
			element := fmt.Sprintf("%s%v", v, i)
			value := cores.GetExcelValue(w.F, w.TableName, element)
			(*w.WeekController.ThisWeekDataMapValue)[k] = value

			element = fmt.Sprintf("%s%v", v, i-1)
			value = cores.GetExcelValue(w.F, w.TableName, element)
			(*w.WeekController.LastWeekDataMapValue)[k] = value
		}
		w.WeekController.WeekDataFormatContent()
		i += 3
	}

	for j := w.HouseDataStarLin; j <= w.HouseAmount; j++ {
		for k, v := range w.WeekController.ThisHouseData {
			element := fmt.Sprintf("%s%v", v, j)
			value := cores.GetExcelValue(w.F, w.TableName, element)
			(*w.WeekController.ThisHouseDataValue)[k] = value

			element = fmt.Sprintf("%s%v", v, j-1)
			value = cores.GetExcelValue(w.F, w.TableName, element)
			(*w.WeekController.LastHouseDataValue)[k] = value
		}
		w.WeekController.WeekHouseFormatContent()
		j += 3
	}
}

func (D *WeekDataService) GameWeekDataService() {

	D.WeekController.WeekGame.ThisGame = cores.GetExcelCols(D.F, D.ThisDataTableName)
	D.WeekController.WeekGame.LastGame = cores.GetExcelCols(D.F, D.LastDataTableName)

	D.WeekController.WeekGame.ThisGame = cores.GameCategory(D.WeekController.WeekGame.ThisGame)
	D.WeekController.WeekGame.LastGame = cores.GameCategory(D.WeekController.WeekGame.LastGame)

	cores.MakeNewDataExcel(D.WeekController.WeekGame.ThisGame, D.WeekController.WeekGame.LastGame)
	newRows, Totalrows := cores.GamesDataExcelSort()
	D.F = cores.OpenExcel("Book1.xlsx")

	// 获取投注量  名称+百分比
	cores.GamesDataSort(newRows, 2)
	betTotal, _ := strconv.ParseFloat(Totalrows[1][2], 64)
	cores.Slicing(&newRows, 3, 3)
	D.WeekController.WeekGame.BetingBest = make([][]string, 6)
	for i := 0; i < len(newRows); i++ {
		D.WeekController.WeekGame.BetingBest[i] = append(D.WeekController.WeekGame.BetingBest[i], newRows[i][0])
		dataFloat, _ := strconv.ParseFloat(newRows[i][2], 64)
		resBeting := dataFloat / betTotal
		resBeting *= 100
		strBeting := strconv.FormatFloat(resBeting, 'f', 2, 64)
		strBeting = fmt.Sprintf("%s%%", strBeting)
		D.WeekController.WeekGame.BetingBest[i] = append(D.WeekController.WeekGame.BetingBest[i], strBeting)
	}

	//获取输赢 前3+后3 名字
	newGamesData := cores.GetExcelRows(D.F, D.TableName)
	winRoLoseNew := make([][]string, len(newGamesData)-2)
	copy(winRoLoseNew, newGamesData[2:])
	cores.GamesDataSort(winRoLoseNew, 7)
	cores.Slicing(&winRoLoseNew, 3, 3)
	D.WeekController.WeekGame.WinOrLose = make([][]string, 6)
	for i := 0; i < len(winRoLoseNew); i++ {
		D.WeekController.WeekGame.WinOrLose[i] = append(D.WeekController.WeekGame.WinOrLose[i], winRoLoseNew[i][0])
	}

	//获取投注量成长  名称+投注量差值
	BetingRate := make([][]string, len(newGamesData)-2)
	copy(BetingRate, newGamesData[2:])
	for i := 0; i < len(BetingRate); i++ {
		this, _ := strconv.Atoi(BetingRate[i][2])
		last, _ := strconv.Atoi(BetingRate[i][3])
		resultSub := this - last
		strNew := fmt.Sprintf("%v", resultSub)
		BetingRate[i] = append(BetingRate[i], strNew)
	}
	cores.GamesDataSort(BetingRate, 10)
	cores.Slicing(&BetingRate, 3, 3)

	D.WeekController.WeekGame.BetingRate = make([][]string, 6)
	for k := range BetingRate {
		D.WeekController.WeekGame.BetingRate[k] = append(D.WeekController.WeekGame.BetingRate[k], BetingRate[k][0])
		resultSub, _ := strconv.ParseFloat(BetingRate[k][10], 64)
		resultSub /= 10000
		resStr := fmt.Sprintf("%v万", int(resultSub))
		D.WeekController.WeekGame.BetingRate[k] = append(D.WeekController.WeekGame.BetingRate[k], resStr)
	}

	TypeGamesTotal := make([][]string, len(newGamesData)-2)
	copy(TypeGamesTotal, newGamesData[2:])

	D.WeekController.GamesExcel.Lottery = make([][]string, 0)
	D.WeekController.GamesExcel.Video = make([][]string, 0)
	var LotteryWinTotal float64
	var videoWinTotal float64
	for k := range TypeGamesTotal {
		betTotal, err := strconv.Atoi(TypeGamesTotal[k][2])
		winOrFloat, _ := strconv.ParseFloat(TypeGamesTotal[k][7], 64)
		D.WeekController.GamesExcel.WinOrLose += winOrFloat
		if err != nil {
			fmt.Println("betotal类型转换失败", err)
		}
		D.WeekController.GamesExcel.BetTotal += betTotal
		if TypeGamesTotal[k][9] == "1" {
			videoWinTotal += winOrFloat
			D.WeekController.GamesExcel.VideoTotal += betTotal
			D.WeekController.GamesExcel.Video = append(D.WeekController.GamesExcel.Video, TypeGamesTotal[k])
		}
		if TypeGamesTotal[k][9] == "2" {
			LotteryWinTotal += winOrFloat
			D.WeekController.GamesExcel.LotteryTotal += betTotal
			D.WeekController.GamesExcel.Lottery = append(D.WeekController.GamesExcel.Lottery, TypeGamesTotal[k])
		}
	}

	newStrWin := fmt.Sprintf("%f", D.WeekController.GamesExcel.WinOrLose)
	newStrLotteryWin := fmt.Sprintf("%f", LotteryWinTotal)
	newStrVideoWin := fmt.Sprintf("%f", videoWinTotal)
	D.WeekController.GamesExcel.LotteryWinPro = cores.ProPortion(newStrWin, newStrLotteryWin)
	D.WeekController.GamesExcel.VideoWinPro = cores.ProPortion(newStrWin, newStrVideoWin)

	cores.GamesDataSort(D.WeekController.GamesExcel.Video, 2)
	cores.Slicing(&D.WeekController.GamesExcel.Video, 3, 0)

	cores.GamesDataSort(D.WeekController.GamesExcel.Lottery, 2)
	cores.Slicing(&D.WeekController.GamesExcel.Lottery, 3, 0)

	betTotalnew := strconv.Itoa(D.WeekController.GamesExcel.BetTotal)
	ltTotalnew := strconv.Itoa(D.WeekController.GamesExcel.LotteryTotal)
	vdTotalnew := strconv.Itoa(D.WeekController.GamesExcel.VideoTotal)
	D.WeekController.GamesExcel.LotteryPro = cores.ProPortion(betTotalnew, ltTotalnew)
	D.WeekController.GamesExcel.VideoPro = cores.ProPortion(betTotalnew, vdTotalnew)

	for k := range D.WeekController.GamesExcel.Lottery {
		lotteryBetol := float64(D.WeekController.GamesExcel.LotteryTotal)
		games, _ := strconv.ParseFloat(D.WeekController.GamesExcel.Lottery[k][2], 64)
		proRes := games / lotteryBetol
		proRes *= 100
		strPro := strconv.FormatFloat(proRes, 'f', 2, 64)
		newStrPro := fmt.Sprintf("%s%%", strPro)
		D.WeekController.GamesExcel.Lottery[k][1] = newStrPro
	}

	for k := range D.WeekController.GamesExcel.Video {
		videoBetol := float64(D.WeekController.GamesExcel.VideoTotal)
		games, _ := strconv.ParseFloat(D.WeekController.GamesExcel.Video[k][2], 64)
		proRes := games / videoBetol
		proRes *= 100
		strPro := strconv.FormatFloat(proRes, 'f', 2, 64)
		newStrPro := fmt.Sprintf("%s%%", strPro)
		D.WeekController.GamesExcel.Video[k][1] = newStrPro
	}

	D.WeekController.GamesDataFormatTxt()
}

func (D *WeekDataService) GameClassService() {
	D.F = cores.OpenExcel("document/周报数据表.xlsx")
	Total := cores.GetExcelRows(D.F, "本周游戏数据")
	G7 := cores.GetExcelRows(D.F, "G7")
	YY := cores.GetExcelRows(D.F, "YY")
	by01 := cores.GetExcelRows(D.F, "BY01")
	by02 := cores.GetExcelRows(D.F, "BY02")
	fillerRows := cores.GetExcelRows(D.F, "TEST")
	BY := cores.MergeSlice(by01, by02, fillerRows)

	Total = cores.GameCategoryRows(Total)
	G7 = cores.GameCategoryRows(G7)
	YY = cores.GameCategoryRows(YY)
	BY = cores.GameCategoryRows(BY)

	D.WeekController.TotalGameClass.Total.BetTotal = cores.SumColumnRowValues(Total, 6)
	D.WeekController.TotalGameClass.Total.OrderAmount = cores.SumColumnRowValues(Total, 1)
	D.WeekController.TotalGameClass.Total.Profit = cores.SumColumnRowValues(Total, 5)
	D.WeekController.TotalGameClass.Total.WinRate = D.WeekController.TotalGameClass.Total.Profit / D.WeekController.TotalGameClass.Total.BetTotal

	ReflectGameClassValue(Total, &D.WeekController.TotalGameClass.TenGame, "1")
	ReflectGameClassValue(Total, &D.WeekController.TotalGameClass.Video, "2")
	ReflectGameClassValue(Total, &D.WeekController.TotalGameClass.Other, "3")
	ReflectGameClassValue(Total, &D.WeekController.TotalGameClass.FiveGame, "4")
	ReflectGameClassValue(Total, &D.WeekController.TotalGameClass.SixMark, "5")

	ReflectGameTotalValue(G7, &D.WeekController.G7GameClass.Total, D.WeekController.TotalGameClass.Total.BetTotal)
	ReflectGameClassValue(G7, &D.WeekController.G7GameClass.TenGame, "1")
	ReflectGameClassValue(G7, &D.WeekController.G7GameClass.Video, "2")
	ReflectGameClassValue(G7, &D.WeekController.G7GameClass.Other, "3")
	ReflectGameClassValue(G7, &D.WeekController.G7GameClass.FiveGame, "4")
	ReflectGameClassValue(G7, &D.WeekController.G7GameClass.SixMark, "5")

	ReflectGameTotalValue(YY, &D.WeekController.YYGameClass.Total, D.WeekController.TotalGameClass.Total.BetTotal)
	ReflectGameClassValue(YY, &D.WeekController.YYGameClass.TenGame, "1")
	ReflectGameClassValue(YY, &D.WeekController.YYGameClass.Video, "2")
	ReflectGameClassValue(YY, &D.WeekController.YYGameClass.Other, "3")
	ReflectGameClassValue(YY, &D.WeekController.YYGameClass.FiveGame, "4")
	ReflectGameClassValue(YY, &D.WeekController.YYGameClass.SixMark, "5")

	ReflectGameTotalValue(BY, &D.WeekController.BYGameClass.Total, D.WeekController.TotalGameClass.Total.BetTotal)
	ReflectGameClassValue(BY, &D.WeekController.BYGameClass.TenGame, "1")
	ReflectGameClassValue(BY, &D.WeekController.BYGameClass.Video, "2")
	ReflectGameClassValue(BY, &D.WeekController.BYGameClass.Other, "3")
	ReflectGameClassValue(BY, &D.WeekController.BYGameClass.FiveGame, "4")
	ReflectGameClassValue(BY, &D.WeekController.BYGameClass.SixMark, "5")

	D.WeekController.GameClassDataFormatTxt()

}

func ReflectGameClassValue(data [][]string, GameClass *models.GameClassBasic, num string) {
	reflectValue := reflect.ValueOf(GameClass).Elem()
	Bet := cores.SumColumnRowValues(data, 6)
	for i := 0; i < reflectValue.NumField(); i++ {
		fieldName := reflectValue.Type().Field(i).Name

		switch fieldName {
		case "OrderAmount":
			GameClass.OrderAmount = cores.SumColumnRowCategory(data, 1, num)
		case "BetTotal":
			GameClass.BetTotal = cores.SumColumnRowCategory(data, 6, num)
		case "Profit":
			GameClass.Profit = cores.SumColumnRowCategory(data, 5, num)
		case "WinRate":
			GameClass.WinRate = GameClass.Profit / GameClass.BetTotal
		case "BettingPro":
			GameClass.BettingPro = GameClass.BetTotal / Bet
		}
	}
}

func ReflectGameTotalValue(data [][]string, GameClass *models.GameClassBasic, BetAmount float64) {
	reflectValue := reflect.ValueOf(GameClass).Elem()
	for i := 0; i < reflectValue.NumField(); i++ {
		fieldName := reflectValue.Type().Field(i).Name

		switch fieldName {
		case "OrderAmount":
			GameClass.OrderAmount = cores.SumColumnRowValues(data, 1)
		case "BetTotal":
			GameClass.BetTotal = cores.SumColumnRowValues(data, 6)
		case "Profit":
			GameClass.Profit = cores.SumColumnRowValues(data, 5)
		case "WinRate":
			GameClass.WinRate = GameClass.Profit / GameClass.BetTotal
		case "BettingPro":
			GameClass.BettingPro = GameClass.BetTotal / BetAmount
		}
	}
}
