package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/common"
)

type Network struct {
	PolyChainID uint64
	Name        string
	Provider    string
	Wrapper     common.Address
	Swapper     common.Address
}

type Account struct {
	AccountNum       uint64
	SenderPublicKey  string
	SenderPrivateKey string
}

type Asset struct {
	ChainId          uint64
	Token            string
	TokenPrecision   int64
	Lptoken          string
	LptokenPrecision int64
	NBtoken          string
}

type Config struct {
	Networks []Network
	Accounts []Account
	Assets   []Asset
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

func (c *Config) GetAccount(index uint64) (accountConfig *Account) {
	for i := 0; i < len(c.Accounts); i++ {
		if c.Accounts[i].AccountNum == index {
			return &c.Accounts[i]
		}
	}
	return nil
}

func (c *Config) GetAsset(index uint64) (tokenConfig *Asset) {
	for i := 0; i < len(c.Assets); i++ {
		if c.Assets[i].ChainId == index {
			return &c.Assets[i]
		}
	}
	return nil
}
