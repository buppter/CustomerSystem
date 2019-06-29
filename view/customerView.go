package main

import (
	"fmt"
	"go_code/customer_system/model"
	"go_code/customer_system/service"
)

type customerView struct {
	//定义必要的字段

	key  string //接收用户输入
	loop bool   //表示是否循环显示主菜单

	//增加一个字段customerService
	customerService *service.CustomerService
}

//显示主菜单
func (cv *customerView) mainMenu() {
	for {
		fmt.Println("******************客户信息管理软件******************")
		fmt.Println("				  1. 添 加 客 户")
		fmt.Println("				  2. 修 改 客 户")
		fmt.Println("				  3. 删 除 客 户")
		fmt.Println("				  4. 客 户 列 表")
		fmt.Println("				  5. 退	    出")
		fmt.Println("请选择(1-5):")

		_, err := fmt.Scanln(&cv.key)
		if err != nil {
			fmt.Println(err)
			continue
		}

		switch cv.key {
		case "1":
			cv.add()
		case "2":
			fmt.Println("修 改 客 户")
		case "3":
			cv.delete()

		case "4":
			cv.list()
		case "5":
			cv.loop = false
		default:
			fmt.Println("你的输入有误，请重新输入...")
		}

		if !cv.loop {
			break
		}
	}

	fmt.Println("你退出了客户关系管理系统")
}

//显示所有客户信息
func (cv *customerView) list() {
	//首先，获取的当前所有的客户信息（切片中）
	customers := cv.customerService.List()

	//显示
	fmt.Println("***************************客户列表**************************")
	fmt.Println("编号\t\t姓名\t\t性别\t\t年龄\t\t电话\t\t\t\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Printf("*************************客户列表完成*************************\n\n\n\n")
}

//得到用户输入，信息构建新的客户，并完成添加
func (cv *customerView) add() {
	fmt.Println("***************************添加客户**************************")
	fmt.Println("姓名：")
	name := ""
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("性别：")
	gender := ""
	_, err = fmt.Scanln(&gender)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("年龄：")
	age := 0
	_, err = fmt.Scanln(&age)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("电话：")
	phone := ""
	_, err = fmt.Scanln(&phone)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("邮箱：")
	email := ""
	_, err = fmt.Scanln(&email)
	if err != nil {
		fmt.Println(err)
	}

	//构建一个新的Customer实例
	//注意：id号，没有让用户输入，ID是唯一的，需要系统分配
	customer := model.NewCustomer2(name, gender, age, phone, email)
	if cv.customerService.Add(customer) {
		fmt.Println("***************************添加成功**************************")
	} else {
		fmt.Println("***************************添加失败**************************")
	}
}

//得到用户输入的ID，删除该ID对应的客户
func (cv *customerView) delete() {
	fmt.Println("***************************删除客户**************************")
	fmt.Println("请选择要删除的客户编号(输入-1退出):")
	id := -1

	_, err := fmt.Scanln(&id)
	if err != nil {
		fmt.Println(err)
	} else if id == -1 {
		return
	}

	fmt.Println("确认删除(Y/N):")
	choice := ""
	_, err = fmt.Scanln(&choice)
	if choice == "Y" || choice == "y" {
		//调用customerService的Delete方法
		if cv.customerService.Delete(id) {
			fmt.Println("***************************删除成功**************************")
		} else {
			fmt.Println("**********************删除失败，该ID号不存在********************")
		}
	}

}
func main() {
	//在main函数中，创建一个customerView实例，并运行主菜单
	customerView := customerView{
		key:  "",
		loop: true,
	}
	//完成对customerView结构体的customerService字段的初始化
	customerView.customerService = service.NewCustomerService()
	//显示主菜单
	customerView.mainMenu()
}
