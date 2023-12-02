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
	Elemnet       *models.WeekDataElement
	TableName     string
	HouseAmount   int
	CompanyAmount int
	F             *excelize.File
}

func (w *WeekDataController) WeekDataUser() {
	for i := 6; i <= w.CompanyAmount; i++ {
		w.Elemnet.HouseName = fmt.Sprintf("A%v", i)
		w.WeekService.WeekData.HouseName = cores.GetExcelValue(w.F, w.TableName, w.Elemnet.HouseName)

		w.Elemnet.HouseCard = fmt.Sprintf("C%v", i)
		w.WeekService.WeekData.HouseCard = cores.GetExcelValue(w.F, w.TableName, w.Elemnet.HouseCard)

		w.Elemnet.HouseAdd = fmt.Sprintf("D%v", i)
		w.WeekService.WeekData.HouseAdd = cores.GetExcelValue(w.F, w.TableName, w.Elemnet.HouseAdd)

		w.Elemnet.HouseExpire = fmt.Sprintf("E%v", i)
		w.WeekService.WeekData.HouseExpire = cores.GetExcelValue(w.F, w.TableName, w.Elemnet.HouseExpire)

		w.Elemnet.PlayerAdd = fmt.Sprintf("F%v", i)
		w.WeekService.WeekData.PlayerAdd = cores.GetExcelValue(w.F, w.TableName, w.Elemnet.PlayerAdd)

		w.Elemnet.PlayerActive = fmt.Sprintf("H%v", i)
		w.WeekService.WeekData.PlayerActive = cores.GetExcelValue(w.F, w.TableName, w.Elemnet.PlayerActive)

		w.Elemnet.PlayerBeting = fmt.Sprintf("I%v", i)
		w.WeekService.WeekData.PlayerBeting = cores.GetExcelValue(w.F, w.TableName, w.Elemnet.PlayerBeting)

		w.Elemnet.PlayerDeposit = fmt.Sprintf("J%v", i)
		w.WeekService.WeekData.PlayerDeposit = cores.GetExcelValue(w.F, w.TableName, w.Elemnet.PlayerDeposit)

		w.Elemnet.DepositAmount = fmt.Sprintf("K%v", i)
		w.WeekService.WeekData.DepositAmount = cores.GetExcelValue(w.F, w.TableName, w.Elemnet.DepositAmount)

		w.Elemnet.Deposit = fmt.Sprintf("L%v", i)
		w.WeekService.WeekData.Deposit = cores.GetExcelValue(w.F, w.TableName, w.Elemnet.Deposit)

		w.Elemnet.PlayerWithdrawal = fmt.Sprintf("M%v", i)
		w.WeekService.WeekData.PlayerWithdrawal = cores.GetExcelValue(w.F, w.TableName, w.Elemnet.PlayerWithdrawal)

		w.Elemnet.WithdrawalAmount = fmt.Sprintf("N%v", i)
		w.WeekService.WeekData.WithdrawalAmount = cores.GetExcelValue(w.F, w.TableName, w.Elemnet.WithdrawalAmount)

		w.Elemnet.Withdrawal = fmt.Sprintf("O%v", i)
		w.WeekService.WeekData.Withdrawal = cores.GetExcelValue(w.F, w.TableName, w.Elemnet.Withdrawal)

		w.Elemnet.BetTotal = fmt.Sprintf("R%v", i)
		w.WeekService.WeekData.BetTotal = cores.GetExcelValue(w.F, w.TableName, w.Elemnet.BetTotal)

		w.Elemnet.WinOrLose = fmt.Sprintf("S%v", i)
		w.WeekService.WeekData.WinOrLose = cores.GetExcelValue(w.F, w.TableName, w.Elemnet.WinOrLose)

		w.Elemnet.Active = fmt.Sprintf("T%v", i)
		w.WeekService.WeekData.Active = cores.GetExcelValue(w.F, w.TableName, w.Elemnet.Active)

		w.Elemnet.ReturnWater = fmt.Sprintf("U%v", i)
		w.WeekService.WeekData.ReturnWater = cores.GetExcelValue(w.F, w.TableName, w.Elemnet.ReturnWater)

		w.Elemnet.Commission = fmt.Sprintf("V%v", i)
		w.WeekService.WeekData.Commission = cores.GetExcelValue(w.F, w.TableName, w.Elemnet.Commission)

		w.Elemnet.PlayerActiveDaily = fmt.Sprintf("X%v", i)
		w.WeekService.WeekData.PlayerActiveDaily = cores.GetExcelValue(w.F, w.TableName, w.Elemnet.PlayerActiveDaily)

		w.Elemnet.PlayerBetingDaily = fmt.Sprintf("Y%v", i)
		w.WeekService.WeekData.PlayerBetingDaily = cores.GetExcelValue(w.F, w.TableName, w.Elemnet.PlayerBetingDaily)
	}
}

type WeekData struct {
	HouseName         string //房间名字
	HouseCard         string //有效房卡
	HouseAdd          string //新增房卡
	HouseExpire       string //到期房卡
	PlayerAdd         string //新增会员
	PlayerActive      string //活跃人数
	PlayerBeting      string //投注人数
	PlayerDeposit     string //存款人数
	DepositAmount     string //存款次数
	Deposit           string //存款金额
	PlayerWithdrawal  string //取款人数
	WithdrawalAmount  string //取款次数
	Withdrawal        string //取款金额
	FirstDeposit      string //首存
	BetTotal          string //有效投注
	WinOrLose         string //游戏输赢
	Active            string //营销活动
	ReturnWater       string //返水
	Commission        string //代理返佣
	PlayerActiveDaily string //日均活跃人数
	PlayerBetingDaily string //日均投注人数

}

type WeekDataElement struct {
	HouseName         string //房间名字
	HouseCard         string //有效房卡
	HouseAdd          string //新增房卡
	HouseExpire       string //到期房卡
	PlayerAdd         string //新增会员
	PlayerActive      string //活跃人数
	PlayerBeting      string //投注人数
	PlayerDeposit     string //存款人数
	DepositAmount     string //存款次数
	Deposit           string //存款金额
	PlayerWithdrawal  string //取款人数
	WithdrawalAmount  string //取款次数
	Withdrawal        string //取款金额
	FirstDeposit      string //首存
	BetTotal          string //有效投注
	WinOrLose         string //游戏输赢
	Active            string //营销活动
	ReturnWater       string //返水
	Commission        string //代理返佣
	PlayerActiveDaily string //日均活跃人数
	PlayerBetingDaily string //日均投注人数

}
