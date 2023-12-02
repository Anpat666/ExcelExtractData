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

func NewWeekDataController(CompanyAmount int, HouseAmount int, ExcelPath string) *WeekDataController {
	return &WeekDataController{
		WeekService: &service.WeekDataService{
			ThisWeekData:         models.ThisWeekDataMap,
			ThisWeekDataMapValue: &models.ThisWeekDataMapValue,
			LastWeekDataMapValue: &models.LastWeekDataMapValue,
			TxtPath:              "document/weekDataReport.txt",
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
}
