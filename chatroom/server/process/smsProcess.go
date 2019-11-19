package process

import (
	"chatroom/client/utils"
	"chatroom/comm/message"
	"encoding/json"
	"fmt"
)

type SmsProcess struct {
}

func (this *SmsProcess) SendGroupmsg(msg *message.Message) (err error) {

	//遍历服务端的onileusers  结构体数组
	//将消息转发出去
	var Sendmsg message.SmsMsgSend
	err = json.Unmarshal([]byte(msg.Data), &Sendmsg)
	if err != nil {
		fmt.Println("json 转移失败")
		return
	}
	data, err := json.Marshal(&msg)
	if err != nil {
		fmt.Println("json 转移失败")
		return
	}
	for id, up := range UserMgrs.OnlieUsers {
		if id == Sendmsg.Userid {
			continue
		}

		//实例化消息连接
		tf := utils.Tranfans{
			Conn: up.Conn,
		}
		//向所有客户端发送消息
		err = tf.Writepkg(data)
		if err != nil {
			fmt.Println("转发消息失败", err)
		}
	}
	return

}

/**
用户一对一发送消息
*/
func (this *SmsProcess) SendOnemsg(msg *message.Message) (err error) {

	//遍历服务端的onileusers  结构体数组
	//将消息转发出去
	var Sendmsg message.SmsMsgSend
	err = json.Unmarshal([]byte(msg.Data), &Sendmsg)
	if err != nil {
		fmt.Println("json 转移失败")
		return
	}
	data, err := json.Marshal(&msg)
	if err != nil {
		fmt.Println("json 转移失败")
		return
	}
	up, err := UserMgrs.GetonileuserById(Sendmsg.Friendid)
	if err != nil {
		fmt.Println(err)
		return
	}

	//实例化消息连接
	tf := utils.Tranfans{
		Conn: up.Conn,
	}
	//向所有客户端发送消息
	err = tf.Writepkg(data)
	if err != nil {
		fmt.Println("转发消息失败", err)
	}

	return

}
