package models

type DailyData struct {
	HouseName   string //房间名字
	Deposit     string //存款金额
	Withdrawal  string //取款金额
	DepSubWith  string //存取差
	BetTotal    string //有效投注
	WinAmount   string //派彩金额
	ReturnWater string //返水
	Active      string //营销活动
	Commission  string //返佣

}

type DailyDataElement struct {
	HouseName   string //房间名字
	Deposit     string //存款金额
	Withdrawal  string //取款金额
	DepSubWith  string //存取差
	BetTotal    string //有效投注
	WinAmount   string //派彩金额
	ReturnWater string //返水
	Active      string //营销活动
	Commission  string //返佣
}

type GameData struct {
	G7Rows [][]string
	YYRows [][]string
	BYRows [][]string
	ZSRows [][]string
}
