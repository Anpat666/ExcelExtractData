package service

import (
	"data/cores"
	"data/models"
	"fmt"
)

type WeekDataService struct {
	WeekGame             *models.GameWeek
	ThisWeekData         map[string]string
	ThisWeekDataMapValue *map[string]string
	LastWeekDataMapValue *map[string]string
	ThisHouseData        map[string]string
	ThisHouseDataValue   *map[string]string
	LastHouseDataValue   *map[string]string
	Content              string
	TxtPath              string
	Link                 string
	ThisData             string
	LastData             string
	DepSubWith           string
	PerCapDep            string
	PerCapBet            string
	StrBetTotal          string
	StrWinOrlose         string
	StrDeposit           string
	StrWithdrawal        string
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

func (w *WeekDataService) WeekHouseFormatContent() {
	w.Content = fmt.Sprintf("--------%s--------\n", (*w.ThisHouseDataValue)["HouseName"])
	cores.UpDataReport(w.Content, w.TxtPath)

	w.Content = fmt.Sprintf("1、新增会员%s人，周活跃人数%s人，周投注人数%s人。\n",
		(*w.ThisHouseDataValue)["PlayerAdd"], (*w.ThisHouseDataValue)["PlayerActive"], (*w.ThisHouseDataValue)["PlayerBeting"])
	cores.UpDataReport(w.Content, w.TxtPath)

	w.PerCapDep = cores.PerCapitaStatistics((*w.ThisHouseDataValue)["Deposit"], (*w.ThisHouseDataValue)["PlayerDeposit"])
	w.PerCapBet = cores.PerCapitaStatistics((*w.ThisHouseDataValue)["BetTotal"], (*w.ThisHouseDataValue)["PlayerBeting"])
	w.PerCapDep = cores.TransitionData(w.PerCapDep)
	w.PerCapBet = cores.TransitionData(w.PerCapBet)
	w.Content = fmt.Sprintf("2、人均存款金额%s，人均投注量%s。\n", w.PerCapDep, w.PerCapBet)
	cores.UpDataReport(w.Content, w.TxtPath)

	w.StrBetTotal = cores.TransitionData((*w.ThisHouseDataValue)["BetTotal"])
	w.StrWinOrlose = cores.TransitionWinOrLose((*w.ThisHouseDataValue)["WinOrLose"])
	w.Content = fmt.Sprintf("3、总投注%s，游戏%s。\n", w.StrBetTotal, w.StrWinOrlose)
	cores.UpDataReport(w.Content, w.TxtPath)

	w.StrDeposit = cores.TransitionData((*w.ThisHouseDataValue)["Deposit"])
	w.StrWithdrawal = cores.TransitionData((*w.ThisHouseDataValue)["Withdrawal"])
	w.DepSubWith = cores.DepositsSubWithdrawal((*w.ThisHouseDataValue)["Deposit"], (*w.ThisHouseDataValue)["Withdrawal"])
	w.DepSubWith = cores.TransitionData(w.DepSubWith)
	w.Link = cores.DepositsAndBetProportion((*w.ThisHouseDataValue)["Deposit"], (*w.ThisHouseDataValue)["BetTotal"])
	w.Content = fmt.Sprintf("4、总存款%s，总取款%s，存提差%s，充投比%s。\n", w.StrDeposit, w.StrWithdrawal, w.DepSubWith, w.Link)
	cores.UpDataReport(w.Content, w.TxtPath)

	(*w.ThisHouseDataValue)["Active"] = cores.TransitionData((*w.ThisHouseDataValue)["Active"])
	(*w.ThisHouseDataValue)["ReturnWater"] = cores.TransitionData((*w.ThisHouseDataValue)["ReturnWater"])
	w.Content = fmt.Sprintf("5、营销活动%s，返水%s。\n", (*w.ThisHouseDataValue)["Active"], (*w.ThisHouseDataValue)["ReturnWater"])
	cores.UpDataReport(w.Content, w.TxtPath)

	(*w.ThisHouseDataValue)["ProfitAndLoss"] = cores.TransitionWinOrLose((*w.ThisHouseDataValue)["ProfitAndLoss"])
	w.Content = fmt.Sprintf("6、总%s \n", (*w.ThisHouseDataValue)["ProfitAndLoss"])
	cores.UpDataReport(w.Content, w.TxtPath)

	PlayerAddLink := cores.ComparisonTool((*w.LastHouseDataValue)["PlayerAdd"], (*w.ThisHouseDataValue)["PlayerAdd"])
	PlayerActiveLink := cores.ComparisonTool((*w.LastHouseDataValue)["PlayerActive"], (*w.ThisHouseDataValue)["PlayerActive"])
	PlayerBetingLink := cores.ComparisonTool((*w.LastHouseDataValue)["PlayerBeting"], (*w.ThisHouseDataValue)["PlayerBeting"])
	DepositLink := cores.ComparisonTool((*w.LastHouseDataValue)["Deposit"], (*w.ThisHouseDataValue)["Deposit"])
	BetTotalLink := cores.ComparisonTool((*w.LastHouseDataValue)["BetTotal"], (*w.ThisHouseDataValue)["BetTotal"])
	w.Content = fmt.Sprintf(
		"7、本周相比上周数据，新增会员数%s,活跃人数%s,投注人数%s,存款数据%s,投注数据%s。\n",
		PlayerAddLink, PlayerActiveLink, PlayerBetingLink, DepositLink, BetTotalLink,
	)
	cores.UpDataReport(w.Content, w.TxtPath)
}
