package cores

import (
	"fmt"
	"strconv"
	"strings"
)

// 环比计算
func ComparisonTool(lastWeekData string, thisWeekData string) string {
	lastWeekDataFloat, _ := strconv.ParseFloat(lastWeekData, 64)
	thisWeekDataFloat, _ := strconv.ParseFloat(thisWeekData, 64)

	if lastWeekDataFloat == thisWeekDataFloat {
		return "持平"
	}
	if thisWeekDataFloat > lastWeekDataFloat {
		res := thisWeekDataFloat/lastWeekDataFloat - 1
		res *= 100
		resStr := strconv.FormatFloat(res, 'f', 2, 64)
		newStr := fmt.Sprintf("成长%s%%", resStr)
		str := strings.Replace(newStr, "-", "", -1)
		return str
	} else {
		res := 1 - thisWeekDataFloat/lastWeekDataFloat
		res *= 100
		resStr := strconv.FormatFloat(res, 'f', 2, 64)
		newStr := fmt.Sprintf("下降%s%%", resStr)
		return newStr
	}
}

// 人均计算
func PerCapitaStatistics(data string, people string) string {
	dataFloat, _ := strconv.ParseFloat(data, 64)
	peopleFloat, _ := strconv.ParseFloat(people, 64)
	res := dataFloat / peopleFloat
	resStr := fmt.Sprintf("%f", res)
	return resStr
}

// 充投比计算，保留2位
func DepositsAndBetProportion(Deposits string, Bet string) string {
	DepositsFloat, _ := strconv.ParseFloat(Deposits, 64)
	BetFloat, _ := strconv.ParseFloat(Bet, 64)
	res := BetFloat / DepositsFloat
	strRes := strconv.FormatFloat(res, 'f', 2, 64)
	return strRes
}

// 存提差计算
func DepositsSubWithdrawal(Deposits string, Withdrawal string) string {
	DepositsFloat, _ := strconv.ParseFloat(Deposits, 64)
	WithdrawalFloat, _ := strconv.ParseFloat(Withdrawal, 64)
	res := DepositsFloat - WithdrawalFloat
	strRes := fmt.Sprintf("%f", res)
	return strRes
}
