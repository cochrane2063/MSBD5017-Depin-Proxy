# Welcome to CLEARNET Node!

Get some rewards for sharing your extra bandwith

## Getting Started


## Run binary file directly
You can download the binary file from the releases section.

This project is ready for linux only, other os support comming soon...

Make sure you have wireguard setup on this system

For instructions to setup wireguard on ubuntu, please follow 

[https://www.digitalocean.com/community/tutorials/how-to-set-up-wireguard-on-ubuntu-22-04](https://www.digitalocean.com/community/tutorials/how-to-set-up-wireguard-on-ubuntu-22-04)

Other distros are very similar

### compile your own binary
If you want to compile your own binary file, make sure you have go installed.

Clone this repository and run inside the project root directory
```bash
go build -o clearnetNode
```

### Run binary file
Make sure to run the binary file with root. Because the proxy needs to access wireguard commands.

```bash
# The keys are provided for testing purposes
export WALLET_PRIVATE_KEY="a671613a49bc7d7f07e9d365fd95705704cf17b787f10ac7f4c3b2de1da319d5"
export INFURA_ENDPOINT="https://sepolia.infura.io/v3/8c1eb7cd8a5a43c9bc5c3d69db182b83"
export CHAIN_ID="11155111"
export CLEARNET_CONTRACT_ADDRESS="0x0305e95225f65db13e98c775dbb95b98178ae73b"
export CLRTOKEN_CONTRACT_ADDRESS="0xf1664c17887767c8f58695846babb349ca61d2e9"
export PORT="8080"
# export VPN_PORT="51820" #optional: by default it will be the port your wireguard server is listening on 
sudo /path/to/file/clearnetNode
```

If the command faileds, you need to give it execute permission. Run
```bash
sudo chmod +x /path/to/file/clearnetNode
```