package services

import (
	"data/controllers"
	"data/cores"
	"data/models"

	"fmt"

	"github.com/xuri/excelize/v2"
)

type DailyDataService struct {
	DailyElement    *models.DailyDataElement
	DailyController *controllers.DailyDataController
	TableName       string
	HouseAmount     int
	F               *excelize.File
}

func NewDailyDataService(HouseAmount int, ExcelPath string, TxtPath string) *DailyDataService {
	return &DailyDataService{
		DailyElement: &models.DailyDataElement{},
		DailyController: &controllers.DailyDataController{
			DailyData: &models.DailyData{},
			GameData:  &models.GameData{},
			TxtPath:   TxtPath,
		},

		TableName:   "SheetJS",
		HouseAmount: HouseAmount,
		F:           cores.OpenExcel(ExcelPath),
	}
}

func (D *DailyDataService) DailyDataUser() {
	cores.ClearDocument(D.DailyController.TxtPath)
	for i := 2; i <= D.HouseAmount; i++ {
		D.DailyElement.HouseName = fmt.Sprintf("A%v", i)
		D.DailyController.DailyData.HouseName = cores.GetExcelValue(D.F, D.TableName, D.DailyElement.HouseName)

		D.DailyElement.Deposit = fmt.Sprintf("D%v", i)
		D.DailyController.DailyData.Deposit = cores.GetExcelValue(D.F, D.TableName, D.DailyElement.Deposit)

		D.DailyElement.Withdrawal = fmt.Sprintf("J%v", i)
		D.DailyController.DailyData.Withdrawal = cores.GetExcelValue(D.F, D.TableName, D.DailyElement.Withdrawal)

		D.DailyElement.DepSubWith = fmt.Sprintf("K%v", i)
		D.DailyController.DailyData.DepSubWith = cores.GetExcelValue(D.F, D.TableName, D.DailyElement.DepSubWith)

		D.DailyElement.BetTotal = fmt.Sprintf("R%v", i)
		D.DailyController.DailyData.BetTotal = cores.GetExcelValue(D.F, D.TableName, D.DailyElement.BetTotal)

		D.DailyElement.WinAmount = fmt.Sprintf("W%v", i)
		D.DailyController.DailyData.WinAmount = cores.GetExcelValue(D.F, D.TableName, D.DailyElement.WinAmount)

		D.DailyElement.ReturnWater = fmt.Sprintf("P%v", i)
		D.DailyController.DailyData.ReturnWater = cores.GetExcelValue(D.F, D.TableName, D.DailyElement.ReturnWater)

		D.DailyElement.Active = fmt.Sprintf("T%v", i)
		D.DailyController.DailyData.Active = cores.GetExcelValue(D.F, D.TableName, D.DailyElement.Active)

		D.DailyElement.Commission = fmt.Sprintf("X%v", i)
		D.DailyController.DailyData.Commission = cores.GetExcelValue(D.F, D.TableName, D.DailyElement.Commission)

		D.DailyController.GetActiveAounmt()

		D.DailyController.GetWinOrLose()

		D.DailyController.GetProfitAndLoss()

		D.DailyController.FormatDailyDataContent()
	}
}

func (D *DailyDataService) GamesDailyData() {
	by01 := cores.GetExcelRows(D.F, "BY01")
	by02 := cores.GetExcelRows(D.F, "BY02")
	fillerRows := cores.GetExcelRows(D.F, "TEST")

	D.DailyController.GameData.G7Rows = cores.GetExcelRows(D.F, "G7")
	D.DailyController.GameData.YYRows = cores.GetExcelRows(D.F, "YY")
	D.DailyController.GameData.BYRows = cores.MergeSlice(by01, by02, fillerRows)
	D.DailyController.GameData.ZSRows = cores.GetExcelRows(D.F, "ZS")

	cores.GamesDataSort(D.DailyController.GameData.G7Rows, 5)
	cores.GamesDataSort(D.DailyController.GameData.YYRows, 5)
	cores.GamesDataSort(D.DailyController.GameData.BYRows, 5)
	cores.GamesDataSort(D.DailyController.GameData.ZSRows, 5)

	cores.Slicing(&D.DailyController.GameData.G7Rows, 3, 3)
	cores.Slicing(&D.DailyController.GameData.YYRows, 3, 3)
	cores.Slicing(&D.DailyController.GameData.BYRows, 3, 3)
	cores.Slicing(&D.DailyController.GameData.ZSRows, 3, 3)

	D.DailyController.Content = "------G7每日游戏输赢-------\n"
	cores.UpDataReport(D.DailyController.Content, D.DailyController.TxtPath)
	for _, v := range D.DailyController.GameData.G7Rows {
		D.DailyController.FormatGameContent(v[0], v[6], v[5])
	}
	D.DailyController.Content = "------YY每日游戏输赢-------\n"
	cores.UpDataReport(D.DailyController.Content, D.DailyController.TxtPath)
	for _, v := range D.DailyController.GameData.YYRows {
		D.DailyController.FormatGameContent(v[0], v[6], v[5])
	}

	D.DailyController.Content = "------BY每日游戏输赢-------\n"
	cores.UpDataReport(D.DailyController.Content, D.DailyController.TxtPath)
	for _, v := range D.DailyController.GameData.BYRows {
		D.DailyController.FormatGameContent(v[0], v[6], v[5])
	}

	D.DailyController.Content = "------ZS每日游戏输赢-------\n"
	cores.UpDataReport(D.DailyController.Content, D.DailyController.TxtPath)
	for _, v := range D.DailyController.GameData.ZSRows {
		D.DailyController.FormatGameContent(v[0], v[6], v[5])
	}

}
