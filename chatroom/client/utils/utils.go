package utils

import (
	"chatroom/comm/message"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

type Tranfans struct {
	Conn net.Conn
	buf  [8096]byte
}

//读取包
func (this *Tranfans) Readpkg() (msg message.Message, err error) {
	_, err = this.Conn.Read(this.buf[:4])
	if err != nil {
		//err = errors.New("read pkg  header error")

		return
	}

	var pkglen uint32
	//类型转换
	pkglen = binary.BigEndian.Uint32(this.buf[:4])

	n, err := this.Conn.Read(this.buf[:pkglen])

	if n != int(pkglen) || err != nil {
		err = errors.New("read pkg  neirong error")
		return
	}

	err = json.Unmarshal(this.buf[:pkglen], &msg)
	if err != nil {
		err = errors.New("read Unmarshal error")
		return
	}

	return
}

//服务器写操作
func (this *Tranfans) Writepkg(data []byte) (err error) {
	//1获取当前内容长度
	var pkg uint32

	pkg = uint32(len(data))

	//var bytes [4]byte
	//2在把内容放松给客户端
	binary.BigEndian.PutUint32(this.buf[:4], pkg)

	//发送内容长度
	n, err := this.Conn.Write(this.buf[:4])
	if n != 4 || err != nil {
		fmt.Println("发送长度失败", err)
		return
	}
	fmt.Println("向服务器写入数据", string(data), this.buf[:4])
	n, err = this.Conn.Write(data)
	fmt.Println(n, "成功写入")
	if n != int(pkg) || err != nil {
		fmt.Println("发送长度失败", err)
		return
	}

	//处理返回消息接收

	return

}
