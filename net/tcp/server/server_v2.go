package server

import (
	"fmt"
	"github.com/eifrigmn/playground/net/tcp/config"
	"github.com/eifrigmn/playground/net/tcp/iface"
	"net"
)

// 初始化
type ServerV2 ServerV1
func NewServerV2(name string, cfg *config.TcpConfig) iface.Server {
	return &ServerV2{
		Name:      name,
		IPVersion: cfg.IpVersion,
		IP:        cfg.Host,
		Port:      cfg.Port,
	}
}

func (s *ServerV2) Stop() {

}
func (s *ServerV2) Start() {
	fmt.Printf("[Start] Server Listenner at %s:%d, starting...\n", s.IP, s.Port)
	go func() {
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve error", err)
			return
		}

		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen error", err)
			return
		}
		fmt.Printf("[Start] Server Listenner at %s:%d, success!\n", s.IP, s.Port)
		var cid uint32 = 0
		for {
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept error", err)
				continue
			}
			//go func() {
			//	for {
			//		buf := make([]byte, 512)
			//		cnt, err := conn.Read(buf)
			//		if err != nil {
			//			if err == io.EOF{
			//				fmt.Printf("client %s is closed!\n", conn.RemoteAddr().String())
			//				conn.Close()
			//				return
			//			}
			//			fmt.Println("receive buf error", err)
			//			continue
			//		}
			//		fmt.Printf("Server receive buf = %s, cnt = %d\n", buf, cnt)
			//		_, err = conn.Write(buf[:cnt])
			//		if err != nil {
			//			fmt.Println("write back error", err)
			//			continue
			//		}
			//	}
			//}()
			dealConn := NewConnection(conn, cid, CallBackToClient)
			cid++
			go dealConn.Start()
		}
	}()
}
func (s *ServerV2) Serve() {
	s.Start()
	select{}
}

func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	//fmt.Println("Conn Handler...")
	fmt.Printf("Server receive buf = %s, cnt = %d\n", data, cnt)
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("write back buff error", err)
		return err
	}
	return nil
}
