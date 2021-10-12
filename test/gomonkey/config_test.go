package gomonkey

import (
	. "github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig_GetConfigName(t *testing.T) {
	p := ApplyFunc(LoadConfig, func(_ string)(*Config,error){
		return &Config{
			Name: "mock_config",
			Data: "mock data of config",
		}, nil
	})
	defer p.Reset()
	cfg, err := LoadConfig("")
	assert.Equal(t, nil, err)
	assert.Equal(t, "mock_config", cfg.GetConfigName())
}
