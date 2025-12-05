package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"net"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	clrtoken "depin-proxy/contracts/CLRToken"
	clearnet "depin-proxy/contracts/ClearNet"
)

var DEFAULT_PRICE_PER_MINUTE = new(big.Int).Exp(big.NewInt(10), big.NewInt(16), nil) // 0.01 CLRToken
var DEFAULT_MIN_STAKE = new(big.Int).Exp(big.NewInt(10), big.NewInt(21), nil)        // 1000 CLRToken

func getNodeWallet() (*ecdsa.PrivateKey, common.Address) {
	walletPrivateKey, err := crypto.HexToECDSA(os.Getenv("WALLET_PRIVATE_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	walletPublicKey := walletPrivateKey.Public()
	publicKeyECDSA, ok := walletPublicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	walletAdress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return walletPrivateKey, walletAdress
}

func initContract() (*clrtoken.Clrtoken, *clearnet.Clearnet, *ethclient.Client, error) {
	endpoint := os.Getenv("INFURA_ENDPOINT")
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		return nil, nil, nil, err
	}

	clearnet_address := common.HexToAddress(os.Getenv("CLEARNET_CONTRACT_ADDRESS"))
	clearnet_instance, err := clearnet.NewClearnet(clearnet_address, client)
	if err != nil {
		return nil, nil, nil, err
	}

	clrtoken_address := common.HexToAddress(os.Getenv("CLRTOKEN_CONTRACT_ADDRESS"))
	clrtoken_instance, err := clrtoken.NewClrtoken(clrtoken_address, client)
	if err != nil {
		return nil, nil, nil, err
	}

	return clrtoken_instance, clearnet_instance, client, nil
}

func getAuth(client *ethclient.Client, walletPrivateKey *ecdsa.PrivateKey, walletAddress common.Address) *bind.TransactOpts {
	nonce, err := client.PendingNonceAt(context.Background(), walletAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainID := big.NewInt(11155111)
	auth, err := bind.NewKeyedTransactorWithChainID(walletPrivateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	return auth
}

func approveCLRTokenSpending(clrtoken_instance *clrtoken.Clrtoken, client *ethclient.Client, auth *bind.TransactOpts) (*types.Transaction, error) {
	minStake := DEFAULT_MIN_STAKE
	approveTx, err := clrtoken_instance.Approve(auth, common.HexToAddress(os.Getenv("CLEARNET_CONTRACT_ADDRESS")), minStake)
	if err != nil {
		return nil, fmt.Errorf("approve tx error: %w", err)
	}
	fmt.Printf("Approve tx sent: %s\n", approveTx.Hash().Hex())
	receipt, err := bind.WaitMined(context.Background(), client, approveTx)
	if err != nil {
		return nil, fmt.Errorf("waiting approve mined: %w", err)
	}
	if receipt.Status == types.ReceiptStatusSuccessful {
		fmt.Println("✅ Transaction successful!")
		return approveTx, nil
	} else {
		return nil, fmt.Errorf("approve tx failed, status=%v", receipt.Status)
	}
}

func registerNode(clearnet_instance *clearnet.Clearnet, client *ethclient.Client, auth *bind.TransactOpts) (*types.Transaction, error) {
	pricePerMinute := DEFAULT_PRICE_PER_MINUTE
	registerTx, err := clearnet_instance.RegisterNode(auth, ip.String(), uint16(port), pricePerMinute)
	if err != nil {
		return nil, fmt.Errorf("registerNode tx error: %w", err)
	}

	fmt.Printf("RegisterTx sent: %s\n", registerTx.Hash().Hex())
	registerReceipt, err := bind.WaitMined(context.Background(), client, registerTx)
	if err != nil {
		return nil, fmt.Errorf("waiting register mined: %w", err)
	}
	if registerReceipt.Status != types.ReceiptStatusSuccessful {
		return nil, fmt.Errorf("register tx failed, status=%v", registerReceipt.Status)
	}
	fmt.Println("✅ Node registered")
	return registerTx, nil
}

func deRegisterNode(clearnet_instance *clearnet.Clearnet, client *ethclient.Client, auth *bind.TransactOpts) (*types.Transaction, error) {
	deRegisterTx, err := clearnet_instance.DeregisterNode(auth)
	if err != nil {
		return nil, fmt.Errorf("DeRegisterNode tx error: %w", err)
	}

	fmt.Printf("DeRegisterTx sent: %s\n", deRegisterTx.Hash().Hex())
	deRegisterReceipt, err := bind.WaitMined(context.Background(), client, deRegisterTx)
	if err != nil {
		return nil, fmt.Errorf("waiting DeRegister mined: %w", err)
	}
	if deRegisterReceipt.Status != types.ReceiptStatusSuccessful {
		return nil, fmt.Errorf("DeRegister tx failed, status=%v", deRegisterReceipt.Status)
	}
	fmt.Println("✅ Node deregistered")
	return deRegisterTx, nil
}

func updateNodePrice(clearnet_instance *clearnet.Clearnet, client *ethclient.Client, auth *bind.TransactOpts, newPricePerMinute *big.Int) (*types.Transaction, error) {
	updateTx, err := clearnet_instance.UpdatePrice(auth, newPricePerMinute)
	if err != nil {
		return nil, fmt.Errorf("UpdateNode tx error: %w", err)
	}

	fmt.Printf("UpdateTx sent: %s\n", updateTx.Hash().Hex())
	updateReceipt, err := bind.WaitMined(context.Background(), client, updateTx)
	if err != nil {
		return nil, fmt.Errorf("waiting update mined: %w", err)
	}
	if updateReceipt.Status != types.ReceiptStatusSuccessful {
		return nil, fmt.Errorf("update tx failed, status=%v", updateReceipt.Status)
	}
	fmt.Println("✅ Node price updated")
	fmt.Println("New Price per Minute (in CLRToken): " + (new(big.Float).Quo(new(big.Float).SetInt(newPricePerMinute), new(big.Float).SetInt(big.NewInt(int64(1e18))))).String())
	return updateTx, nil
}

func updateNodeInfo(clearnet_instance *clearnet.Clearnet, client *ethclient.Client, auth *bind.TransactOpts, newIP string, newPort uint16) (*types.Transaction, error) {
	updateTx, err := clearnet_instance.UpdateNodeInfo(auth, newIP, newPort)
	if err != nil {
		return nil, fmt.Errorf("UpdateNodeInfo tx error: %w", err)
	}

	fmt.Printf("UpdateNodeInfo Tx sent: %s\n", updateTx.Hash().Hex())
	updateReceipt, err := bind.WaitMined(context.Background(), client, updateTx)
	if err != nil {
		return nil, fmt.Errorf("waiting updateNodeInfo mined: %w", err)
	}
	if updateReceipt.Status != types.ReceiptStatusSuccessful {
		return nil, fmt.Errorf("updateNodeInfo tx failed, status=%v", updateReceipt.Status)
	}
	fmt.Println("✅ Node info updated")
	fmt.Println("Ip: " + ip.String())
	fmt.Println("Port: " + fmt.Sprintf("%d", port))
	fmt.Println("Remember to update your firewall settings!")
	return updateTx, nil
}

func getNodeStatus(clearnet_instance *clearnet.Clearnet, walletAddress common.Address) (NodeInfo, error) {
	nodeInfo, err := clearnet_instance.GetNodeInfo(&bind.CallOpts{}, walletAddress)
	if err != nil {
		return NodeInfo{}, fmt.Errorf("GetNodeInfo call error: %w", err)
	}

	return NodeInfo{
		ip:                 net.ParseIP(nodeInfo.IpAddress),
		port:               nodeInfo.Port,
		pricePerMinute:     nodeInfo.PricePerMinute,
		reputationScore:    nodeInfo.ReputationScore,
		totalMinutesServed: nodeInfo.TotalMinutesServed,
		totalEarnings:      nodeInfo.TotalEarnings,
	}, nil
}

func printNodeStatus(nodeInfo NodeInfo) {
	fmt.Println("Node Status:")
	fmt.Println("IP Address: ", nodeInfo.ip)
	fmt.Println("Port: ", nodeInfo.port)
	fmt.Println("Price Per Minute (in CLRToken): ", new(big.Float).Quo(new(big.Float).SetInt(nodeInfo.pricePerMinute), new(big.Float).SetInt(big.NewInt(int64(1e18)))).String())
	fmt.Println("Reputation Score: ", nodeInfo.reputationScore.String())
	fmt.Println("Total Minutes Served: ", nodeInfo.totalMinutesServed.String())
	fmt.Println("Total Earnings (in CLRToken): ", new(big.Float).Quo(new(big.Float).SetInt(nodeInfo.totalEarnings), new(big.Float).SetInt(big.NewInt(int64(1e18)))).String())
}

func getPaymentChannelInfo(clearnet_instance *clearnet.Clearnet, clientAddress common.Address) (PaymentChannel, error) {
	channelInfo, err := clearnet_instance.GetPaymentChannelInfo(&bind.CallOpts{}, clientAddress)
	if err != nil {
		return PaymentChannel{}, fmt.Errorf("GetPaymentChannelInfo call error: %w", err)
	}

	return PaymentChannel{
		balance:  channelInfo.Balance,
		nonce:    channelInfo.Nonce,
		isActive: channelInfo.IsActive,
	}, nil
}
