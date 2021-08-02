package tools

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/KSlashh/crosschain_test/abi"
	"github.com/KSlashh/crosschain_test/config"
	"github.com/KSlashh/crosschain_test/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var DefaultGasLimit uint64 = 300000
var DefaultId *big.Int = big.NewInt(15)
var FeeRPC string = "https://bridge.poly.network/testnet/v1/getfee/"
var AddressZero string = "0000000000000000000000000000000000000000"
var Zero *big.Int = big.NewInt(0)
var EmptyHash = common.Hash{}

type GetFeeReq struct {
	SrcChainId    uint64
	Hash          string
	SwapTokenHash string
	DstChainId    uint64
}

type GetFeeRsp struct {
	SrcChainId               uint64
	Hash                     string
	DstChainId               uint64
	UsdtAmount               string
	TokenAmount              string
	TokenAmountWithPrecision string
	SwapTokenHash            string
	Balance                  string
	BalanceWithPrecision     string
}

func Swap(client *ethclient.Client, fromAsset common.Address, toPoolId uint64, toChainId uint64, toAsset []byte, toAddress []byte, amount *big.Int, conf *config.Network) (common.Hash, error) {
	privateKey, err := crypto.HexToECDSA(conf.SenderPrivateKey)
	if err != nil {
		return EmptyHash, fmt.Errorf(fmt.Sprint("fail to swap ", err))
	}
	ISwapperContract, err := abi.NewISwapper(conf.Swapper, client)
	if err != nil {
		return EmptyHash, fmt.Errorf(fmt.Sprint("fail to swap ", err))
	}
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return EmptyHash, fmt.Errorf(fmt.Sprint("fail to swap ", err))
	}
	fee, err := GetFee(conf.PolyChainID, toChainId, AddressZero)
	if err != nil {
		return EmptyHash, fmt.Errorf(fmt.Sprint("fail to swap ", err))
	}
	auth, err := MakeAuth(client, privateKey, fee, DefaultGasLimit, chainId)
	if err != nil {
		return EmptyHash, fmt.Errorf(fmt.Sprint("fail to swap ", err))
	}
	tx, err := ISwapperContract.Swap(auth, fromAsset, toPoolId, toChainId, toAsset, toAddress, amount, Zero, fee, DefaultId)
	if err != nil {
		return EmptyHash, fmt.Errorf(fmt.Sprint("fail to swap ", err))
	}
	err = WaitTxConfirm(client, tx.Hash())
	if err != nil {
		return EmptyHash, fmt.Errorf(fmt.Sprint("fail to swap ", err))
	}
	return tx.Hash(), nil
}

func Add(client *ethclient.Client, fromAsset common.Address, toPoolId uint64, toChainId uint64, toAddress []byte, amount *big.Int, conf *config.Network) (common.Hash, error) {
	privateKey, err := crypto.HexToECDSA(conf.SenderPrivateKey)
	if err != nil {
		return EmptyHash, fmt.Errorf(fmt.Sprint("fail to AddLiquidity ", err))
	}
	ISwapperContract, err := abi.NewISwapper(conf.Swapper, client)
	if err != nil {
		return EmptyHash, fmt.Errorf(fmt.Sprint("fail to AddLiquidity ", err))
	}
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return EmptyHash, fmt.Errorf(fmt.Sprint("fail to AddLiquidity ", err))
	}
	fee, err := GetFee(conf.PolyChainID, toChainId, AddressZero)
	if err != nil {
		return EmptyHash, fmt.Errorf(fmt.Sprint("fail to AddLiquidity ", err))
	}
	auth, err := MakeAuth(client, privateKey, fee, DefaultGasLimit, chainId)
	if err != nil {
		return EmptyHash, fmt.Errorf(fmt.Sprint("fail to AddLiquidity ", err))
	}
	tx, err := ISwapperContract.AddLiquidity(auth, fromAsset, toPoolId, toChainId, toAddress, amount, Zero, fee, DefaultId)
	if err != nil {
		return EmptyHash, fmt.Errorf(fmt.Sprint("fail to AddLiquidity ", err))
	}
	err = WaitTxConfirm(client, tx.Hash())
	if err != nil {
		return EmptyHash, fmt.Errorf(fmt.Sprint("fail to AddLiquidity ", err))
	}
	return tx.Hash(), nil
}

func Remove(client *ethclient.Client, fromAsset common.Address, toPoolId uint64, toChainId uint64, toAsset []byte, toAddress []byte, amount *big.Int, conf *config.Network) (common.Hash, error) {
	privateKey, err := crypto.HexToECDSA(conf.SenderPrivateKey)
	if err != nil {
		return EmptyHash, fmt.Errorf(fmt.Sprint("fail to RemoveLiquidity ", err))
	}
	ISwapperContract, err := abi.NewISwapper(conf.Swapper, client)
	if err != nil {
		return EmptyHash, fmt.Errorf(fmt.Sprint("fail to RemoveLiquidity ", err))
	}
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return EmptyHash, fmt.Errorf(fmt.Sprint("fail to RemoveLiquidity ", err))
	}
	fee, err := GetFee(conf.PolyChainID, toChainId, AddressZero)
	if err != nil {
		return EmptyHash, fmt.Errorf(fmt.Sprint("fail to RemoveLiquidity ", err))
	}
	auth, err := MakeAuth(client, privateKey, fee, DefaultGasLimit, chainId)
	if err != nil {
		return EmptyHash, fmt.Errorf(fmt.Sprint("fail to RemoveLiquidity ", err))
	}
	tx, err := ISwapperContract.RemoveLiquidity(auth, fromAsset, toPoolId, toChainId, toAsset, toAddress, amount, Zero, fee, DefaultId)
	if err != nil {
		return EmptyHash, fmt.Errorf(fmt.Sprint("fail to RemoveLiquidity ", err))
	}
	err = WaitTxConfirm(client, tx.Hash())
	if err != nil {
		return EmptyHash, fmt.Errorf(fmt.Sprint("fail to RemoveLiquidity ", err))
	}
	return tx.Hash(), nil
}

func Lock(client *ethclient.Client, fromAsset common.Address, toChainId uint64, toAddress []byte, amount *big.Int, conf *config.Network) (common.Hash, error) {
	privateKey, err := crypto.HexToECDSA(conf.SenderPrivateKey)
	if err != nil {
		return EmptyHash, fmt.Errorf(fmt.Sprint("fail to lock ", err))
	}
	IWrapperContract, err := abi.NewIWrapper(conf.Wrapper, client)
	if err != nil {
		return EmptyHash, fmt.Errorf(fmt.Sprint("fail to lock ", err))
	}
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return EmptyHash, fmt.Errorf(fmt.Sprint("fail to lock ", err))
	}
	fee, err := GetFee(conf.PolyChainID, toChainId, AddressZero)
	if err != nil {
		return EmptyHash, fmt.Errorf(fmt.Sprint("fail to lock ", err))
	}
	auth, err := MakeAuth(client, privateKey, fee, DefaultGasLimit, chainId)
	if err != nil {
		return EmptyHash, fmt.Errorf(fmt.Sprint("fail to lock ", err))
	}
	tx, err := IWrapperContract.Lock(auth, fromAsset, toChainId, toAddress, amount, fee, DefaultId)
	if err != nil {
		return EmptyHash, fmt.Errorf(fmt.Sprint("fail to lock ", err))
	}
	err = WaitTxConfirm(client, tx.Hash())
	if err != nil {
		return EmptyHash, fmt.Errorf(fmt.Sprint("fail to lock ", err))
	}
	return tx.Hash(), nil
}

func Approve(client *ethclient.Client, token, spender common.Address, amount *big.Int, conf *config.Network) error {
	privateKey, err := crypto.HexToECDSA(conf.SenderPrivateKey)
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("fail to approve token %s ,", token.Hex()), err)
	}
	IERC20Contract, err := abi.NewIERC20(token, client)
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("fail to approve token %s ,", token.Hex()), err)
	}
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("fail to approve token %s ,", token.Hex()), err)
	}
	auth, err := MakeAuth(client, privateKey, Zero, DefaultGasLimit, chainId)
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("fail to approve token %s ,", token.Hex()), err)
	}
	tx, err := IERC20Contract.Approve(auth, spender, amount)
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("fail to approve token %s ,", token.Hex()), err)
	}
	err = WaitTxConfirm(client, tx.Hash())
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("fail to approve token %s ,", token.Hex()), err)
	}
	return nil
}

func AllowanceOf(client *ethclient.Client, token, owner, spender common.Address) (*big.Int, error) {
	ERC20Contract, err := abi.NewIERC20Caller(token, client)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("fail while get allowance token of %s ,", token.Hex()), err)
	}
	return ERC20Contract.Allowance(nil, owner, spender)
}

func MakeAuth(client *ethclient.Client, key *ecdsa.PrivateKey, value *big.Int, gasLimit uint64, chainId *big.Int) (*bind.TransactOpts, error) {
	authAddress := crypto.PubkeyToAddress(*key.Public().(*ecdsa.PublicKey))
	nonce, err := client.PendingNonceAt(context.Background(), authAddress)
	if err != nil {
		return nil, fmt.Errorf("makeAuth, addr %s, err %v", authAddress.Hex(), err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("makeAuth, get suggest gas price err: %v", err)
	}

	// auth := bind.NewKeyedTransactor(key)
	auth, err := bind.NewKeyedTransactorWithChainID(key, chainId)
	if err != nil {
		return nil, fmt.Errorf("makeAuth, bind.NewKeyedTransactorWithChainID err: %v", err)
	}
	auth.From = authAddress
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = value
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice

	return auth, nil
}

func WaitTxConfirm(client *ethclient.Client, hash common.Hash) error {
	ticker := time.NewTicker(time.Second * 1)
	end := time.Now().Add(60 * time.Second)
	for now := range ticker.C {
		_, pending, err := client.TransactionByHash(context.Background(), hash)
		if err != nil {
			log.Debug("failed to call TransactionByHash: %v", err)
			continue
		}
		if !pending {
			break
		}
		if now.Before(end) {
			continue
		}
		log.Info("Transaction pending for more than 1 min, check transaction %s on explorer yourself, make sure it's confirmed.", hash.Hex())
		return nil
	}

	tx, err := client.TransactionReceipt(context.Background(), hash)
	if err != nil {
		return fmt.Errorf("faild to get receipt %s", hash.Hex())
	}

	if tx.Status == 0 {
		return fmt.Errorf("receipt failed %s", hash.Hex())
	}

	return nil
}

func GetFee(SrcChainId, DstChainId uint64, Hash string) (*big.Int, error) {
	para := &GetFeeReq{SrcChainId, Hash, Hash, DstChainId}
	data, err := json.Marshal(para)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", FeeRPC, strings.NewReader(string(data)))
	req.Header.Set("Accepts", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("response status code: %d", resp.StatusCode)
	}
	respBody, _ := ioutil.ReadAll(resp.Body)
	res := new(GetFeeRsp)
	err = json.Unmarshal(respBody, res)
	if err != nil {
		return nil, err
	}
	f, err := strconv.ParseFloat(res.TokenAmountWithPrecision, 64)
	if err != nil {
		return nil, fmt.Errorf("response return strange TokenAmountWithPrecision %s, %s", res.TokenAmountWithPrecision, err.Error())
	}
	fee := big.NewInt(int64(f) + 1)
	return fee, nil
}
