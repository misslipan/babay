package process

import (
	"chatroom/client/utils"
	"chatroom/comm/message"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

//显示登陆完成菜单
func ShowDol() {
	var key int
	var mysend string
	for {
		fmt.Printf("------------你好%d登陆成功----------------------\n\n", Curusers.Userid)
		fmt.Println("\t\t\t 1.显示在线列表")
		fmt.Println("\t\t\t 2.发送消息")
		fmt.Println("\t\t\t 3.消息列表")
		fmt.Println("\t\t\t 4.退出登陆")
		fmt.Println("\t\t\t 请选择(1-3)")

		//因为我们一直会用到这个实例 Smsprocess
		fmt.Scanf("%d\t", &key)

		smsProa := &SmsProcess{}
		switch key {
		case 1:
			// fmt.Println("在线列表")
			outputonline()

		case 2:
			fmt.Println("发送消息\t")
			fmt.Scanf("%s\t", &mysend)
			err := smsProa.SendGroupMes(mysend)
			fmt.Println(err)

		case 3:
			fmt.Println("消息列表")
		case 4:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("在线列表")

		}
	}
}

func ProcessServer(conn net.Conn) {

	tf := &utils.Tranfans{
		Conn: conn,
	}

	for {

		fmt.Println(tf.Conn.LocalAddr().String())
		fmt.Println("客户端等待服务器发送消息")
		myge, err := tf.Readpkg()
		if err != nil {

			fmt.Println("tf.readpkg fail", err)
			return

		}
		switch myge.Type {
		case message.NotifystatusType:
			var notfiyonline *message.Notifystatus
			err := json.Unmarshal([]byte(myge.Data), &notfiyonline)
			if err != nil {
				fmt.Println("json.Unmarshal fila", err)
				return
			}
			UpdateUserStatus(notfiyonline)
		case message.SmsMsgSendType:
			fmt.Println(myge)

		default:
			fmt.Println("没有获取到服务器的操作指令")

		}

		fmt.Println("msg=", myge)
	}
}
