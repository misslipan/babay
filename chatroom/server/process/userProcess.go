package process

import (
	"chatroom/comm/message"
	"chatroom/server/model"
	"chatroom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn   net.Conn
	Userid int
}

//用户状态检测通知其他用户上线
func (this *UserProcess) Notfiyuser(userid int) {

	for id, up := range UserMgrs.OnlieUsers {
		if id == userid {
			continue
		}
		// 将数据转发给各个用户
		up.NotifyMenoline(userid)
	}

}

func (this *UserProcess) NotifyMenoline(userid int) {

	var msga message.Message
	msga.Type = message.NotifystatusType

	var statusmsg message.Notifystatus
	statusmsg.Userid = userid
	statusmsg.Status = message.Onliestatus
	data, err := json.Marshal(statusmsg)
	if err != nil {
		fmt.Println("序列话操作失败")
		return
	}
	msga.Data = string(data)
	data, err = json.Marshal(msga)
	if err != nil {
		fmt.Println("序列话操作失败")
		return
	}

	//实例化机构体
	tf := &utils.Tranfans{
		Conn: this.Conn,
	}

	//发送消息
	err = tf.Writepkg(data)
	if err != nil {
		fmt.Println("数据发送失败", err)
		return
	}

}

//服务器验证登陆信息

func (this *UserProcess) ServerProceLongin(msg *message.Message) (err error) {

	//拿到登陆消息的结构体
	var longin message.LoginMsg
	//将客户端传过来的结构进行反序列化处理
	err = json.Unmarshal([]byte(msg.Data), &longin)
	if err != nil {
		fmt.Println("json.unmarshal  fail ", err)
		return
	}
	// fmt.Println(msg)

	//设置返回参数
	var longinResMsg message.LoginReMmsg
	//设置消息类型
	msg.Type = message.LoginReMmsgType

	// if longin.UserId == 100 && longin.PassWord == "123456" {
	// 	longinResMsg.Code = 200
	// 	longinResMsg.ErrorMsg = "登陆成功"
	// } else {
	// 	longinResMsg.Code = 500
	// 	longinResMsg.ErrorMsg = "该用户还未注册"
	// }
	//验证用户登陆信息是否正确
	user, err := model.MyuserDao.LoginCheck(longin.UserId, longin.PassWord)
	if err != nil {
		longinResMsg.Code = 500
		longinResMsg.ErrorMsg = "登陆失败"
	} else {

		longinResMsg.Code = 200

		this.Userid = user.Userid
		UserMgrs.AddonliUser(this)
		this.Notfiyuser(user.Userid)
		//将当前在线用户放入到userids
		for id, _ := range UserMgrs.OnlieUsers {
			longinResMsg.Usersids = append(longinResMsg.Usersids, id)
		}

		longinResMsg.ErrorMsg = "登陆成功"
		fmt.Println(longinResMsg)

	}

	data, err := json.Marshal(longinResMsg)
	if err != nil {
		fmt.Println("json.marshal fail ", err)
		return
	}
	msg.Data = string(data)

	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("json.marshal fail ", err)
		return
	}

	//实例化机构体
	tf := &utils.Tranfans{
		Conn: this.Conn,
	}

	//发送消息
	err = tf.Writepkg(data)
	return
}

//注册用户
func (this *UserProcess) ServerProceRegister(msg *message.Message) (err error) {
	//拿到注册消息的结构体
	var Regmes message.RegisterMes
	//将客户端传过来的结构进行反序列化处理
	err = json.Unmarshal([]byte(msg.Data), &Regmes.User)

	if err != nil {
		fmt.Println("json.unmarshal  fail ", err)
		return
	}
	//设置返回参数
	var ResitserResMsg message.RegisterResMes
	//设置消息类型
	msg.Type = message.RegisterErrType

	err = model.MyuserDao.Register(&Regmes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			ResitserResMsg.Code = 505
			ResitserResMsg.ErrorMsg = err.Error()
		} else {
			ResitserResMsg.Code = 506
			ResitserResMsg.ErrorMsg = "发生未知错误"
		}

	} else {

		ResitserResMsg.Code = 200
		ResitserResMsg.ErrorMsg = "注册成功"

	}

	data, err := json.Marshal(ResitserResMsg)
	if err != nil {
		fmt.Println("json.marshal fail ", err)
		return
	}
	msg.Data = string(data)

	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("json.marshal fail ", err)
		return
	}
	//实例化机构体
	tf := &utils.Tranfans{
		Conn: this.Conn,
	}

	//发送消息
	err = tf.Writepkg(data)
	return

}
