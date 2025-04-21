package utils

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func RewardUser(wallet string) error {
	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		return fmt.Errorf("failed to connect to the Ethereum client: %v", err)
	}

	abiFile, err := os.ReadFile("erc20_abi.json")
	if err != nil {
		return fmt.Errorf("failed to read ABI file: %v", err)
	}

	parsedABI, err := abi.JSON(strings.NewReader(string(abiFile)))
	if err != nil {
		return fmt.Errorf("failed to parse ABI: %v", err)
	}

	contractAddr := common.HexToAddress("0x04fe36e3dcb026c68d3688a231abcc68c46b2b64")
	contract := bind.NewBoundContract(contractAddr, parsedABI, client, client, client)

	privateKey, err := crypto.HexToECDSA("0xaf55ec40390874ec03bb95317af6dde825cd150eb6f25d0cccd67eee8f9499c1")
	if err != nil {
		return fmt.Errorf("failed to load private key: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))
	if err != nil {
		return fmt.Errorf("failed to create transactor: %v", err)
	}

	auth.GasLimit = uint64(300000)
	auth.GasPrice = big.NewInt(2000000000)

	to := common.HexToAddress(wallet)
	amount := new(big.Int)
	amount.SetString("10000000000000000000", 10)

	tx, err := contract.Transact(auth, "reward", to, amount)
	if err != nil {
		return fmt.Errorf("failed to send reward transaction: %v", err)
	}

	log.Printf("Reward sent! Tx: %s\n", tx.Hash().Hex())

	ctx := context.Background()
	receipt, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		return fmt.Errorf("failed to wait for transaction to be mined: %v", err)
	}

	if receipt.Status == 1 {
		log.Printf("Transaction successful! Block number: %d\n", receipt.BlockNumber)
	} else {
		log.Printf("Transaction failed! Block number: %d\n", receipt.BlockNumber)
	}

	return nil
}
