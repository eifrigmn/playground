package iface

import "net"

type IConnection interface {
	// 启动连接，当前连接准备开始工作
	Start()
	// 停止连接
	Stop()
	// 获取当前连接所绑定的socket conn
	GetTCPConnection() *net.TCPConn
	// 获取当前连接模块的id
	GetConnID() uint32
	// 获取远程客户端的TCP状态 IP port
	GetRemoteAddr() net.Addr
	// 发送数据，将数据发送给远程的客户端
	Send(data []byte) error
}

// 定义一个处理连接业务的方法
// conn 当前连接
// data 待处理的数据
// cnt 待处理的数据长度
type HandleFunc func(conn *net.TCPConn, data []byte, cnt int) error