# Welcome to CLEARNET Node!

Get some rewards for sharing your extra bandwith

## Getting Started
You can use docker compose to run the server node.

Make sure you have docker and docker compose installed.

A docker compose file with private key are privided. Please use the private key included in the compose file. Because the server node requires sepolia ether and CLR Tokens to register. If you are using a private key with insufficient eth or CLT Tokens, the registration will fail.

Download the docker-compose.yml file from this repo and cd into the directory with the docker-compose.yml file.

Run
```bash
docker compose up -d
```
If you are running an older version of docker compose, you should use
```bash
docker-compose up -d
```

Make sure you have the required ports open in your firewall!

Set up port forwarding if necessary.

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
# Please use the key provided for testing since it already has ether and CLR Tokens.
export WALLET_PRIVATE_KEY="f0ce4f9c5521bfb662312a18794928ba3c24ac8afa55c4687d894452818faaeb"
export INFURA_ENDPOINT="https://sepolia.infura.io/v3/8c1eb7cd8a5a43c9bc5c3d69db182b83"
export CHAIN_ID="11155111"
export CLEARNET_CONTRACT_ADDRESS="0x0305e95225f65db13e98c775dbb95b98178ae73b"
export CLRTOKEN_CONTRACT_ADDRESS="0xf1664c17887767c8f58695846babb349ca61d2e9"
export PROXY_PORT="8080"
# export VPN_PORT="51820" #optional: by default it will be the port your wireguard server is listening on 
sudo /path/to/file/clearnetNode
```

If the command faileds, you need to give it execute permission. Run
```bash
sudo chmod +x /path/to/file/clearnetNode
```
