package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"

	"golang.zx2c4.com/wireguard/wgctrl"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

var DEFAULT_WIREGUARD_DEVICE = "wg0"

type ProtectedArray struct {
	mu    sync.Mutex
	array []bool
}

func (pa *ProtectedArray) setValue(index int, value bool) {
	pa.mu.Lock()
	defer pa.mu.Unlock()
	pa.array[index] = value
}

func (pa *ProtectedArray) getValue(index int) bool {
	pa.mu.Lock()
	defer pa.mu.Unlock()
	return pa.array[index]
}

var ipAvailable = ProtectedArray{
	array: make([]bool, 256),
}

func initializeIPPool() {
	for i := 1; i < 255; i++ {
		ipAvailable.setValue(i, true)
	}
	ipAvailable.setValue(0, false)
	ipAvailable.setValue(1, false)
	ipAvailable.setValue(255, false)
}

func allocateIP() net.IP {
	for i := 2; i < 255; i++ {
		if ipAvailable.getValue(i) {
			ipAvailable.setValue(i, false)
			return net.IPv4(10, 8, 0, byte(i))
		}
	}
	return nil
}

func releaseIP(ip net.IP) {
	fmt.Printf("Releasing IP: %s\n", ip.String())
	lastOctet := int(ip[len(ip)-1])
	ipAvailable.setValue(lastOctet, true)
}

func getCIDRFromIP(ip net.IP) string {
	return ip.String() + "/24"
}

func getWireguardPort() (int, error) {
	env_vpn_port, exists := os.LookupEnv("VPN_PORT")
	if exists {
		port, err := strconv.Atoi(env_vpn_port)
		if err == nil {
			return port, nil
		}
	}
	wgClient, err := wgctrl.New()
	if err != nil {
		return 0, fmt.Errorf("failed to create wgctrl client: %w", err)
	}
	defer wgClient.Close()

	device, err := wgClient.Device(DEFAULT_WIREGUARD_DEVICE)
	if err != nil {
		return 0, fmt.Errorf("failed to get Wireguard device: %w", err)
	}

	return device.ListenPort, nil
}

func getWireguardPublicKey() (string, error) {
	wgClient, err := wgctrl.New()
	if err != nil {
		return "", fmt.Errorf("failed to create wgctrl client: %w", err)
	}
	defer wgClient.Close()

	device, err := wgClient.Device(DEFAULT_WIREGUARD_DEVICE)
	if err != nil {
		return "", fmt.Errorf("failed to get Wireguard device: %w", err)
	}

	publicKey := device.PublicKey.String()
	return publicKey, nil
}

func addWireguardPeer(clientIP net.IP, peerPublicKeyString string) error {
	wgClient, err := wgctrl.New()
	if err != nil {
		return fmt.Errorf("failed to create wgctrl client: %w", err)
	}
	defer wgClient.Close()

	peerPubKey, err := wgtypes.ParseKey(peerPublicKeyString)
	if err != nil {
		return fmt.Errorf("failed to parse peer public key: %w", err)
	}
	var peerConfigs []wgtypes.PeerConfig
	peerConfigs = append(peerConfigs, wgtypes.PeerConfig{
		PublicKey:                   peerPubKey,
		Remove:                      false,
		UpdateOnly:                  false,
		PresharedKey:                nil,
		Endpoint:                    nil,
		PersistentKeepaliveInterval: nil,
		ReplaceAllowedIPs:           true,
		AllowedIPs:                  []net.IPNet{{IP: clientIP, Mask: net.CIDRMask(32, 32)}},
	})

	config := wgtypes.Config{
		PrivateKey:   nil,
		ListenPort:   nil,
		FirewallMark: nil,
		ReplacePeers: false,
		Peers:        peerConfigs,
	}

	err = wgClient.ConfigureDevice(DEFAULT_WIREGUARD_DEVICE, config)

	if err != nil {
		return fmt.Errorf("failed to add Wireguard peer: %w", err)
	}

	return nil
}

func removeWireguardPeer(peerPublicKeyString string) error {
	wgClient, err := wgctrl.New()
	if err != nil {
		return fmt.Errorf("failed to create wgctrl client: %w", err)
	}
	defer wgClient.Close()

	peerPubKey, err := wgtypes.ParseKey(peerPublicKeyString)
	if err != nil {
		return fmt.Errorf("failed to parse peer public key: %w", err)
	}
	var peerConfigs []wgtypes.PeerConfig
	peerConfigs = append(peerConfigs, wgtypes.PeerConfig{
		PublicKey:                   peerPubKey,
		Remove:                      true,
		UpdateOnly:                  true,
		PresharedKey:                nil,
		Endpoint:                    nil,
		PersistentKeepaliveInterval: nil,
		ReplaceAllowedIPs:           false,
		AllowedIPs:                  nil,
	})

	config := wgtypes.Config{
		PrivateKey:   nil,
		ListenPort:   nil,
		FirewallMark: nil,
		ReplacePeers: false,
		Peers:        peerConfigs,
	}

	err = wgClient.ConfigureDevice(DEFAULT_WIREGUARD_DEVICE, config)

	if err != nil {
		return fmt.Errorf("failed to remove Wireguard peer: %w", err)
	}

	return nil
}

func clearWireguardPeers() error {
	wgClient, err := wgctrl.New()
	if err != nil {
		return fmt.Errorf("failed to create wgctrl client: %w", err)
	}
	defer wgClient.Close()

	var peerConfigs []wgtypes.PeerConfig

	config := wgtypes.Config{
		PrivateKey:   nil,
		ListenPort:   nil,
		FirewallMark: nil,
		ReplacePeers: true,
		Peers:        peerConfigs,
	}

	err = wgClient.ConfigureDevice(DEFAULT_WIREGUARD_DEVICE, config)

	if err != nil {
		return fmt.Errorf("failed to clear Wireguard peers: %w", err)
	}

	return nil
}

// func getTrafficMetric() int {
// 	var activeConnections int = 0
// 	for i := 2; i < 255; i++ {
// 		if !ipAvailable.getValue(i) {
// 			activeConnections++
// 		}
// 	}
// 	var metric = activeConnections
// 	return metric
// }
