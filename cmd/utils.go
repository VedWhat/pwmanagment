package cmd

import (
	"context"
	"io/ioutil"
	"log"

	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/client/rpc"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/program/sysprog"
	"github.com/portto/solana-go-sdk/types"
)

type Wallet struct {
	account types.Account
	c       *client.Client
}

func createNewWallet(RPCEndPoint string) Wallet {
	newAcc := types.NewAccount()
	data := []byte(newAcc.PrivateKey)
	err := ioutil.WriteFile("data", data, 0644)
	if err != nil {
		log.Fatal("Error while creating wallet:{}", err)
	}

	return Wallet{
		account: newAcc,
		c:       client.NewClient(RPCEndPoint),
	}
}

func importOldWallet(RPCEndpoint string) (Wallet, error) {
	contents, err := ioutil.ReadFile("data")
	privateKey := []byte(string(contents))
	if err != nil {
		log.Fatal("Error reading file:{}", err)
	}
	account, err := types.AccountFromBytes(privateKey)
	if err != nil {
		log.Fatal("Error while importing wallet:{}", err)
		return Wallet{}, err
	}

	return Wallet{
		account: account,
		c:       client.NewClient(RPCEndpoint),
	}, nil

}

func getBalance() (uint64, error) {
	wallet, _ := importOldWallet(rpc.DevnetRPCEndpoint)
	balance, err := wallet.c.GetBalance(
		context.TODO(),
		wallet.account.PublicKey.ToBase58(),
	)

	if err != nil {
		return 0, err
	}
	return balance, nil

}

func requestAirdrop(amount uint64) (string, error) {
	// request for SOL using RequestAirdrop()
	wallet, _ := importOldWallet(rpc.DevnetRPCEndpoint)
	amount = amount * 1e9 // turning SOL into lamports
	txhash, err := wallet.c.RequestAirdrop(
		context.TODO(),
		wallet.account.PublicKey.ToBase58(),
		amount,
	)
	if err != nil {
		log.Fatal("Error while requesting airdrop:{}", err)
		return "", err
	}
	return txhash, nil
}

func transferFunds(receiver string, amount uint64) (string, error) {
	wallet, _ := importOldWallet(rpc.DevnetRPCEndpoint)
	response, err := wallet.c.GetRecentBlockhash(context.TODO())
	if err != nil {
		return "", err
	}
	message := types.NewMessage(
		wallet.account.PublicKey,
		[]types.Instruction{
			sysprog.Transfer(
				wallet.account.PublicKey,
				common.PublicKeyFromString(receiver),
				amount,
			),
		},
		response.Blockhash,
	)

	tx, err := types.NewTransaction(message, []types.Account{wallet.account, wallet.account})
	if err != nil {
		return "", err
	}

	txhash, err := wallet.c.SendTransaction2(context.TODO(), tx)
	if err != nil {
		return "", err
	}
	return txhash, nil
}
