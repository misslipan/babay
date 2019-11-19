package process

import (
	"fmt"
)

var (
	UserMgrs *UserMgr
)

type UserMgr struct {
	OnlieUsers map[int]*UserProcess
}

//初始化
func init() {

	UserMgrs = &UserMgr{
		OnlieUsers: make(map[int]*UserProcess, 1024),
	}

}

//添加用户信息进数组
func (this *UserMgr) AddonliUser(up *UserProcess) {
	this.OnlieUsers[up.Userid] = up
}

//删除
func (this *UserMgr) DelonliUser(userId int) {
	delete(this.OnlieUsers, userId)

}

//返回全部
func (this *UserMgr) Getallonli() map[int]*UserProcess {
	return this.OnlieUsers
}

//根据用户id获取用户
func (this *UserMgr) GetonileuserById(userId int) (up *UserProcess, err error) {
	up, ok := this.OnlieUsers[userId]
	if !ok {
		err = fmt.Errorf("用户%d不存在", userId)
		return
	}
	return
}
