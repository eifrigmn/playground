package server

import (
	"github.com/eifrigmn/playground/net/tcp/iface"
)

type Request struct {
	conn iface.IConnection
	data []byte
}

// 得到当前连接
func (r *Request) GetConnection() iface.IConnection{
	return r.conn
}
// 得到请求的消息数据
func (r *Request)GetData() []byte {
	return r.data
}
