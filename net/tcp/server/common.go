package server

import (
	"github.com/eifrigmn/playground/net/tcp/iface"
	"net"
)

type tcpServer struct {
	listener *net.TCPListener
}

// zinx-v1.0
type ServerV1 struct {
	Name string
	IPVersion string
	IP string
	Port int
}

type ServerV3 struct {
	Name string
	IPVersion string
	IP string
	Port int
	// 当前Server添加一个router
	Router iface.IRouter
}
