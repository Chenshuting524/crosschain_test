package main

import (
	//"flag"

	//"fmt"
	"math/big"
	"sync"

	"crosschain_test/config"
	"crosschain_test/log"
	"crosschain_test/tools"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var confFile = "./config/sampleConfig.json"

//var BigNum = big.NewInt(0).Mul(big.NewInt(1000000000000000000), big.NewInt(1000000000000000000))

type txparam struct {
	fromAsset     *config.Asset
	fromChainId   uint64
	toChainId     uint64
	toAsset       *config.Asset
	toAddress     string
	toPoolId      uint64
	DefaultAmount *big.Int
}

func newtxparam(s string, net *config.Config) *txparam {
	input := &txparam{}
	//选择源链
	ChainIdlist := []uint64{2, 79, 202}
	r := tools.RandomNormalInt64(0, 10000, 5000, 1000)
	//fmt.Println("r", r)
	input.fromChainId = ChainIdlist[r%len(ChainIdlist)]
	//目标链
	for {
		r1 := tools.RandomNormalInt64(0, 10000, 5000, 1000)
		x := ChainIdlist[r1%len(ChainIdlist)]
		//fmt.Println("r1", r)
		if x != input.fromChainId {
			input.toChainId = x
			break
		}
	}
	//input.fromChainId = 202
	//input.toChainId = 2

	//选择pool id
	input.toPoolId = 5

	//根据原链和目标链chainid以及pool id 查询Asset
	input.fromAsset = net.GetAsset(input.fromChainId)
	input.toAsset = net.GetAsset(input.toChainId)

	//数值随机

	randamount := int64(tools.RandomNormalInt64(0, 100000, 50000, 5000))
	//fmt.Println("amount", randamount)
	//baseamount := big.NewInt(1000000)
	//input.DefaultAmount = new(big.Int).Mul(big.NewInt(randamount), baseamount)
	input.DefaultAmount = big.NewInt(randamount)

	input.toAddress = s
	return input
}

func main() {
	//读取配置文件
	conf, err := config.LoadConfig(confFile)
	if err != nil {
		log.Fatal("LoadConfig fail", err)
	}
	net := conf
	//功能list
	alist := []string{"swap", "add", "remove", "lock"}

	//遍历account,每一个account开一个线程(account 添加到config) 执行随机动作
	accounts := conf.Accounts
	var wg sync.WaitGroup

	for i := 0; i < len(accounts); i++ {
		acc := &accounts[i]
		wg.Add(1)
		go func(acc *config.Account, net *config.Config, alist []string) {
			defer wg.Done()
			//每一个account每一轮随机选择一个动作
			for i := 0; i <= 10; i++ {
				log.Info("AccountNum", acc.AccountNum, "cycleNum", i)
				//获取随机数
				//idx := tools.RandomNormalInt64(0, 10000, 5000, 1000)
				//获取随机动作
				//function1 := alist[idx%len(alist)]
				function1 := "lock"
				//根据动作，生成入参
				param := newtxparam(acc.SenderPublicKey, net)
				//进入动作
				RandomFunc(function1, param, acc, net)

				//本轮交易确认后进入下一轮
			}

		}(acc, net, alist)

	}
	wg.Wait()

}

func RandomFunc(function string, param *txparam, acc *config.Account, net *config.Config) {
	switch function {
	case "swap":
		log.Info("Start Swap...", param.fromChainId, " to ", param.toChainId, " amount ", param.DefaultAmount)
		netCfg := net.GetNetwork(param.fromChainId)
		if netCfg == nil {
			log.Errorf("network with chainId %d not found in config file", param.fromChainId)
		}
		client, err := ethclient.Dial(netCfg.Provider)
		if err != nil {
			log.Errorf("fail to dial client %s of network %d", netCfg.Provider, param.fromChainId)
		}
		//数值精度处理
		b := new(big.Int).Exp(big.NewInt(10), big.NewInt(param.fromAsset.TokenPrecision), nil)
		Amount := new(big.Int).Mul(param.DefaultAmount, b)

		err = tools.Approve(client, common.HexToAddress(param.fromAsset.Token), netCfg.Swapper, Amount, netCfg, acc)
		if err != nil {
			log.Fatal("fail to approve ", err)
		}
		txHash, err := tools.Swap(
			client, common.HexToAddress(param.fromAsset.Token), param.toPoolId, param.toChainId,
			common.HexToAddress(param.toAsset.Token).Bytes(), common.HexToAddress(param.toAddress).Bytes(), Amount, netCfg, acc)
		if err != nil {
			log.Fatal("fail while swap ", err)
		}
		log.Info("Swap succeed ", param.fromChainId, " to ", param.toChainId, " tx hash: ", txHash.Hex())

	case "add":
		log.Info("Start Add...", param.fromChainId, " to ", param.toChainId, " amount ", param.DefaultAmount)
		netCfg := net.GetNetwork(param.fromChainId)
		if netCfg == nil {
			log.Errorf("network with chainId %d not found in config file", param.fromChainId)
		}
		client, err := ethclient.Dial(netCfg.Provider)
		if err != nil {
			log.Errorf("fail to dial client %s of network %d", netCfg.Provider, param.fromChainId)
		}

		//数值精度处理
		b := new(big.Int).Exp(big.NewInt(10), big.NewInt(param.fromAsset.TokenPrecision), nil)
		Amount := new(big.Int).Mul(param.DefaultAmount, b)

		err = tools.Approve(client, common.HexToAddress(param.fromAsset.Token), netCfg.Swapper, Amount, netCfg, acc)
		if err != nil {
			log.Fatal("fail to approve ", err)
		}
		txHash, err := tools.Add(
			client, common.HexToAddress(param.fromAsset.Token), param.toPoolId, param.toChainId,
			common.HexToAddress(param.toAddress).Bytes(), Amount, netCfg, acc)
		if err != nil {
			log.Fatal("fail while add ", err)
		}
		log.Info("Add succeed ", param.fromChainId, " to ", param.toChainId, " tx hash: ", txHash.Hex())

	case "remove":
		log.Info("Start remove ...", param.fromChainId, " to ", param.toChainId, " amount ", param.DefaultAmount)
		netCfg := net.GetNetwork(param.fromChainId)
		if netCfg == nil {
			log.Errorf("network with chainId %d not found in config file", param.fromChainId)
		}
		client, err := ethclient.Dial(netCfg.Provider)
		if err != nil {
			log.Errorf("fail to dial client %s of network %d", netCfg.Provider, param.fromChainId)
		}

		//数值精度处理
		b := new(big.Int).Exp(big.NewInt(10), big.NewInt(param.fromAsset.LptokenPrecision), nil)
		Amount := new(big.Int).Mul(param.DefaultAmount, b)

		err = tools.Approve(client, common.HexToAddress(param.fromAsset.Lptoken), netCfg.Swapper, Amount, netCfg, acc)
		if err != nil {
			log.Fatal("fail to approve ", err)
		}

		txHash, err := tools.Remove(
			client, common.HexToAddress(param.fromAsset.Lptoken), param.toPoolId, param.toChainId,
			common.HexToAddress(param.toAsset.Token).Bytes(), common.HexToAddress(param.toAddress).Bytes(), param.DefaultAmount, netCfg, acc)
		if err != nil {
			log.Fatal("fail while remove ", err)
		}
		log.Info("Remove succeed ", param.fromChainId, " to ", param.toChainId, " tx hash: ", txHash.Hex())

	case "lock":
		log.Info("Start lock...", param.fromChainId, " to ", param.toChainId, " amount ", param.DefaultAmount)
		netCfg := net.GetNetwork(param.fromChainId)
		if netCfg == nil {
			log.Errorf("network with chainId %d not found in config file", param.fromChainId)
		}
		client, err := ethclient.Dial(netCfg.Provider)
		if err != nil {
			log.Errorf("fail to dial client %s of network %d", netCfg.Provider, param.fromChainId)
		}
		err = tools.Approve(client, common.HexToAddress(param.fromAsset.NBtoken), netCfg.Wrapper, param.DefaultAmount, netCfg, acc)
		if err != nil {
			log.Fatal("fail to approve ", err)
		}

		txHash, err := tools.Lock(
			client, common.HexToAddress(param.fromAsset.NBtoken), param.toChainId,
			common.HexToAddress(param.toAddress).Bytes(), param.DefaultAmount, netCfg, acc)
		if err != nil {
			log.Fatal("fail while lock ", err)
		}
		log.Info("Lock succeed ", param.fromChainId, " to ", param.toChainId, " tx hash: ", txHash.Hex())

	default:
		log.Fatal("unknown func")
	}
}
