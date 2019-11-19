package model

type User struct {
	Userid   int    `json:"userid"`
	Username string `json:"username"`
	Password string `json:"password"`
}
