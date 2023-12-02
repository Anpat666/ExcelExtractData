package models

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
	ProfitAndLoss     string //总盈亏
	PlayerActiveDaily string //日均活跃人数
	PlayerBetingDaily string //日均投注人数

}

var ThisWeekDataMap = map[string]string{
	"HouseName":         "A",
	"HouseCard":         "C",
	"HouseAdd":          "D",
	"HouseExpire":       "E",
	"PlayerAdd":         "F",
	"PlayerActive":      "H",
	"PlayerBeting":      "I",
	"PlayerDeposit":     "J",
	"DepositAmount":     "K",
	"Deposit":           "L",
	"PlayerWithdrawal":  "M",
	"WithdrawalAmount":  "N",
	"Withdrawal":        "O",
	"FirstDeposit":      "P",
	"BetTotal":          "R",
	"WinOrLose":         "S",
	"Active":            "T",
	"ReturnWater":       "U",
	"Commission":        "V",
	"ProfitAndLoss":     "W",
	"PlayerActiveDaily": "X",
	"PlayerBetingDaily": "Y",
}

var ThisWeekDataMapValue = map[string]string{
	"HouseName":         "A",
	"HouseCard":         "C",
	"HouseAdd":          "D",
	"HouseExpire":       "E",
	"PlayerAdd":         "F",
	"PlayerActive":      "H",
	"PlayerBeting":      "I",
	"PlayerDeposit":     "J",
	"DepositAmount":     "K",
	"Deposit":           "L",
	"PlayerWithdrawal":  "M",
	"WithdrawalAmount":  "N",
	"Withdrawal":        "O",
	"FirstDeposit":      "P",
	"BetTotal":          "R",
	"WinOrLose":         "S",
	"Active":            "T",
	"ReturnWater":       "U",
	"Commission":        "V",
	"ProfitAndLoss":     "W",
	"PlayerActiveDaily": "X",
	"PlayerBetingDaily": "Y",
}

var LastWeekDataMapValue = map[string]string{
	"HouseName":         "A",
	"HouseCard":         "C",
	"HouseAdd":          "D",
	"HouseExpire":       "E",
	"PlayerAdd":         "F",
	"PlayerActive":      "H",
	"PlayerBeting":      "I",
	"PlayerDeposit":     "J",
	"DepositAmount":     "K",
	"Deposit":           "L",
	"PlayerWithdrawal":  "M",
	"WithdrawalAmount":  "N",
	"Withdrawal":        "O",
	"FirstDeposit":      "P",
	"BetTotal":          "R",
	"WinOrLose":         "S",
	"Active":            "T",
	"ReturnWater":       "U",
	"Commission":        "V",
	"ProfitAndLoss":     "W",
	"PlayerActiveDaily": "X",
	"PlayerBetingDaily": "Y",
}
