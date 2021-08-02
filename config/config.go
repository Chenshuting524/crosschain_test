package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/common"
)

type Network struct {
	PolyChainID      uint64
	Name             string
	Provider         string
	Wrapper          common.Address
	Swapper          common.Address
	SenderPrivateKey string
}

type Config struct {
	Networks []Network
}

func LoadConfig(confFile string) (config *Config, err error) {
	jsonBytes, err := ioutil.ReadFile(confFile)
	if err != nil {
		return
	}

	config = &Config{}
	err = json.Unmarshal(jsonBytes, config)
	return
}

func (c *Config) GetNetwork(index uint64) (netConfig *Network) {
	for i := 0; i < len(c.Networks); i++ {
		if c.Networks[i].PolyChainID == index {
			return &c.Networks[i]
		}
	}
	return nil
}
