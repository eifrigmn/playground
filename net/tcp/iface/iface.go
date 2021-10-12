package iface



// 服务端接口
type Server interface {
	Start()
	Stop()
	Serve()
}

type ServerV3 interface {
	Server
	// 路由功能：给当前的服务注册一个路由方法，供客户端的连接处理使用
	AddRouter(r IRouter)
}
