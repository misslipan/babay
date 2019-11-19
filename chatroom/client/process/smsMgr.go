package process

import (
	"chatroom/comm/message"
	"encoding/json"
	"fmt"
)

//显示大家一起聊天信息
func outputonSmsend(msg *message.Message) (err error) {
	var Smssend *message.SmsMsgSend
	err = json.Unmarshal([]byte(msg.Data), &Smssend)
	if err != nil {
		fmt.Println("json.Unmarshal fila", err)
		return
	}
	fmt.Println(Smssend.Userid, "对大家说:", Smssend.Content)
	return
}

//显示一对一的消息
func outOneSmsend(msg *message.Message) (err error) {
	var Smssend *message.SmsMsgSend
	err = json.Unmarshal([]byte(msg.Data), &Smssend)
	if err != nil {
		fmt.Println("json.Unmarshal fila", err)
		return
	}
	fmt.Println(Smssend.Friendid, "对大你说:", Smssend.Content)
	return
}
