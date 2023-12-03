package service

import (
	"data/cores"
	"data/models"
	"fmt"
	"strconv"
)

type ServiceDailyData struct {
	DailyData     *models.DailyData
	GameData      *models.GameData
	ActiveAounmt  string //优惠返佣返水总和
	WinOrLose     string //游戏输赢
	ProfitAndLoss string //总盈亏
	Content       string //输出内容
	TxtPath       string //写入的路径
	ExcelPath     string //读取Excel路径
}

// 返水+活动营销+代理返佣
func (s *ServiceDailyData) GetActiveAounmt() {
	rwFloat, err := strconv.ParseFloat(s.DailyData.ReturnWater, 64)
	if err != nil {
		fmt.Println("返水转化Float64类型错误", err)
	}

	ActiveFloat, err2 := strconv.ParseFloat(s.DailyData.Active, 64)
	if err2 != nil {
		fmt.Println("营销活动转化Float64类型错误", err2)
	}

	ComFloat, err3 := strconv.ParseFloat(s.DailyData.Commission, 64)
	if err3 != nil {
		fmt.Println("代理返佣转化Float64类型错误", err3)
	}

	ActiveAounmt := rwFloat + ActiveFloat + ComFloat
	s.ActiveAounmt = strconv.FormatFloat(ActiveAounmt, 'f', -1, 64)
}

// 有效-派彩 = 游戏输赢
func (s *ServiceDailyData) GetWinOrLose() {
	betFloat, err := strconv.ParseFloat(s.DailyData.BetTotal, 64)
	if err != nil {
		fmt.Println("有效投注转化Float64类型错误", err)
	}

	winFloat, err := strconv.ParseFloat(s.DailyData.WinAmount, 64)
	if err != nil {
		fmt.Println("派彩金额转化Float64类型错误", err)
	}

	WinOrLose := betFloat - winFloat
	s.WinOrLose = strconv.FormatFloat(WinOrLose, 'f', -1, 64)
}

// 子公司、房主总输赢   游戏输赢-总营销
func (s *ServiceDailyData) GetProfitAndLoss() {
	WinOrLoseFloat, err := strconv.ParseFloat(s.WinOrLose, 64)
	if err != nil {
		fmt.Println("游戏输赢转化Float64类型错误", err)
	}

	ActiveAounmtFloat, err := strconv.ParseFloat(s.ActiveAounmt, 64)
	if err != nil {
		fmt.Println("营销总支出转化Float64类型错误", err)
	}

	ProfitAndLoss := WinOrLoseFloat - ActiveAounmtFloat
	s.ProfitAndLoss = strconv.FormatFloat(ProfitAndLoss, 'f', -1, 64)
}

func (s *ServiceDailyData) FormatDailyDataContent() {
	Deposit := cores.TransitionData(s.DailyData.Deposit)
	Withdrawal := cores.TransitionData(s.DailyData.Withdrawal)
	DepSubWith := cores.TransitionData(s.DailyData.DepSubWith)
	BetTotal := cores.TransitionData(s.DailyData.BetTotal)
	WinOrLose := cores.TransitionWinOrLose(s.WinOrLose)
	ActiveAounmt := cores.TransitionData(s.ActiveAounmt)
	ProfitAndLoss := cores.IsCompanyWinOrLose(s.ProfitAndLoss)
	s.Content = fmt.Sprintf("%s:存款%s，取款%s，存取差%s，有效投注%s，游戏%s，优惠返佣返水总和%s，%s \n",
		s.DailyData.HouseName, Deposit, Withdrawal, DepSubWith, BetTotal, WinOrLose, ActiveAounmt, ProfitAndLoss,
	)

	cores.UpDataReport(s.Content, s.TxtPath)
}

func (s *ServiceDailyData) FormatGameContent(GameName string, BetTotal string, WinOrLose string) {
	BetTotal = cores.TransitionData(BetTotal)
	WinOrLose = cores.TransitionWinOrLose(WinOrLose)
	s.Content = fmt.Sprintf("%s:有效投注%s,%s \n", GameName, BetTotal, WinOrLose)

	cores.UpDataReport(s.Content, s.TxtPath)
}
