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
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"

	clrtoken "depin-proxy/contracts/CLRToken"
	clearnet "depin-proxy/contracts/ClearNet"
)

type Connection struct {
	walletAddress        string
	vpnClientPublicKey   string
	ip                   net.IP
	receivedNonce        *big.Int
	connectionStartTime  *big.Int
	agreedPricePerMinute *big.Int
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

type ProtectedConnections struct {
	mu          sync.RWMutex
	connections []Connection
}

var protectedConnections ProtectedConnections = ProtectedConnections{
	connections: make([]Connection, 0),
}

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
			fmt.Println("Current Connections:", protectedConnections.getConnectionsCount())
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
	go disconnectWatcher()

	http.HandleFunc("/connect", connectionHandler)
	// http.HandleFunc("/disconnect", disconectHandler)

	fmt.Println("LocalIP:", ip.String())
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

// func generateIV() string {
// 	key, err := crypto.GenerateKey()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	iv := crypto.FromECDSA(key)
// 	ivs = append(ivs, hexutil.Encode(iv))
// 	fmt.Println("Generated IV:", hexutil.Encode(iv))
// 	return hexutil.Encode(iv)
// }

// func isValidIV(iv string) bool {
// 	for i, v := range ivs {
// 		if v == iv {
// 			fmt.Println("Valid IV:", iv)
// 			ivs = append(ivs[:i], ivs[i+1:]...)
// 			return true
// 		}
// 	}
// 	return false
// }

func isValidNonce(nonce *big.Int, clientAddress common.Address) bool {
	currentNonce, err := getCurrentNonce(clearnet_instance, clientAddress)
	if err != nil {
		fmt.Println("Error getting current nonce:", err)
		return false
	}
	return nonce.Cmp(new(big.Int).Add(currentNonce, big.NewInt(1))) == 0
}

func disconnectWatcher() {
	for {
		time.Sleep(60 * time.Second)
		currentTime := big.NewInt(time.Now().Unix())
		for i := 0; i < protectedConnections.getConnectionsCount(); i++ {
			conn := protectedConnections.getConnection(i)
			elapsedMinutes := new(big.Int).Div(new(big.Int).Sub(currentTime, conn.connectionStartTime), big.NewInt(60))
			totalCost := new(big.Int).Mul(elapsedMinutes, conn.agreedPricePerMinute)
			paymentChannel, err := getPaymentChannelInfo(clearnet_instance, common.HexToAddress(conn.walletAddress))
			if err != nil {
				fmt.Println("Error getting payment channel info:", err)
				continue
			}

			if conn.receivedNonce.Cmp(new(big.Int).Add(paymentChannel.nonce, big.NewInt(1))) != 0 {
				err = removeWireguardPeer(conn.vpnClientPublicKey)
				if err != nil {
					fmt.Println("Error removing Wireguard peer:", err)
					continue
				}
				protectedConnections.removeConnection(conn.walletAddress)
				fmt.Println("Client nonce increased")
				fmt.Println("Connection count:", protectedConnections.getConnectionsCount())
				i--
				continue
			}

			if paymentChannel.balance.Cmp(totalCost) == -1 {
				fmt.Println("Disconnecting wallet address due to insufficient balance:", conn.walletAddress)
				err = removeWireguardPeer(conn.vpnClientPublicKey)
				if err != nil {
					fmt.Println("Error removing Wireguard peer:", err)
					continue
				}
				protectedConnections.removeConnection(conn.walletAddress)
				fmt.Println("Connection count:", protectedConnections.getConnectionsCount())
				i--
			}
		}
	}
}

// func disconectHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	body, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		http.Error(w, "Error reading request body", http.StatusInternalServerError)
// 		return
// 	}

// 	bodyStr := string(body)
// 	s := strings.Split(bodyStr, "\n")
// 	nonce, ok := new(big.Int).SetString(s[0], 10)
// 	if !ok {
// 		http.Error(w, "Invalid nonce format", http.StatusBadRequest)
// 		return
// 	}
// 	signature := s[1]
// 	recoveredAddr, err := verifySignature(nonce.String(), signature)
// 	if err != nil {
// 		http.Error(w, "Error verifying signature: "+err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	if !isValidNonce(nonce, common.HexToAddress(recoveredAddr)) {
// 		http.Error(w, "Invalid or reused nonce", http.StatusBadRequest)
// 		return
// 	}
// 	fmt.Println("Disconnecting Wallet Address: ", recoveredAddr)

// 	connection := protectedConnections.getConnectionByWalletAddress(recoveredAddr)
// 	if connection == nil {
// 		http.Error(w, "No active connection found for this wallet address", http.StatusBadRequest)
// 		return
// 	}

// 	err = removeWireguardPeer(connection.vpnClientPublicKey)
// 	if err != nil {
// 		http.Error(w, "Error removing Wireguard peer: "+err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	protectedConnections.removeConnection(recoveredAddr)

// 	fmt.Println("Connection count:", protectedConnections.getConnectionsCount())

// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("Disconnected"))
// }

func connectionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
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
	nonce, ok := new(big.Int).SetString(s[0], 10)
	if !ok {
		http.Error(w, "Invalid nonce format", http.StatusBadRequest)
		return
	}
	vpnClientPublicKey := s[1]
	connectionStartTime, ok := new(big.Int).SetString(s[2], 10)
	if !ok {
		http.Error(w, "Invalid connection start time format", http.StatusBadRequest)
		return
	}
	currentTime := big.NewInt(time.Now().Unix())
	if currentTime.Cmp(connectionStartTime) == -1 {
		http.Error(w, "Connection start time is in the future", http.StatusBadRequest)
		return
	}

	pricePerMinute, ok := new(big.Int).SetString(s[3], 10)
	if !ok {
		http.Error(w, "Invalid price per minute format", http.StatusBadRequest)
		return
	}
	signature := s[4]
	fmt.Println("VPN Client Public Key:", vpnClientPublicKey)
	fmt.Println("Signature:", signature)

	recoveredAddr, err := verifySignature(nonce.String()+vpnClientPublicKey+connectionStartTime.String()+pricePerMinute.String(), signature)
	if err != nil {
		http.Error(w, "Error verifying signature: "+err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("Wallet Address:", recoveredAddr)
	if !isValidNonce(nonce, common.HexToAddress(recoveredAddr)) {
		http.Error(w, "Invalid or reused nonce", http.StatusBadRequest)
		return
	}
	if protectedConnections.connectionExists(recoveredAddr) {
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
	wireguardClientIP := protectedConnections.addConnection(recoveredAddr, vpnClientPublicKey, nonce, connectionStartTime, pricePerMinute)

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

	nodeSignature, err := signMessage(nonce.String() + connectionStartTime.String() + pricePerMinute.String())
	if err != nil {
		http.Error(w, "Error signing message: "+err.Error(), http.StatusInternalServerError)
		return
	}

	responseDataJson, err := json.Marshal(struct {
		Message                  string
		WireguardClientCIDR      string
		WireguardServerPublicKey string
		WireguardPort            int
		WireguardDNS             string
		NodeSignature            string
	}{
		Message:                  "Connection accepted",
		WireguardClientCIDR:      getCIDRFromIP(wireguardClientIP),
		WireguardServerPublicKey: wireguardPublicKey,
		WireguardPort:            wireguardPort,
		WireguardDNS:             "1.1.1.1, 8.8.8.8",
		NodeSignature:            nodeSignature,
	})
	if err != nil {
		http.Error(w, "Error generating response: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(responseDataJson)
}

func signMessage(message string) (string, error) {
	prefixedMessage := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)
	messageHash := crypto.Keccak256Hash([]byte(prefixedMessage))
	signature, err := crypto.Sign(messageHash.Bytes(), walletPrivateKey)
	if err != nil {
		return "", err
	}
	return hexutil.Encode(signature), nil
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

func _connectionExists(walletAddress string, connections []Connection) bool {
	for _, conn := range connections {
		if conn.walletAddress == walletAddress {
			return true
		}
	}
	return false
}

func (c *ProtectedConnections) connectionExists(walletAddress string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return _connectionExists(walletAddress, c.connections)
}

func (c *ProtectedConnections) getConnectionByWalletAddress(walletAddress string) *Connection {
	c.mu.RLock()
	defer c.mu.RUnlock()
	for _, conn := range c.connections {
		if conn.walletAddress == walletAddress {
			return &conn
		}
	}
	return nil
}

func (c *ProtectedConnections) getConnection(i int) Connection {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.connections[i]
}

func (c *ProtectedConnections) getConnectionsCount() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.connections)
}

func (c *ProtectedConnections) addConnection(walletAddress string, vpnClientPublicKey string, receivedNonce *big.Int, connectionStartTime *big.Int, pricePerMinute *big.Int) net.IP {
	c.mu.Lock()
	defer c.mu.Unlock()
	if !_connectionExists(walletAddress, c.connections) {
		ip := allocateIP()
		connection := Connection{
			walletAddress:        walletAddress,
			vpnClientPublicKey:   vpnClientPublicKey,
			ip:                   ip,
			receivedNonce:        receivedNonce,
			connectionStartTime:  connectionStartTime,
			agreedPricePerMinute: pricePerMinute,
		}
		c.connections = append(c.connections, connection)
		fmt.Println("Connection count:", len(c.connections))
		return ip
	}
	fmt.Println("Connection already exists for wallet address:", walletAddress)
	return nil
}

func (c *ProtectedConnections) removeConnection(recoveredAddr string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for i, conn := range c.connections {
		if conn.walletAddress == recoveredAddr {
			releaseIP(conn.ip)
			c.connections = append(c.connections[:i], c.connections[i+1:]...)
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
