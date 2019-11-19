package message

const (
	LonginType         = "LoginMes"
	LoginReMmsgType    = "LoginRes"
	RegisterResMesType = "RegisterResMes" //注册返回消息常量
	RegisterErrType    = "Registererr"
	NotifystatusType   = "NotifystatusType" //用户状态常量类型
	SmsMsgSendType     = "Smsmsgsend"
	SmsMsgOneType      = "OnevsOnesend" //一对一发送消息
)

const (
	Onliestatus = iota
	Lixianstatus
)

//消息结构体
type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息内容
}

//登陆消息内容
type LoginMsg struct {
	UserId   int    `json:"userid"`
	PassWord string `json:"password"`
	UserName string `json:"username"`
}

//登陆返回消息内容
type LoginReMmsg struct {
	Code     int `json:"code"` //状态码
	Usersids []int
	ErrorMsg string `json:"errormsg"` //返回信息
}

//注册模型
type RegisterMes struct {
	User User
}

type RegisterResMes struct {
	Code     int    `json:"code"`     //状态码
	ErrorMsg string `json:"errormsg"` //返回信息
}

/**
用户状态
*/
type Notifystatus struct {
	Userid int `json:"userid"`
	Status int `json:"status"`
}

/**
消息发送
*/
type SmsMsgSend struct {
	Content string `json:"content"`
	User
	Friendid int `json:"friendid"`
}
