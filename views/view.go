package views

import (
	"data/controllers"
	"fmt"
)

type User struct {
	Choose        int
	HouseAmount   int
	GameAmount    int
	CompanyAmount int
	ExcelPath     string
}

func MenuView() {

	fmt.Println("-------请选择您要做的事情------")
	fmt.Println("1、每日数据回报+每日游戏输赢")
	fmt.Println("2、周报数据")
	fmt.Println("3、月报数据")
	fmt.Println("4、 退出")
	fmt.Print("请输入：")
}

func (u *User) dailyData() {
	fmt.Println("-------正在进行初始化------")

	fmt.Println("在document文件中必须存在《下级代理帐变报表.xlsx》，文件名必须正确")

	fmt.Printf("请输入需要制作的子公司+房主数量：")
	fmt.Scanln(&u.HouseAmount)
	u.HouseAmount++

	fmt.Printf("请输入需要制作的游戏统计数量：")
	fmt.Scanln(&u.GameAmount)

	u.ExcelPath = "document/下级代理帐变报表.xlsx"
	daily := controllers.NewDailyDataController(u.HouseAmount, u.GameAmount, u.ExcelPath)

	daily.DailyDataUser()
}

func (u *User) weekData() {
	fmt.Println("-------正在进行初始化------")

	fmt.Println("在document文件中必须存在《周报数据表.xlsx》，文件名必须正确")

	fmt.Printf("请输入需要制作的公司层级数量：")
	fmt.Scanln(&u.CompanyAmount)
	u.CompanyAmount = 2 + u.CompanyAmount*4

	fmt.Printf("请输入需要制作的房间统计数量：")
	fmt.Scanln(&u.HouseAmount)
	u.HouseAmount = 30 + u.HouseAmount*4

	u.ExcelPath = "document/周报数据表.xlsx"
	week := controllers.NewWeekDataController(u.CompanyAmount, u.HouseAmount, u.ExcelPath)

	week.WeekDataUser()
}

func (u *User) ChooseMenu() {

	for {
		MenuView()
		fmt.Scanln(&u.Choose)
		switch u.Choose {
		case 1:
			u.dailyData()
			fmt.Println("已制作完成，请在document文件夹打开<dailyDataReport.txt>获取")
			return
		case 2:
			u.weekData()
			fmt.Println("已制作完成，请在document文件夹打开<weekDataReport.txt>获取")
			return
		case 3:
		case 4:
			fmt.Println("已退出，请关闭")
			return
		default:
			fmt.Println("输入错误请重新输入")
		}
	}
}
