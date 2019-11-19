package main

import (
	"chatroom/comm/message"
	"chatroom/server/process"
	"chatroom/server/utils"
	"fmt"
	"io"
	"net"
)

type Processr struct {
	Conn net.Conn
}

//服务器消息类型中控函数,根据用户发送消息进行处理
func (this *Processr) ServerProce(msg *message.Message) (err error) {
	//查看服务器收到信息没有
	switch msg.Type {
	case message.LonginType:
		user := &process.UserProcess{
			Conn: this.Conn,
		}
		err = user.ServerProceLongin(msg)
	case message.RegisterResMesType:
		user := &process.UserProcess{
			Conn: this.Conn,
		}
		err = user.ServerProceRegister(msg)

	case message.SmsMsgSendType:
		Send := &process.SmsProcess{}
		err = Send.SendGroupmsg(msg)
	default:

	}
	return

}

func (this *Processr) Process1() (err error) {

	for {
		//读数据封装一个函数 readpkg()
		tf := &utils.Tranfans{
			Conn: this.Conn,
		}
		msgs, err := tf.Readpkg()
		fmt.Println("客户端送消息", msgs)
		if err != nil {

			if err == io.EOF {
				fmt.Println("客户端已经退出", err)
				return err
			} else {
				fmt.Println("错误操作", err)
				return err
			}

		}
		err = this.ServerProce(&msgs)
		if err != nil {
			fmt.Println("错误操作", err)
			return err
		}
		// fmt.Println("内容", msgs)

	}

}
