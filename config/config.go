package config

import (
	"encoding/json"
	"io/ioutil"
)

type NodeConfig struct {
	Server *ServerConfig `json:"server"`
	// The storage
	Storage *StorageConfig `json:"storage"`
}

type RPCConfig struct {
	Host string
}

type TLSConfig struct {
	Enabled  bool   `json:"enabled"`
	KeyPath  string `json:"key"`
	CertPath string `json:"cert"`
}

type ServerConfig struct {
	// The RPC server that should be connected to
	RPC  *RPCConfig `json:"rpc"`
	TLS  *TLSConfig `json:"tls"`
	Web  *WebConfig
	Host string `json:"host"`
}

type StorageConfig struct {
	Bucket  string `json:"bucket"`
	AuthKey string `json:"authKey"`
}

type WebConfig struct {
	MaxStreams int `json:"maxStreams"`
}

func ReadConfig(configPath string) (*NodeConfig, error) {
	var conf NodeConfig

	confFile, err := ioutil.ReadFile(configPath)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(confFile, &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}
