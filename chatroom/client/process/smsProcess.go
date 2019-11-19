package process

import (
	"chatroom/client/utils"
	"chatroom/comm/message"
	"encoding/json"
	"fmt"
)

type SmsProcess struct {
}

/**

发送群聊消息
*/

func (this *SmsProcess) SendGroupMes(conetent string) (err error) {

	var mes message.Message
	mes.Type = message.SmsMsgSendType
	var mesdata message.SmsMsgSend
	//将消息内容放入结构体
	mesdata.Content = conetent
	mesdata.Userid = Curusers.Userid
	mesdata.UserStatus = Curusers.UserStatus

	//序列花mesdata
	data, err := json.Marshal(mesdata)
	if err != nil {
		fmt.Println("sendgroup json.marshal fila", err.Error())
		return
	}
	mes.Data = string(data)

	//序列花mesdata
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("sendgroup json.marshal fila", err.Error())
		return
	}

	//获取连接发送消息
	tf := &utils.Tranfans{
		Conn: Curusers.Conn,
	}

	err = tf.Writepkg(data)
	if err != nil {
		fmt.Println("发送失败", err)
		return
	}
	return
}
