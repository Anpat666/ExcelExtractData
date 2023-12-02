package service

import (
	"data/cores"
	"fmt"
)

type WeekDataService struct {
	ThisWeekData         map[string]string
	ThisWeekDataMapValue *map[string]string
	LastWeekData         map[string]string
	LastWeekDataMapValue *map[string]string
	Content              string
	TxtPath              string
	Link                 string
	ThisData             string
	LastData             string
	DepSubWith           string
	ContentSlices        []string
}

func (w *WeekDataService) WeekDataFormatContent() {
	w.Content = fmt.Sprintf("--------%s--------\n", (*w.ThisWeekDataMapValue)["HouseName"])
	cores.UpDataReport(w.Content, w.TxtPath)

	w.Content = fmt.Sprintf("1、有效房卡%s个，新增房卡%s个，到期房卡%s个。\n", (*w.ThisWeekDataMapValue)["HouseCard"], (*w.ThisWeekDataMapValue)["HouseAdd"], (*w.ThisWeekDataMapValue)["HouseExpire"])
	cores.UpDataReport(w.Content, w.TxtPath)

	w.Link = cores.ComparisonTool((*w.LastWeekDataMapValue)["PlayerActiveDaily"], (*w.ThisWeekDataMapValue)["PlayerActiveDaily"])
	w.Content = fmt.Sprintf("2、日均活跃人数：上周%s人，本周%s人，%s。 \n", (*w.LastWeekDataMapValue)["PlayerActiveDaily"], (*w.ThisWeekDataMapValue)["PlayerActiveDaily"], w.Link)
	cores.UpDataReport(w.Content, w.TxtPath)

	w.Link = cores.ComparisonTool((*w.LastWeekDataMapValue)["PlayerBetingDaily"], (*w.ThisWeekDataMapValue)["PlayerBetingDaily"])
	w.Content = fmt.Sprintf("3、日均投注人数：上周%s人，本周%s人，%s。\n", (*w.LastWeekDataMapValue)["PlayerBetingDaily"], (*w.ThisWeekDataMapValue)["PlayerBetingDaily"], w.Link)
	cores.UpDataReport(w.Content, w.TxtPath)

	w.ThisData = cores.PerCapitaStatistics((*w.ThisWeekDataMapValue)["Deposit"], (*w.ThisWeekDataMapValue)["PlayerDeposit"])
	w.LastData = cores.PerCapitaStatistics((*w.LastWeekDataMapValue)["Deposit"], (*w.LastWeekDataMapValue)["PlayerDeposit"])
	w.Link = cores.ComparisonTool(w.ThisData, w.LastData)
	w.ThisData = cores.TransitionData(w.ThisData)
	w.LastData = cores.TransitionData(w.LastData)
	w.Content = fmt.Sprintf("4、人均存款金额：上周%s，本周%s，%s。\n", w.ThisData, w.LastData, w.Link)
	cores.UpDataReport(w.Content, w.TxtPath)

	w.ThisData = cores.PerCapitaStatistics((*w.ThisWeekDataMapValue)["BetTotal"], (*w.ThisWeekDataMapValue)["PlayerBeting"])
	w.LastData = cores.PerCapitaStatistics((*w.LastWeekDataMapValue)["BetTotal"], (*w.LastWeekDataMapValue)["PlayerBeting"])
	w.Link = cores.ComparisonTool(w.LastData, w.ThisData)
	w.ThisData = cores.TransitionData(w.ThisData)
	w.LastData = cores.TransitionData(w.LastData)
	w.Content = fmt.Sprintf("5、人均投注量：上周%s，本周%s，%s。\n", w.LastData, w.ThisData, w.Link)
	cores.UpDataReport(w.Content, w.TxtPath)

	w.Link = cores.ComparisonTool((*w.LastWeekDataMapValue)["PlayerAdd"], (*w.ThisWeekDataMapValue)["PlayerAdd"])
	w.Content = fmt.Sprintf("6、新增会员：上周%s人，本周%s人，%s。\n", (*w.LastWeekDataMapValue)["PlayerAdd"], (*w.ThisWeekDataMapValue)["PlayerAdd"], w.Link)
	cores.UpDataReport(w.Content, w.TxtPath)

	w.Link = cores.ComparisonTool((*w.LastWeekDataMapValue)["PlayerActive"], (*w.ThisWeekDataMapValue)["PlayerActive"])
	w.Content = fmt.Sprintf("7、活跃人数：上周%s人，本周%s人，%s。\n", (*w.LastWeekDataMapValue)["PlayerActive"], (*w.ThisWeekDataMapValue)["PlayerActive"], w.Link)
	cores.UpDataReport(w.Content, w.TxtPath)

	w.Link = cores.ComparisonTool((*w.LastWeekDataMapValue)["PlayerBeting"], (*w.ThisWeekDataMapValue)["PlayerBeting"])
	w.Content = fmt.Sprintf("8、投注人数：上周%s人，本周%s人，%s。\n", (*w.LastWeekDataMapValue)["PlayerBeting"], (*w.ThisWeekDataMapValue)["PlayerBeting"], w.Link)
	cores.UpDataReport(w.Content, w.TxtPath)

	w.ThisData = cores.TransitionData((*w.ThisWeekDataMapValue)["BetTotal"])
	w.LastData = cores.TransitionData((*w.LastWeekDataMapValue)["BetTotal"])
	w.Link = cores.ComparisonTool((*w.LastWeekDataMapValue)["BetTotal"], (*w.ThisWeekDataMapValue)["BetTotal"])
	w.Content = fmt.Sprintf("9、投注量：上周%s，本周%s，%s。\n", w.LastData, w.ThisData, w.Link)
	cores.UpDataReport(w.Content, w.TxtPath)

	w.ThisData = cores.TransitionData((*w.ThisWeekDataMapValue)["Deposit"])
	w.LastData = cores.TransitionData((*w.LastWeekDataMapValue)["Deposit"])
	w.Link = cores.ComparisonTool((*w.LastWeekDataMapValue)["Deposit"], (*w.ThisWeekDataMapValue)["Deposit"])
	w.Content = fmt.Sprintf("10、存款量：上周%s，本周%s，%s。\n", w.LastData, w.ThisData, w.Link)
	cores.UpDataReport(w.Content, w.TxtPath)

	w.ThisData = cores.TransitionData((*w.ThisWeekDataMapValue)["Withdrawal"])
	w.LastData = cores.TransitionData((*w.LastWeekDataMapValue)["Withdrawal"])
	w.Link = cores.ComparisonTool((*w.LastWeekDataMapValue)["Withdrawal"], (*w.ThisWeekDataMapValue)["Withdrawal"])
	w.Content = fmt.Sprintf("11、取款量：上周%s，本周%s，%s。\n", w.LastData, w.ThisData, w.Link)
	cores.UpDataReport(w.Content, w.TxtPath)

	w.ThisData = cores.TransitionData((*w.ThisWeekDataMapValue)["WinOrLose"])
	w.LastData = cores.TransitionData((*w.LastWeekDataMapValue)["WinOrLose"])
	w.Link = cores.ComparisonTool((*w.LastWeekDataMapValue)["WinOrLose"], (*w.ThisWeekDataMapValue)["WinOrLose"])
	w.Content = fmt.Sprintf("12、游戏盈利：上周%s，本周%s，%s。\n", w.LastData, w.ThisData, w.Link)
	cores.UpDataReport(w.Content, w.TxtPath)

	w.ThisData = cores.TransitionData((*w.ThisWeekDataMapValue)["ProfitAndLoss"])
	w.LastData = cores.TransitionData((*w.LastWeekDataMapValue)["ProfitAndLoss"])
	w.Link = cores.ComparisonTool((*w.LastWeekDataMapValue)["ProfitAndLoss"], (*w.ThisWeekDataMapValue)["ProfitAndLoss"])
	w.Content = fmt.Sprintf("13、总盈利：上周%s，本周%s，%s。\n", w.LastData, w.ThisData, w.Link)
	cores.UpDataReport(w.Content, w.TxtPath)

	w.DepSubWith = cores.DepositsSubWithdrawal((*w.ThisWeekDataMapValue)["Deposit"], (*w.ThisWeekDataMapValue)["Withdrawal"])
	w.Link = cores.DepositsAndBetProportion((*w.ThisWeekDataMapValue)["Deposit"], (*w.ThisWeekDataMapValue)["BetTotal"])
	w.DepSubWith = cores.TransitionData(w.DepSubWith)
	w.Content = fmt.Sprintf("14、存提差%s，充投比%s。\n", w.DepSubWith, w.Link)
	cores.UpDataReport(w.Content, w.TxtPath)

	(*w.ThisWeekDataMapValue)["Active"] = cores.TransitionData((*w.ThisWeekDataMapValue)["Active"])
	(*w.ThisWeekDataMapValue)["ReturnWater"] = cores.TransitionData((*w.ThisWeekDataMapValue)["ReturnWater"])
	(*w.ThisWeekDataMapValue)["Commission"] = cores.TransitionData((*w.ThisWeekDataMapValue)["Commission"])
	w.Content = fmt.Sprintf("15、营销活动%s，返水%s,返佣%s。\n", (*w.ThisWeekDataMapValue)["Active"], (*w.ThisWeekDataMapValue)["ReturnWater"], (*w.ThisWeekDataMapValue)["Commission"])
	cores.UpDataReport(w.Content, w.TxtPath)

}
