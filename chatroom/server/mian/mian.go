package main

import (
	"chatroom/server/model"
	"fmt"
	"net"
)

//处理客户端通信
func processr(conn net.Conn) {
	defer conn.Close()
	Pr := &Processr{
		Conn: conn,
	}
	err := Pr.Process1()
	if err != nil {
		fmt.Println("kes  ", err)
		return
	}

}

//实梨花UserDao
func initUserdao() {
	model.MyuserDao = model.NesUserDao(Pool)
}

func main() {
	//创建连接池
	Initredis()
	//实例化redis连接
	initUserdao()
	fmt.Println("服务器监听中.......")
	listen, err := net.Listen("tcp", ":8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("监听失败")
		return
	}

	//一旦监听成功,等待连接
	for {
		fmt.Println("等待客户端连接....")

		conn, err := listen.Accept()

		if err != nil {
			fmt.Println("客户端连接失败")
			continue
		}

		//处理客户端

		go processr(conn)

	}

}
