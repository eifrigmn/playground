package config

type TcpConfig struct {
	Host string
	Port int
	IpVersion string
}

func LoadConfig() *TcpConfig {
	return &TcpConfig{
		Host: "0.0.0.0",
		Port: 8972,
		IpVersion: "tcp",
	}
}
