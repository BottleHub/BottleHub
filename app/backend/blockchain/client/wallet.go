package client

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func addressHelper(addr string) common.Address {
	address := common.HexToAddress(addr)

	return address
}

func GetBalance(addr string) *big.Float {
	account := addressHelper(addr)
	balance, err := ConnectClient().BalanceAt(context.Background(), account, nil)

	if err != nil {
		log.Fatal(err)
	}

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	return ethValue
}

func GetPendingBalance(addr string) *big.Float {
	account := addressHelper(addr)
	balance, err := ConnectClient().PendingBalanceAt(context.Background(), account)

	if err != nil {
		log.Fatal(err)
	}

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	return ethValue
}

func CreateWallet() (string, string, string, string) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyHex := hexutil.Encode(privateKeyBytes) //[2:]

	hash1 := sha3.NewLegacyKeccak256()
	hash1.Write(privateKeyBytes[1:])
	privateAddress := hexutil.Encode(hash1.Sum(nil)[12:])

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Cannot Assert Type: PublicKey is not of Type *ecdsa.PublicKey")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	publicKeyHex := hexutil.Encode(publicKeyBytes) //[4:]

	hash2 := sha3.NewLegacyKeccak256()
	hash2.Write(publicKeyBytes[1:])
	publicAddress := hexutil.Encode(hash2.Sum(nil)[12:])

	return privateKeyHex, privateAddress, publicKeyHex, publicAddress
}
