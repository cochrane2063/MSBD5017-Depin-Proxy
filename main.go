package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"io"
	"log"
	"math/big"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"

	clearnet "depin-proxy/contracts"
)

type Connection struct {
	walletAddress      string
	vpnClientPublicKey string
}

var ip net.IP

var port int

var ivs []string

var connections []Connection

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

func initContract() (*clearnet.Clearnet, *ethclient.Client, error) {
	endpoint := os.Getenv("INFURA_ENDPOINT")
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		return nil, nil, err
	}

	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		return nil, nil, err
	}

	address := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
	instance, err := clearnet.NewClearnet(address, client)
	if err != nil {
		return nil, nil, err
	}

	fmt.Println("Latest Block Number:", blockNumber)
	return instance, client, nil
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

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
	ip = GetLocalIP()
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("Invalid PORT value:" + err.Error())
	}

	instance, client, err := initContract()
	if err != nil {
		log.Fatal(err)
	}
	walletPrivateKey, walletAddress := getNodeWallet()
	auth := getAuth(client, walletPrivateKey, walletAddress)

	// _, _ = instance, auth
	tx, err := instance.RegisterNode(auth, ip.String(), uint16(port), big.NewInt(1000000000000000000))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())

	registedNodes, err := instance.GetActiveNodes(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Registered Nodes:", registedNodes)

	http.HandleFunc("/connect", connectionHandler)
	http.HandleFunc("/disconnect", disconectHandler)

	fmt.Println("LocalIP:", ip.String())
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

func generateIV() string {
	key, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	iv := crypto.FromECDSA(key)
	ivs = append(ivs, hexutil.Encode(iv))
	fmt.Println("Generated IV:", hexutil.Encode(iv))
	return hexutil.Encode(iv)
}

func isValidIV(iv string) bool {
	for i, v := range ivs {
		if v == iv {
			fmt.Println("Valid IV:", iv)
			ivs = append(ivs[:i], ivs[i+1:]...)
			return true
		}
	}
	return false
}

func disconectHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodGet {
		iv := generateIV()
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(iv))
		return
	}
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
	iv := s[0]
	signature := s[1]

	if !isValidIV(iv) {
		http.Error(w, "Invalid or reused IV", http.StatusBadRequest)
		return
	}
	recoveredAddr, err := verifySignature(iv, signature)
	if err != nil {
		http.Error(w, "Error verifying signature: "+err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("Disconnecting Wallet Address: ", recoveredAddr)

	removeConnection(recoveredAddr)

	fmt.Println("Connection count:", len(connections))

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Disconnected"))
}

func connectionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodGet {
		iv := generateIV()
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(iv))
		return
	}
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
	iv := s[0]
	vpnClientPublicKey := s[1]
	signature := s[2]
	fmt.Println("VPN Client Public Key:", vpnClientPublicKey)
	fmt.Println("Signature:", signature)

	// Verify signature
	if !isValidIV(iv) {
		http.Error(w, "Invalid or reused IV", http.StatusBadRequest)
		return
	}
	recoveredAddr, err := verifySignature(iv+vpnClientPublicKey, signature)
	if err != nil {
		http.Error(w, "Error verifying signature: "+err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("Wallet Address:", recoveredAddr)

	addConnection(Connection{
		walletAddress:      recoveredAddr,
		vpnClientPublicKey: vpnClientPublicKey,
	})
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Connection accepted"))

}

func verifySignature(message, sigHex string) (string, error) {
	// decode signature
	sig, err := hexutil.Decode(sigHex)
	if err != nil {
		return "", err
	}
	if len(sig) != 65 {
		return "", fmt.Errorf("invalid signature length: %d", len(sig))
	}

	// adjust V if needed (MetaMask returns 27/28)
	if sig[64] >= 27 {
		sig[64] -= 27
	}

	// recreate prefixed message
	prefixed := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)
	hash := crypto.Keccak256Hash([]byte(prefixed))

	// recover public key
	pubKey, err := crypto.SigToPub(hash.Bytes(), sig)
	if err != nil {
		return "", err
	}

	// derive address and compare (case-insensitive)
	recoveredAddr := crypto.PubkeyToAddress(*pubKey).Hex()
	return recoveredAddr, nil
}

func connectionExists(vpnClientPublicKey string) bool {
	for _, conn := range connections {
		if conn.vpnClientPublicKey == vpnClientPublicKey {
			return true
		}
	}
	return false
}

func addConnection(connection Connection) {
	if !connectionExists(connection.vpnClientPublicKey) {
		connections = append(connections, connection)
	}
	fmt.Println("Connection count:", len(connections))
}

func removeConnection(recoveredAddr string) {
	for i, conn := range connections {
		if conn.walletAddress == recoveredAddr {
			connections = append(connections[:i], connections[i+1:]...)
			break
		}
	}
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
