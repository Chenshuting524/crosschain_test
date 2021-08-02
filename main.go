package main

import (
	"flag"
	"math/big"

	"github.com/KSlashh/crosschain_test/config"
	"github.com/KSlashh/crosschain_test/log"
	"github.com/KSlashh/crosschain_test/tools"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var confFile string
var function string
var toAddress string
var fromAsset string
var toAsset string
var toPoolId uint64
var toChainId uint64
var fromChainId uint64

var DefaultAmount = big.NewInt(666)
var BigNum = big.NewInt(0).Mul(big.NewInt(1000000000000000000), big.NewInt(1000000000000000000))

func init() {
	flag.StringVar(&confFile, "conf", "./config.json", "configuration file path")
	flag.StringVar(&toAddress, "toAddress", "0x7FB1484882e4A3A7a4e31f0eb33bf3dD3d95f797", "")
	flag.StringVar(&fromAsset, "fromAsset", "0x9c5e484438fbFf492f23f52a7f666F60122360F8", "")
	flag.StringVar(&toAsset, "toAsset", "0xF6C69CdF9F03Ec22d12FED30f218b36A0Cd31Bb2", "")
	flag.Uint64Var(&fromChainId, "fromChainId", 79, "")
	flag.Uint64Var(&toChainId, "toChainId", 202, "")
	flag.Uint64Var(&toPoolId, "toPoolId", 5, "")
	flag.StringVar(&function, "func", "add", "choose function to run:\n"+
		"  add\n"+
		"  remove\n"+
		"  swap\n"+
		"  lock")
	flag.Parse()
}

func main() {
	conf, err := config.LoadConfig(confFile)
	if err != nil {
		log.Fatal("LoadConfig fail", err)
	}
	switch function {
	case "swap":
		log.Info("Start...")
		netCfg := conf.GetNetwork(fromChainId)
		if netCfg == nil {
			log.Errorf("network with chainId %d not found in config file", fromChainId)
		}
		client, err := ethclient.Dial(netCfg.Provider)
		if err != nil {
			log.Errorf("fail to dial client %s of network %d", netCfg.Provider, fromChainId)
		}
		err = tools.Approve(client, common.HexToAddress(fromAsset), netCfg.Swapper, BigNum, netCfg)
		if err != nil {
			log.Fatal("fail to approve ", err)
		}
		cnt := 0
		for {
			cnt += 1
			txHash, err := tools.Swap(
				client, common.HexToAddress(fromAsset), toPoolId, toChainId,
				common.HexToAddress(toAsset).Bytes(), common.HexToAddress(toAddress).Bytes(), DefaultAmount, netCfg)
			if err != nil {
				log.Fatal("fail while swap ", err)
			}
			log.Info(cnt, " Swap succeed, tx hash: ", txHash.Hex())
		}
	case "add":
		log.Info("Start...")
		netCfg := conf.GetNetwork(fromChainId)
		if netCfg == nil {
			log.Errorf("network with chainId %d not found in config file", fromChainId)
		}
		client, err := ethclient.Dial(netCfg.Provider)
		if err != nil {
			log.Errorf("fail to dial client %s of network %d", netCfg.Provider, fromChainId)
		}
		err = tools.Approve(client, common.HexToAddress(fromAsset), netCfg.Swapper, BigNum, netCfg)
		if err != nil {
			log.Fatal("fail to approve ", err)
		}
		cnt := 0
		for {
			cnt += 1
			txHash, err := tools.Add(
				client, common.HexToAddress(fromAsset), toPoolId, toChainId,
				common.HexToAddress(toAddress).Bytes(), DefaultAmount, netCfg)
			if err != nil {
				log.Fatal("fail while add ", err)
			}
			log.Info(cnt, " Add succeed, tx hash: ", txHash.Hex())
		}
	case "remove":
		log.Info("Start...")
		netCfg := conf.GetNetwork(fromChainId)
		if netCfg == nil {
			log.Errorf("network with chainId %d not found in config file", fromChainId)
		}
		client, err := ethclient.Dial(netCfg.Provider)
		if err != nil {
			log.Errorf("fail to dial client %s of network %d", netCfg.Provider, fromChainId)
		}
		err = tools.Approve(client, common.HexToAddress(fromAsset), netCfg.Swapper, BigNum, netCfg)
		if err != nil {
			log.Fatal("fail to approve ", err)
		}
		cnt := 0
		for {
			cnt += 1
			txHash, err := tools.Remove(
				client, common.HexToAddress(fromAsset), toPoolId, toChainId,
				common.HexToAddress(toAsset).Bytes(), common.HexToAddress(toAddress).Bytes(), DefaultAmount, netCfg)
			if err != nil {
				log.Fatal("fail while remove ", err)
			}
			log.Info(cnt, " Remove succeed, tx hash: ", txHash.Hex())
		}
	case "lock":
		log.Info("Start...")
		netCfg := conf.GetNetwork(fromChainId)
		if netCfg == nil {
			log.Errorf("network with chainId %d not found in config file", fromChainId)
		}
		client, err := ethclient.Dial(netCfg.Provider)
		if err != nil {
			log.Errorf("fail to dial client %s of network %d", netCfg.Provider, fromChainId)
		}
		err = tools.Approve(client, common.HexToAddress(fromAsset), netCfg.Wrapper, BigNum, netCfg)
		if err != nil {
			log.Fatal("fail to approve ", err)
		}
		cnt := 0
		for {
			cnt += 1
			txHash, err := tools.Lock(
				client, common.HexToAddress(fromAsset), toChainId,
				common.HexToAddress(toAddress).Bytes(), DefaultAmount, netCfg)
			if err != nil {
				log.Fatal("fail while lock ", err)
			}
			log.Info(cnt, " Lock succeed, tx hash: ", txHash.Hex())
		}
	default:
		log.Fatal("unknown func")
	}
}
