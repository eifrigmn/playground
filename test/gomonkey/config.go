package gomonkey

import (
	"encoding/json"
	"os"
)

type Config struct {
	Name string
	Data string
}
func LoadConfig(path string)(*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var cfgByte []byte
	_,err = f.Read(cfgByte)
	if err != nil {
		return nil, err
	}
	var cfg Config
	err = json.Unmarshal(cfgByte, &cfg)
	return &cfg, err

}

func (c *Config)GetConfigName() string {
	return c.Name
}
