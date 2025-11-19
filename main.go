package main

import (
	// "context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
	// "github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// endpoint := "https://sepolia.infura.io/v3/8c1eb7cd8a5a43c9bc5c3d69db182b83"
	// client, err := ethclient.Dial(endpoint)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// blockNumber, err := client.BlockNumber(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Latest Block Number:", blockNumber)
	fmt.Println("LocalIP:", GetLocalIP())
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
