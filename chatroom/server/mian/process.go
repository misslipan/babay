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
		//多人发送消息
	case message.SmsMsgSendType:
		Send := &process.SmsProcess{}
		err = Send.SendGroupmsg(msg)
	case message.SmsMsgOneType:
		Send := &process.SmsProcess{}
		err = Send.SendOnemsg(msg)
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
