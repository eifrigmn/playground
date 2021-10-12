package server

import (
	"fmt"
	"github.com/eifrigmn/playground/net/tcp/iface"
	"io"
	"net"
)

type Connection struct {
	// 当前连接的socket TCP套接字
	Conn *net.TCPConn
	ConnID uint32
	isClosed bool
	handleAPI iface.HandleFunc
	// 告知当前连接已停止的channel
	ExitChan chan bool
}

// 初始化连接模块
func NewConnection(conn *net.TCPConn, connID uint32, callbackApi iface.HandleFunc) iface.IConnection {
	return &Connection{
		Conn: conn,
		ConnID: connID,
		handleAPI: callbackApi,
		isClosed: false,
		ExitChan: make(chan bool, 1),
	}
}
func (c *Connection) StartReader(){
	fmt.Println("connection reader is running...", c.ConnID)
	defer fmt.Printf("%d reader is exit, remote addr is %s\n", c.ConnID, c.Conn.RemoteAddr().String())
	defer c.Stop()
	for {
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Printf("client id=%d, addr=%s is closed!\n", c.ConnID, c.Conn.RemoteAddr().String())
				break
			}
			fmt.Println("recve buff error", err)
			continue
		}

		if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
			fmt.Printf("%d handle error %v", c.ConnID, err)
			break
		}
	}

}
func (c *Connection) Start(){
	fmt.Println("connection start", c.ConnID)
	go c.StartReader()
}

func (c *Connection) Stop(){
	fmt.Println("stop", c.ConnID)
	if c.isClosed {
		return
	}
	c.isClosed = true
	c.Conn.Close()
	close(c.ExitChan)
}
// 获取当前连接所绑定的socket conn
func (c *Connection) GetTCPConnection() *net.TCPConn{
	return c.Conn
}
// 获取当前连接模块的id
func (c *Connection)GetConnID() uint32 {
	return c.ConnID
}
// 获取远程客户端的TCP状态 IP port
func (c *Connection)GetRemoteAddr() net.Addr{
	return c.Conn.RemoteAddr()
}
// 发送数据，将数据发送给远程的客户端
func (c *Connection)Send(data []byte) error{
	return nil
}
