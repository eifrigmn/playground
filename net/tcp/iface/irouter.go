package iface

// 路由的抽象接口
type IRouter interface {
	PreHandle(req IRequest)
	Handle(req IRequest)
	PostHandle(req IRequest)
}
