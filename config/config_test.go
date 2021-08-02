package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestGenerateConfig(t *testing.T) {
	Eth := Network{
		2,
		"eth",
		"https://mainnet.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161",
		common.HexToAddress("0xDc37471Af6a8aB7f45F444c5a3Ef4758281bE32C"),
		common.HexToAddress("0x1E7A3e54494F300dC66181621E23eE657E22D725"),
		"{put sender's sk here}",
	}
	Bsc := Network{
		79,
		"bsc",
		"https://bsc-dataseed.binance.org",
		common.HexToAddress("0x9f9F15CC407F7b26f55D71D43f993580a9107007"),
		common.HexToAddress("0x6407673152fc40B648C3348BeE66D162d39923c6"),
		"{put sender's sk here}",
	}
	Heco := Network{
		7,
		"heco",
		"https://http-mainnet-node.huobichain.com",
		common.HexToAddress("0xCC8407Ee04AaC2AdC0E6A55E7E97176C701146cd"),
		common.HexToAddress("0x5B18542A4cD54C7c1641bdF63541F19072b79260"),
		"{put sender's sk here}",
	}
	Matic := Network{
		202,
		"matic",
		"http://rpc-mumbai.matic.today",
		common.HexToAddress("0xD5d63Dce45E0275Ca76a8b2e9BD8C11679A57D0D"),
		common.HexToAddress("0x1B0C55be400e2a7D924032B257Fbc75Bbfd256E7"),
		"{put sender's sk here}",
	}
	nets := []Network{Eth, Bsc, Heco, Matic}
	res, err := json.Marshal(&Config{nets})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s \n", res)
	err = ioutil.WriteFile("./sampleConfig.json", res, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestReadConfig(t *testing.T) {
	conf, err := LoadConfig("sampleConfig.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	net1 := conf.GetNetwork(12)
	if net1 == nil {
		fmt.Println("1 nil")
		fmt.Println(conf.Networks[3])
		return
	}
	fmt.Printf("chain %d has name %s", net1.PolyChainID, net1.Name)
}
