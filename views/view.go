package views

import (
	"data/cores"
	"data/services"
	"fmt"
)

type UserView struct {
	Choose        int
	HouseAmount   int
	GameChoose    int
	CompanyAmount int
	ExcelPath     string
	TxtPath       string
	ThisTableName string
	LastTableName string
}

func MenuView() {

	fmt.Println("-------请选择您要做的事情------")
	fmt.Println("1、每日数据回报+每日游戏输赢")
	fmt.Println("2、周报数据")
	fmt.Println("3、月报数据")
	fmt.Println("4、 退出")
	fmt.Print("请输入：")
}

func (u *UserView) dailyDataView() {
	fmt.Println("-------正在进行初始化------")
	fmt.Println("在document文件中必须存在《下级代理帐变报表.xlsx》，文件名必须正确")
	fmt.Printf("请输入需要制作的子公司+房主数量：")
	fmt.Scanln(&u.HouseAmount)
	u.HouseAmount++

	u.ExcelPath = "document/下级代理帐变报表.xlsx"
	u.TxtPath = "document/dailyDataReport.txt"
	daily := services.NewDailyDataService(u.HouseAmount, u.ExcelPath, u.TxtPath)

	daily.DailyDataUser()
	cores.ReplaceDocument("分公司", "房主", u.TxtPath, 4)

	for {
		fmt.Printf("是否需要制作游戏输赢统计(输入1需要或者输入2不需要)：")
		fmt.Scanln(&u.GameChoose)
		if u.GameChoose == 1 || u.GameChoose == 2 {
			if u.GameChoose == 1 {
				daily.GamesDailyData()
				break
			} else {
				break
			}
		} else {
			fmt.Println("选择有误，请重新输入")
		}
	}

}

func (u *UserView) weekDataView() {
	fmt.Println("-------正在进行初始化------")

	fmt.Println("在document文件中必须存在《周报数据表.xlsx》，文件名必须正确")

	fmt.Printf("请输入需要制作的房间统计数量：")
	fmt.Scanln(&u.HouseAmount)
	u.HouseAmount = 30 + u.HouseAmount*4

	u.CompanyAmount = 18
	u.ExcelPath = "document/周报数据表.xlsx"
	u.TxtPath = "document/weekDataReport.txt"

	week := services.NewWeekDataService(u.CompanyAmount, u.HouseAmount, u.ExcelPath, u.TxtPath)

	week.WeekDataUser()
	for {
		fmt.Printf("是否需要制作周游戏数据回报(输入1需要或者输入2不需要)：")
		fmt.Scanln(&u.GameChoose)
		if u.GameChoose == 1 || u.GameChoose == 2 {
			if u.GameChoose == 1 {
				week.GameWeekData()
				break
			} else {
				break
			}
		} else {
			fmt.Println("选择有误，请重新输入")
		}
	}

}

func (u *UserView) monDataView() {
	fmt.Println("-------正在进行初始化------")

	fmt.Println("在document文件中必须存在《月报数据表.xlsx》，文件名必须正确")

	fmt.Printf("请输入需要制作的房间统计数量：")
	fmt.Scanln(&u.HouseAmount)
	u.HouseAmount = 30 + u.HouseAmount*4

	u.CompanyAmount = 18
	u.ExcelPath = "document/月报数据表.xlsx"
	u.TxtPath = "document/monDataReport.txt"
	u.ThisTableName = "本月游戏数据"
	u.LastTableName = "上月游戏数据"
	week := services.NewWeekDataService(u.CompanyAmount, u.HouseAmount, u.ExcelPath, u.TxtPath)

	week.WeekDataUser()
	for {
		fmt.Printf("是否需要制作月游戏数据回报(输入1需要或者输入2不需要)：")
		fmt.Scanln(&u.GameChoose)
		if u.GameChoose == 1 || u.GameChoose == 2 {
			if u.GameChoose == 1 {
				week.GameWeekData()
				break
			} else {
				break
			}
		} else {
			fmt.Println("选择有误，请重新输入")
		}
	}

	cores.ReplaceDocument("周", "月", u.TxtPath, 0)
}

func (u *UserView) ChooseMenu() {

	for {
		MenuView()
		fmt.Scanln(&u.Choose)
		switch u.Choose {
		case 1:
			u.dailyDataView()
			fmt.Println("已制作完成，请在document文件夹打开<dailyDataReport.txt>获取")
			return
		case 2:
			u.weekDataView()
			fmt.Println("已制作完成，请在document文件夹打开<weekDataReport.txt>获取")
			return
		case 3:
			u.monDataView()
			fmt.Println("已制作完成，请在document文件夹打开<monDataReport.txt>获取")
			return
		case 4:
			fmt.Println("已退出，请关闭")
			return
		default:
			fmt.Println("输入错误请重新输入")
		}
	}
}
