package process

import (
	"chatroom/client/model"
	"chatroom/comm/message"
	"fmt"
)

var Useronlie map[int]*message.User = make(map[int]*message.User, 10)

var Curusers model.CurUser

//显示在线用户
func outputonline() {
	fmt.Println("当前在线列表:")
	for id, _ := range Useronlie {
		fmt.Println("用户id:\t", id)
	}
}

//修改用户状态
func UpdateUserStatus(notifyMes *message.Notifystatus) {

	user, ok := Useronlie[notifyMes.Userid]
	if !ok {
		user = &message.User{
			Userid: notifyMes.Userid,
		}

	}
	user.UserStatus = notifyMes.Status
	Useronlie[notifyMes.Userid] = user
	outputonline()
}
