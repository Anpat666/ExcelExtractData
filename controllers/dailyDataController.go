package controllers

import (
	"data/cores"
	"data/models"
	"data/service"
	"fmt"

	"github.com/xuri/excelize/v2"
)

type DailyDataController struct {
	DailyElement  *models.DailyDataElement
	DailyService  *service.ServiceDailyData
	GameElement   *models.GameElement
	TableName     string
	GameTableName string
	HouseAmount   int
	GameAmount    int
	F             *excelize.File
}

func NewDailyDataController(HouseAmount int, GameAmount int, ExcelPath string) *DailyDataController {
	return &DailyDataController{
		DailyElement: &models.DailyDataElement{},
		DailyService: &service.ServiceDailyData{
			DailyData: &models.DailyData{},
			GameData:  &models.GameData{},
			TxtPath:   "document/dailyDataReport.txt",
		},
		GameElement:   &models.GameElement{},
		TableName:     "SheetJS",
		GameTableName: "Sheet1",
		HouseAmount:   HouseAmount,
		GameAmount:    GameAmount,
		F:             cores.OpenExcel(ExcelPath),
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

	for k := 1; k <= D.GameAmount; k++ {
		D.GameElement.GameName = fmt.Sprintf("A%v", k)
		D.DailyService.GameData.GameName = cores.GetExcelValue(D.F, D.GameTableName, D.GameElement.GameName)

		D.GameElement.BetTotal = fmt.Sprintf("B%v", k)
		D.DailyService.GameData.BetTotal = cores.GetExcelValue(D.F, D.GameTableName, D.GameElement.BetTotal)

		D.GameElement.WinOrLose = fmt.Sprintf("C%v", k)
		D.DailyService.GameData.WinOrLose = cores.GetExcelValue(D.F, D.GameTableName, D.GameElement.WinOrLose)

		D.DailyService.FormatGameContent()
	}
}
