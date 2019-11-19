package message

type User struct {
	Userid     int    `json:"userid"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	UserStatus int    `json:"userstatus"`
}
