package controllers

import (
	"data/cores"
	"data/models"
	"data/service"
	"fmt"

	"github.com/xuri/excelize/v2"
)

type WeekDataController struct {
	WeekService   *service.WeekDataService
	TableName     string
	HouseAmount   int
	CompanyAmount int
	F             *excelize.File
}

func NewWeekDataController(CompanyAmount int, HouseAmount int, ExcelPath string, TxtPath string) *WeekDataController {
	return &WeekDataController{
		WeekService: &service.WeekDataService{
			ThisWeekData:         models.ThisWeekDataMap,
			ThisWeekDataMapValue: &models.ThisWeekDataMapValue,
			LastWeekDataMapValue: &models.LastWeekDataMapValue,
			ThisHouseData:        models.ThisHouseData,
			ThisHouseDataValue:   &models.ThisHouseDataValue,
			LastHouseDataValue:   &models.LastHouseDataValue,
			TxtPath:              TxtPath,
		},

		TableName:     "Sheet1",
		CompanyAmount: CompanyAmount,
		HouseAmount:   HouseAmount,
		F:             cores.OpenExcel(ExcelPath),
	}
}

func (w *WeekDataController) WeekDataUser() {

	cores.ClearDocument(w.WeekService.TxtPath)
	for i := 6; i <= w.CompanyAmount; i++ {
		for k, v := range w.WeekService.ThisWeekData {
			element := fmt.Sprintf("%s%v", v, i)
			value := cores.GetExcelValue(w.F, w.TableName, element)
			(*w.WeekService.ThisWeekDataMapValue)[k] = value

			element = fmt.Sprintf("%s%v", v, i-1)
			value = cores.GetExcelValue(w.F, w.TableName, element)
			(*w.WeekService.LastWeekDataMapValue)[k] = value
		}
		w.WeekService.WeekDataFormatContent()
		i += 3
	}

	for j := 28; j <= w.HouseAmount; j++ {
		for k, v := range w.WeekService.ThisHouseData {
			element := fmt.Sprintf("%s%v", v, j)
			value := cores.GetExcelValue(w.F, w.TableName, element)
			(*w.WeekService.ThisHouseDataValue)[k] = value

			element = fmt.Sprintf("%s%v", v, j-1)
			value = cores.GetExcelValue(w.F, w.TableName, element)
			(*w.WeekService.LastHouseDataValue)[k] = value

		}
		w.WeekService.WeekHouseFormatContent()
		if j == 40 || j == 55 {
			j += 6
			continue
		}
		j += 3
	}
}

func (D *WeekDataController) GameWeekData() {
	D.WeekService.WeekGame.ThisGame = cores.GetExcelCols(D.F, "本周游戏数据")
	D.WeekService.WeekGame.LastGame = cores.GetExcelCols(D.F, "上周游戏数据")
}
