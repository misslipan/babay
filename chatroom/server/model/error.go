package model

import (
	"errors"
)

var (
	ERROR_LOGING_EXIT = errors.New("该用户不存在请注册")
	ERROR_PASSD       = errors.New("密码错误")
	ERROR_USER_EXISTS = errors.New("用户已存在")
)
