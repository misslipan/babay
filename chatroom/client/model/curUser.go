package model

import (
	"chatroom/comm/message"
	"net"
)

type CurUser struct {
	Conn net.Conn
	message.User
}
