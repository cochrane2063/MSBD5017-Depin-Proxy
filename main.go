package main

import (
	"bufio"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"

	clrtoken "depin-proxy/contracts/CLRToken"
	clearnet "depin-proxy/contracts/ClearNet"
)

type Connection struct {
	walletAddress      string
	vpnClientPublicKey string
	ip                 net.IP
}

type NodeInfo struct {
	ip                 net.IP
	port               uint16
	pricePerMinute     *big.Int
	reputationScore    *big.Int
	totalMinutesServed *big.Int
	totalEarnings      *big.Int
}

type PaymentChannel struct {
	balance  *big.Int
	nonce    *big.Int
	isActive bool
}

var ip net.IP
var port int

var clrtoken_instance *clrtoken.Clrtoken
var clearnet_instance *clearnet.Clearnet
var client *ethclient.Client
var walletPrivateKey *ecdsa.PrivateKey
var walletAddress common.Address

var ivs []string
var connections []Connection

func dialogue() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to the CLEARNET VPN Node!")
	fmt.Println("Available commands:")
	fmt.Println("  register     - Register this node on ClearNet")
	fmt.Println("  deregister   - Deregister this node from ClearNet")
	fmt.Println("  update price - Update this node's price per minute")
	fmt.Println("  status       - Show current status")
	fmt.Println("  help         - Show available commands")
	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Println("stdin read error:", err)
			return
		}
		cmd := strings.TrimSpace(line)
		switch cmd {
		case "help":
			fmt.Println("Available commands:")
			fmt.Println("  register     - Register this node on ClearNet")
			fmt.Println("  deregister   - Deregister this node from ClearNet")
			fmt.Println("  update price - Update this node's price per minute")
			fmt.Println("  status       - Show current status")
			fmt.Println("  help         - Show available commands")
		case "register":
			fmt.Println("Registering")
			approveCLRTokenSpending(clrtoken_instance, client, getAuth(client, walletPrivateKey, walletAddress))
			registerNode(clearnet_instance, client, getAuth(client, walletPrivateKey, walletAddress))
		case "deregister":
			fmt.Println("Deregistering")
			deRegisterNode(clearnet_instance, client, getAuth(client, walletPrivateKey, walletAddress))
		case "update price":
			fmt.Println("Please enter the new price per minute in CLRToken (e.g., 0.01): ")
			line, err := reader.ReadString('\n')
			if err != nil {
				log.Println("stdin read error:", err)
				return
			}
			priceFloat, err := strconv.ParseFloat(strings.TrimSpace(line), 64)
			if err != nil {
				fmt.Println("Invalid price value:" + err.Error())
				continue
			}
			newPrice := new(big.Int).Mul(big.NewInt(int64(priceFloat*1e18)), big.NewInt(1))
			updateNodePrice(clearnet_instance, client, getAuth(client, walletPrivateKey, walletAddress), newPrice)
		case "status":
			fmt.Println("Connection Status:")
			fmt.Println("Current Connections:", len(connections))
			nodeInfo, err := getNodeStatus(clearnet_instance, walletAddress)
			if err != nil {
				fmt.Println("Error getting node status:", err)
				continue
			}
			printNodeStatus(nodeInfo)
		default:
			fmt.Println("unknown command:", cmd)
		}
	}
}

func updateChanges() {
	nodeInfo, err := getNodeStatus(clearnet_instance, walletAddress)
	if err != nil {
		fmt.Println("Node not registered")
		return
	}
	if !nodeInfo.ip.Equal(ip) || nodeInfo.port != uint16(port) {
		updateNodeInfo(clearnet_instance, client, getAuth(client, walletPrivateKey, walletAddress), ip.String(), uint16(port))
	}
}

func initialize() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
	ip = GetLocalIP()
	port, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("Invalid PORT value:" + err.Error())
	}
	initializeIPPool()
	err = clearWireguardPeers()
	if err != nil {
		log.Fatal("Failed to clear Wireguard peers: " + err.Error())
	}

	clrtoken_instance, clearnet_instance, client, err = initContract()
	if err != nil {
		log.Fatal(err)
	}
	walletPrivateKey, walletAddress = getNodeWallet()
}

func main() {
	initialize()
	updateChanges()

	go dialogue()

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

	connection := getConnectionByWalletAddress(recoveredAddr)
	if connection == nil {
		http.Error(w, "No active connection found for this wallet address", http.StatusBadRequest)
		return
	}

	err = removeWireguardPeer(connection.vpnClientPublicKey)
	if err != nil {
		http.Error(w, "Error removing Wireguard peer: "+err.Error(), http.StatusInternalServerError)
		return
	}

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

	if connectionExists(recoveredAddr) {
		http.Error(w, "Connection already exists", http.StatusBadRequest)
		return
	}

	paymentChannel, err := getPaymentChannelInfo(clearnet_instance, common.HexToAddress(recoveredAddr))
	if err != nil {
		http.Error(w, "Error getting payment channel info: "+err.Error(), http.StatusBadRequest)
		return
	}
	if !paymentChannel.isActive {
		http.Error(w, "Client not registered", http.StatusBadRequest)
		return
	}
	wireguardClientIP := addConnection(recoveredAddr, vpnClientPublicKey)

	wireguardPublicKey, err := getWireguardPublicKey()
	if err != nil {
		fmt.Println("Error getting Wireguard public key:", err)
	}

	wireguardPort, err := getWireguardPort()
	if err != nil {
		fmt.Println("Error getting Wireguard port:", err)
	}

	err = addWireguardPeer(wireguardClientIP, vpnClientPublicKey)
	if err != nil {
		http.Error(w, "Error adding Wireguard peer: "+err.Error(), http.StatusInternalServerError)
		return
	}

	responseDataJson, err := json.Marshal(struct {
		Message                  string
		WireguardClientCIDR      string
		WireguardServerPublicKey string
		WireguardPort            int
		WireguardDNS             string
	}{
		Message:                  "Connection accepted",
		WireguardClientCIDR:      getCIDRFromIP(wireguardClientIP),
		WireguardServerPublicKey: wireguardPublicKey,
		WireguardPort:            wireguardPort,
		WireguardDNS:             "1.1.1.1, 8.8.8.8",
	})
	if err != nil {
		http.Error(w, "Error generating response: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(responseDataJson)
}

func verifySignature(message, sigHex string) (string, error) {
	sig, err := hexutil.Decode(sigHex)
	if err != nil {
		return "", err
	}
	if len(sig) != 65 {
		return "", fmt.Errorf("invalid signature length: %d", len(sig))
	}

	if sig[64] >= 27 {
		sig[64] -= 27
	}

	prefixed := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)
	hash := crypto.Keccak256Hash([]byte(prefixed))

	pubKey, err := crypto.SigToPub(hash.Bytes(), sig)
	if err != nil {
		return "", err
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey).Hex()
	return recoveredAddr, nil
}

func connectionExists(walletAddress string) bool {
	for _, conn := range connections {
		if conn.walletAddress == walletAddress {
			return true
		}
	}
	return false
}

func getConnectionByWalletAddress(walletAddress string) *Connection {
	for _, conn := range connections {
		if conn.walletAddress == walletAddress {
			return &conn
		}
	}
	return nil
}

func addConnection(walletAddress string, vpnClientPublicKey string) net.IP {
	if !connectionExists(walletAddress) {
		ip := allocateIP()
		connection := Connection{
			walletAddress:      walletAddress,
			vpnClientPublicKey: vpnClientPublicKey,
			ip:                 ip,
		}
		connections = append(connections, connection)
		fmt.Println("Connection count:", len(connections))
		return ip
	}
	fmt.Println("Connection already exists for wallet address:", walletAddress)
	return nil
}

func removeConnection(recoveredAddr string) {
	for i, conn := range connections {
		if conn.walletAddress == recoveredAddr {
			releaseIP(conn.ip)
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
