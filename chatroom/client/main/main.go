package main

import (
	"chatroom/client/process"
	"fmt"
	"os"
)

var (
	userId   int
	password string
	name     string
)

func main() {

	//输入指令
	var key int

	for true {
		fmt.Println("------------------欢迎来到聊天室----------------------")
		fmt.Println("\t\t\t 1,登陆聊天室")
		fmt.Println("\t\t\t 2.注册聊天室")
		fmt.Println("\t\t\t 3.退出系统")
		fmt.Println("\t\t\t 请选择(1-3)")

		fmt.Scanln(&key)

		switch key {
		case 1:
			fmt.Println("请输入账号")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("输入密码")
			fmt.Scanf("%s\n", &password)
			users := &process.UserProcess{}
			err := users.Login(userId, password)
			fmt.Println(err)

		case 2:
			fmt.Println("注册聊天")
			fmt.Println("请输入账号")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("输入密码")
			fmt.Scanf("%s\n", &password)
			fmt.Println("用户呢称")
			fmt.Scanf("%s\n", &name)
			//调用
			users := &process.UserProcess{}
			users.Register(userId, password, name)
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("输入有误,请重新输入")

		}
	}

}
