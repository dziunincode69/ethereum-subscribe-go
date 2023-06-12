package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

var (
	ctx         = context.Background()
	url         = "ws://127.0.0.1:8546"
	rpcClient   *rpc.Client
	err         error
	gethclients *gethclient.Client
	txChan      = make(chan *types.Transaction)
)

func main() {
	connect := ConnectToNetwork()
	if connect {
		gethclients.SubscribeFullPendingTransactions(ctx, txChan)
		for {
			txs := <-txChan
			input := hex.EncodeToString(txs.Data())
			to := txs.To()
			gasPrice := txs.GasPrice()
			froms, _ := types.Sender(types.LatestSignerForChainID(big.NewInt(1)), txs)
			from := froms
			fmt.Println(input, to, gasPrice, from)
		}

	}
}

func ConnectToNetwork() bool {
	rpcClient, err = rpc.DialContext(ctx, url)
	if err != nil {
		log.Fatal(err)
	}
	gethclients = gethclient.New(rpcClient)
	return true
}
