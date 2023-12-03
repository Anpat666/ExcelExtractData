package controllers

import (
	"data/cores"
	"data/models"
	"data/service"
	"fmt"

	"github.com/xuri/excelize/v2"
)

type DailyDataController struct {
	DailyElement *models.DailyDataElement
	DailyService *service.ServiceDailyData
	TableName    string
	HouseAmount  int
	F            *excelize.File
}

func NewDailyDataController(HouseAmount int, ExcelPath string, TxtPath string) *DailyDataController {
	return &DailyDataController{
		DailyElement: &models.DailyDataElement{},
		DailyService: &service.ServiceDailyData{
			DailyData: &models.DailyData{},
			GameData:  &models.GameData{},
			TxtPath:   TxtPath,
		},

		TableName:   "SheetJS",
		HouseAmount: HouseAmount,
		F:           cores.OpenExcel(ExcelPath),
	}
}

func (D *DailyDataController) DailyDataUser() {
	cores.ClearDocument(D.DailyService.TxtPath)
	for i := 2; i <= D.HouseAmount; i++ {
		D.DailyElement.HouseName = fmt.Sprintf("A%v", i)
		D.DailyService.DailyData.HouseName = cores.GetExcelValue(D.F, D.TableName, D.DailyElement.HouseName)

		D.DailyElement.Deposit = fmt.Sprintf("D%v", i)
		D.DailyService.DailyData.Deposit = cores.GetExcelValue(D.F, D.TableName, D.DailyElement.Deposit)

		D.DailyElement.Withdrawal = fmt.Sprintf("J%v", i)
		D.DailyService.DailyData.Withdrawal = cores.GetExcelValue(D.F, D.TableName, D.DailyElement.Withdrawal)

		D.DailyElement.DepSubWith = fmt.Sprintf("K%v", i)
		D.DailyService.DailyData.DepSubWith = cores.GetExcelValue(D.F, D.TableName, D.DailyElement.DepSubWith)

		D.DailyElement.BetTotal = fmt.Sprintf("R%v", i)
		D.DailyService.DailyData.BetTotal = cores.GetExcelValue(D.F, D.TableName, D.DailyElement.BetTotal)

		D.DailyElement.WinAmount = fmt.Sprintf("W%v", i)
		D.DailyService.DailyData.WinAmount = cores.GetExcelValue(D.F, D.TableName, D.DailyElement.WinAmount)

		D.DailyElement.ReturnWater = fmt.Sprintf("P%v", i)
		D.DailyService.DailyData.ReturnWater = cores.GetExcelValue(D.F, D.TableName, D.DailyElement.ReturnWater)

		D.DailyElement.Active = fmt.Sprintf("T%v", i)
		D.DailyService.DailyData.Active = cores.GetExcelValue(D.F, D.TableName, D.DailyElement.Active)

		D.DailyElement.Commission = fmt.Sprintf("X%v", i)
		D.DailyService.DailyData.Commission = cores.GetExcelValue(D.F, D.TableName, D.DailyElement.Commission)

		D.DailyService.GetActiveAounmt()

		D.DailyService.GetWinOrLose()

		D.DailyService.GetProfitAndLoss()

		D.DailyService.FormatDailyDataContent()
	}
}

func (D *DailyDataController) GamesDailyData() {
	D.DailyService.GameData.G7Rows = cores.GetExcelRows(D.F, "G7")
	D.DailyService.GameData.YYRows = cores.GetExcelRows(D.F, "YY")
	D.DailyService.GameData.BYRows = cores.GetExcelRows(D.F, "BY")

	cores.GamesDataSort(D.DailyService.GameData.G7Rows, 5)
	cores.GamesDataSort(D.DailyService.GameData.YYRows, 5)
	cores.GamesDataSort(D.DailyService.GameData.BYRows, 5)

	cores.Slicing(&D.DailyService.GameData.G7Rows, 3, 3)
	cores.Slicing(&D.DailyService.GameData.YYRows, 3, 3)
	cores.Slicing(&D.DailyService.GameData.BYRows, 3, 3)

	D.DailyService.Content = "------G7每日游戏输赢-------\n"
	cores.UpDataReport(D.DailyService.Content, D.DailyService.TxtPath)
	for _, v := range D.DailyService.GameData.G7Rows {
		D.DailyService.FormatGameContent(v[0], v[6], v[5])
	}
	D.DailyService.Content = "------YY每日游戏输赢-------\n"
	cores.UpDataReport(D.DailyService.Content, D.DailyService.TxtPath)
	for _, v := range D.DailyService.GameData.YYRows {
		D.DailyService.FormatGameContent(v[0], v[6], v[5])
	}

	D.DailyService.Content = "------BY每日游戏输赢-------\n"
	cores.UpDataReport(D.DailyService.Content, D.DailyService.TxtPath)
	for _, v := range D.DailyService.GameData.BYRows {
		D.DailyService.FormatGameContent(v[0], v[6], v[5])
	}

}
