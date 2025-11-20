package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/big"
	"net"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	endpoint := "https://sepolia.infura.io/v3/8c1eb7cd8a5a43c9bc5c3d69db182b83"
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		log.Fatal(err)
	}

	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("") // TODO: fill in contract address
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	callUsingBindings(client, address)

	http.HandleFunc("/connect", connectionHandler)

	fmt.Println("Latest Block Number:", blockNumber)
	fmt.Println("LocalIP:", GetLocalIP())
	http.ListenAndServe(":8080", nil)
}

func callUsingBindings(client *ethclient.Client, contractAddr common.Address) {
	ctx := context.Background()

	contract, err := MyContract(contractAddr, client)
	if err != nil {
		log.Fatal(err)
	}

	privateKeyHex := "" // TODO
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatal(err)
	}
	chainID := big.NewInt(11155111) // TODO: Sepolia chain ID
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}

}

func connectionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	bodyStr := string(body)
	s := strings.Split(bodyStr, "\n")
	signature := s[0]
	publicKey := s[1]
	payload := strings.Split(s[2], " ")

	walletAddr = payload[0]
	wireguardPubKey = payload[1]

	// Verify signature

}

func GetLocalIP() net.IP {
	res, err := http.Get("https://api.ipify.org")
	if err != nil {
		conn, err := net.Dial("udp", "1.1.1.1:80")
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		localAddress := conn.LocalAddr().(*net.UDPAddr)
		return localAddress.IP
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	ip := net.ParseIP(strings.TrimSpace(string(body)))
	return ip
}
