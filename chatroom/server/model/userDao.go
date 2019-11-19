package model

import (
	"chatroom/comm/message"
	"encoding/json"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

var (
	MyuserDao *UserDao
)

type UserDao struct {
	Pool *redis.Pool
}

//工厂模式实例化
func NesUserDao(pool *redis.Pool) (arr *UserDao) {

	arr = &UserDao{
		Pool: pool,
	}
	return
}

func (this *UserDao) GetuserbyId(id int) (user *User, err error) {
	//连接rides
	conn := this.Pool.Get()

	defer conn.Close()

	value, err := redis.String(conn.Do("Hget", "user", id))
	if err != nil {
		if redis.ErrNil == err {
			err = ERROR_LOGING_EXIT
		}
		return
	}
	user = &User{}
	err = json.Unmarshal([]byte(value), user)
	if err != nil {
		fmt.Println("json.unmarshal fail")
		return
	}
	return
}

//验证登陆
func (this *UserDao) LoginCheck(id int, password string) (user *User, err error) {
	user, err = this.GetuserbyId(id)
	if err != nil {
		return
	}
	if user.Password != password {
		err = ERROR_PASSD
		return
	}
	return
}

/**
注册
*/
func (this *UserDao) Register(user *message.User) (err error) {
	conn := this.Pool.Get()
	defer conn.Close()
	_, err = this.GetuserbyId(user.Userid)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}

	datas, err := json.Marshal(user)
	_, err = conn.Do("Hset", "user", user.Userid, string(datas))
	return
}
