package ecomm

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

type Erc20Client struct {
	contract   *gateway.Contract
	token_name string
}

// const (
// 	walletInfoDir = "/wallet"
// )

func NewErc20Client(token_name string) *Erc20Client {
	//log.Println("============ application-golang starts ============")

	err := os.Setenv("DISCOVERY_AS_LOCALHOST", "true")
	if err != nil {
		log.Fatalf("Error setting DISCOVERY_AS_LOCALHOST environment variable: %v", err)
	}

	walletPath := "wallet"
	// remove any existing wallet from prior runs
	os.RemoveAll(walletPath)
	wallet, err := gateway.NewFileSystemWallet(walletPath)
	if err != nil {
		log.Fatalf("Failed to create wallet: %v", err)
	}

	if !wallet.Exists("appUser") {
		err = populateWallet(wallet)
		if err != nil {
			log.Fatalf("Failed to populate wallet contents: %v", err)
		}
	}

	ccpPath := filepath.Join(
		"../../../",
		"fabric-samples",
		"test-network",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"connection-org1.yaml",
	)

	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(ccpPath))),
		gateway.WithIdentity(wallet, "appUser"),
	)

	if err != nil {
		log.Fatalf("Failed to connect to gateway: %v", err)
	}
	defer gw.Close()

	network, err := gw.GetNetwork("mychannel")
	if err != nil {
		log.Fatalf("Failed to get network: %v", err)
	}

	return &Erc20Client{
		contract:   network.GetContract(token_name),
		token_name: token_name,
	}
}

func (cc *Erc20Client) GetTokenName() string {
	return cc.token_name
}

func (cc *Erc20Client) Initialize(name string, symbol string, decimals string) ([]byte, error) {
	return cc.contract.SubmitTransaction("Initialize", name, symbol, decimals)
}

func (cc *Erc20Client) Mint(amount string) ([]byte, error) {
	return cc.contract.SubmitTransaction("Mint", amount)
}

func (cc *Erc20Client) Transfer(recipient string, amount string) ([]byte, error) {
	return cc.contract.SubmitTransaction("Transfer", recipient, amount)
}

func (cc *Erc20Client) TotalSupply() (string, error) {
	res, err := cc.contract.SubmitTransaction("TotalSupply")

	// total_supply := new(big.Int)
	// json.Unmarshal(res, total_supply)

	return string(res), err
}

func (cc *Erc20Client) BalanceOf(account string) (string, error) {
	res, err := cc.contract.EvaluateTransaction("BalanceOf", account)
	return string(res), err
}

func (cc *Erc20Client) SubmitTransaction(name string, args ...string) ([]byte, error) {
	return cc.contract.SubmitTransaction(name, args...)
}

func (cc *Erc20Client) EvaluateTransaction(name string, args ...string) ([]byte, error) {
	return cc.contract.EvaluateTransaction(name, args...)
}

func populateWallet(wallet *gateway.Wallet) error {
	log.Println("============ Populating wallet ============")
	credPath := filepath.Join(
		"../../..",
		"fabric-samples",
		"test-network",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"users",
		"User1@org1.example.com",
		"msp",
	)

	certPath := filepath.Join(credPath, "signcerts", "cert.pem")
	// read the certificate pem
	cert, err := os.ReadFile(filepath.Clean(certPath))
	if err != nil {
		return err
	}

	keyDir := filepath.Join(credPath, "keystore")
	// there's a single file in this dir containing the private key
	files, err := os.ReadDir(keyDir)
	if err != nil {
		return err
	}
	if len(files) != 1 {
		return fmt.Errorf("keystore folder should have contain one file")
	}
	keyPath := filepath.Join(keyDir, files[0].Name())
	key, err := os.ReadFile(filepath.Clean(keyPath))
	if err != nil {
		return err
	}

	identity := gateway.NewX509Identity("Org1MSP", string(cert), string(key))

	return wallet.Put("appUser", identity)
}
