package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

//获取network配置
func TestGenerateConfig(t *testing.T) {
	//获取链配置
	Eth := Network{
		2,
		"eth",
		"https://ropsten.infura.io/v3/7902ff0c436e4909b4f17b61a5cc9479",
		common.HexToAddress("0xDc37471Af6a8aB7f45F444c5a3Ef4758281bE32C"),
		common.HexToAddress("0x1E7A3e54494F300dC66181621E23eE657E22D725"),
	}
	Bsc := Network{
		79,
		"bsc",
		"https://data-seed-prebsc-1-s1.binance.org:8545/",
		common.HexToAddress("0x9f9F15CC407F7b26f55D71D43f993580a9107007"),
		common.HexToAddress("0x6407673152fc40B648C3348BeE66D162d39923c6"),
	}
	Heco := Network{
		7,
		"heco",
		"https://http-testnet.hecochain.com",
		common.HexToAddress("0xCC8407Ee04AaC2AdC0E6A55E7E97176C701146cd"),
		common.HexToAddress("0x5B18542A4cD54C7c1641bdF63541F19072b79260"),
	}
	Matic := Network{
		202,
		"matic",
		"https://rpc-mumbai.maticvigil.com",
		common.HexToAddress("0xD5d63Dce45E0275Ca76a8b2e9BD8C11679A57D0D"),
		common.HexToAddress("0x1B0C55be400e2a7D924032B257Fbc75Bbfd256E7"),
	}
	nets := []Network{Eth, Bsc, Heco, Matic}

	//获取account配置
	A1 := Account{
		1,
		"0x6Ac449ADE24174238DF325749bD5ea87B02BF7f6",
		"privatekey",
	}
	A2 := Account{
		2,
		"0x0c888cca1190940ebC156d4Cf13cbF880A83E4A3",
		"privatekey",
	}

	accounts := []Account{A1, A2}

	//获取Asset配置
	Eth_Token := Asset{
		2,
		"0xf5FfaDc05ac8A2f451832B5a06Fab7d3EC55fd27",
		6,
		"0xe98D6AeEF5E76E5eB04fE1Ab8f4659c73671da60",
		18,
		"0xf4d3994fe7ad65485e63d7f139c72ea81becefa4",
	}
	Bsc_Token := Asset{
		79,
		"0x9c5e484438fbFf492f23f52a7f666F60122360F8",
		18,
		"0x08b19C6405D7Fe6C8651a166eC1563907574159c",
		18,
		"0x7466366d6d303f41a1876c45b1acfbc2b17123e4",
	}
	Polygon_Token := Asset{
		202,
		"0xF6C69CdF9F03Ec22d12FED30f218b36A0Cd31Bb2",
		6,
		"0x34d4a23A1FC0C694f0D74DDAf9D8d564cfE2D430",
		18,
		"0x76d56873426cee920f9e8c439ee0a2cecf50d979",
	}
	tokens := []Asset{Eth_Token, Polygon_Token, Bsc_Token}

	res, err := json.Marshal(&Config{nets, accounts, tokens})
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
		fmt.Println(conf.Networks[1])
	}
	account1 := conf.GetAccount(66)
	if account1 == nil {
		fmt.Println("2 nil")
		fmt.Println(conf.Accounts)
	}
	return
	Assets1 := conf.GetAsset(66)
	if Assets1 == nil {
		fmt.Println("2 nil")
		fmt.Println(conf.Assets)
	}
	return

}
