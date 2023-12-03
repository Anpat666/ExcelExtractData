package models

type GameWeek struct {
	ThisGame [][]string
	LastGame [][]string
}

var ThisWeekDataMap = map[string]string{
	"HouseName":         "A", //房间名字
	"HouseCard":         "C", //有效房卡
	"HouseAdd":          "D", //新增房卡
	"HouseExpire":       "E", //到期房卡
	"PlayerAdd":         "F", //新增会员
	"PlayerActive":      "H", //活跃人数
	"PlayerBeting":      "I", //投注人数
	"PlayerDeposit":     "J", //存款人数
	"DepositAmount":     "K", //存款次数
	"Deposit":           "L", //存款金额
	"PlayerWithdrawal":  "M", //取款人数
	"WithdrawalAmount":  "N", //取款次数
	"Withdrawal":        "O", //取款金额
	"FirstDeposit":      "P", //首存
	"BetTotal":          "R", //有效投注
	"WinOrLose":         "S", //游戏输赢
	"Active":            "T", //营销活动
	"ReturnWater":       "U", //返水
	"Commission":        "V", //代理返佣
	"ProfitAndLoss":     "W", //总盈亏
	"PlayerActiveDaily": "X", //日均活跃人数
	"PlayerBetingDaily": "Y", //日均投注人数
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

var ThisHouseData = map[string]string{
	"HouseName":          "A", //房间名字
	"PlayerAdd":          "C", //新增会员
	"PlayerActive":       "D", //活跃人数
	"PlayerBeting":       "E", //投注人数
	"PlayerDeposit":      "F", //存款人数
	"DepositAmount":      "G", //存款次数
	"Deposit":            "H", //存款金额
	"PlayerWithdrawal":   "I", //取款人数
	"WithdrawalAmount":   "J", //取款次数
	"Withdrawal":         "K", //取款金额
	"PlayerFirstDeposit": "L", //首存人数
	"BetTotal":           "M", //有效投注
	"WinOrLose":          "N", //游戏输赢
	"Active":             "O", //营销活动
	"ReturnWater":        "P", //返水
	"Commission":         "Q", //代理返佣
	"ProfitAndLoss":      "R", //总盈亏
}

var ThisHouseDataValue = map[string]string{
	"HouseName":          "A",
	"PlayerAdd":          "C",
	"PlayerActive":       "D",
	"PlayerBeting":       "E",
	"PlayerDeposit":      "F",
	"DepositAmount":      "G",
	"Deposit":            "H",
	"PlayerWithdrawal":   "I",
	"WithdrawalAmount":   "J",
	"Withdrawal":         "K",
	"PlayerFirstDeposit": "L",
	"BetTotal":           "M",
	"WinOrLose":          "N",
	"Active":             "O",
	"ReturnWater":        "P",
	"Commission":         "Q",
	"ProfitAndLoss":      "R",
}

var LastHouseDataValue = map[string]string{
	"HouseName":          "A",
	"PlayerAdd":          "C",
	"PlayerActive":       "D",
	"PlayerBeting":       "E",
	"PlayerDeposit":      "F",
	"DepositAmount":      "G",
	"Deposit":            "H",
	"PlayerWithdrawal":   "I",
	"WithdrawalAmount":   "J",
	"Withdrawal":         "K",
	"PlayerFirstDeposit": "L",
	"BetTotal":           "M",
	"WinOrLose":          "N",
	"Active":             "O",
	"ReturnWater":        "P",
	"Commission":         "Q",
	"ProfitAndLoss":      "R",
}
