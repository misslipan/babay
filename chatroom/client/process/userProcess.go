package process

import (
	"chatroom/client/utils"
	"chatroom/comm/message"
	_ "encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type UserProcess struct {
}

func (this *UserProcess) Login(userId int, password string) (err error) {
	// fmt.Printf("userid=%d  pass=%s", userId, password)

	// return nil
	//连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8889")

	if err != nil {
		fmt.Println("连接失败")
		return
	}
	defer conn.Close()
	//消息内容结构提
	var mes message.Message
	mes.Type = message.LonginType
	//登陆结构定义
	var logmes message.LoginMsg
	logmes.UserId = userId
	logmes.PassWord = password
	//序列化消息内容
	datas, err := json.Marshal(logmes)
	if err != nil {
		fmt.Println("错误操作", err)
		return
	}
	//将序列话的数据放进去
	mes.Data = string(datas)

	//将全部内容序列化
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("错误操作", err)
		return
	}

	Pk := &utils.Tranfans{Conn: conn}
	err = Pk.Writepkg(data)
	if err != nil {
		fmt.Println("发送内容失败", err)
		return
	}

	//获取服务端返回消息
	resmyg, err := Pk.Readpkg()

	if err != nil {
		fmt.Println("读取消息", err)
		return
	}

	var logres message.LoginReMmsg

	err = json.Unmarshal([]byte(resmyg.Data), &logres)
	if err != nil {
		fmt.Println("json.unmarshal  fail", err)
		return
	}

	if logres.Code == 200 {

		//初始化当前本地用户连接
		Curusers.Conn = conn
		Curusers.Userid = userId
		Curusers.UserStatus = message.Onliestatus

		fmt.Println("所有在线用户如下\t", err)

		for _, v := range logres.Usersids {

			if v == userId {
				continue
			}

			users := &message.User{
				Userid:     v,
				UserStatus: message.Onliestatus,
			}
			Useronlie[v] = users

			fmt.Println("在线用户\t", v)

		}

		fmt.Println("\n\n")
		go ProcessServer(conn)

		ShowDol()

	} else {
		fmt.Println(logres.ErrorMsg)

	}

	return
}

//注册请求
func (this *UserProcess) Register(userId int, password string, name string) (err error) {
	//连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8889")

	if err != nil {
		fmt.Println("连接失败")
		return
	}
	defer conn.Close()
	//消息内容结构提
	var mes message.Message
	mes.Type = message.RegisterResMesType
	//登陆结构定义
	var Regmes message.RegisterMes
	Regmes.User.Userid = userId
	Regmes.User.Password = password
	Regmes.User.Username = name
	//序列化消息内容
	datas, err := json.Marshal(Regmes.User)
	if err != nil {
		fmt.Println("错误操作", err)
		return
	}
	//将序列话的数据放进去
	mes.Data = string(datas)
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("错误操作", err)
		return
	}

	Pk := &utils.Tranfans{Conn: conn}
	err = Pk.Writepkg(data)
	if err != nil {
		fmt.Println("发送内容失败", err)
		return
	}
	//获取服务端返回消息
	resmyg, err := Pk.Readpkg()
	if err != nil {
		fmt.Println("读取消息", err)
		return
	}

	fmt.Println("客户端", resmyg)

	var Regmesa message.RegisterResMes

	err = json.Unmarshal([]byte(resmyg.Data), &Regmesa)
	if err != nil {
		fmt.Println("json.unmarshal  fail", err)
		return
	}

	if Regmesa.Code == 200 {
		fmt.Println(Regmesa.ErrorMsg)
		// os.Exit(0)

	} else {
		fmt.Println(Regmesa.ErrorMsg)
		os.Exit(0)
	}

	return
}
