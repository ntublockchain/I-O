package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/Guy1m0/Blockchain-I-O/cclib"
	"github.com/Guy1m0/Blockchain-I-O/contracts/cb1p_auction"
	"github.com/Guy1m0/Blockchain-I-O/contracts/english_auction"
	"github.com/Guy1m0/Blockchain-I-O/contracts/stable_coin"
	"github.com/Guy1m0/Blockchain-I-O/examples/ecomm"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/olekukonko/tablewriter"

	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

const (
	key_path = "../../keys/"
	rootKey  = "../../keys/key0"

	auctioneerKey = "../../keys/key1"
	keyFolder     = "../../keys/"
	bidder1Key    = "../../keys/key2"
	bidder2Key    = "../../keys/key3"
	password      = "password"

	contractInfoFile = "../tmp/contract_info.json"
	userInfoFile     = "../tmp/user_info.json"
	auctionInfoFile  = "../tmp/auction_info.json"
	logCSVPath       = "../tmp/log.csv"

	defaultHeaders = "AssetID,Event,KeyWords,StartTime,EndTime,KafkaReceived,GasCost,Note,TimeElapsed,KafkaTime\n"
)

var (
	rootT *bind.TransactOpts

	token_name = "MDai"

	ethClient *ethclient.Client
	quoClient *ethclient.Client

	plt = "eth"
	amt = "100"
)

func main() {
	rootT, _ = cclib.NewTransactor(rootKey, password)
	ethClient = ecomm.NewEthClient()
	quoClient = ecomm.NewQuorumClient()

	if ecomm.CheckClientValidity(quoClient) {
		fmt.Println("Successfully connected to the Ethereum client.")
	} else {
		fmt.Println("Failed to connect to the Ethereum client.")
	}

	command := flag.String("c", "", "command")
	usr := flag.String("usr", "", "user name")
	flag.StringVar(&token_name, "t", token_name, "Stable coin token name")
	flag.StringVar(&amt, "amt", amt, "Set new user initial balance")
	flag.StringVar(&plt, "p", plt, "Platform for new user")
	flag.Parse()

	switch *command {
	case "test":
		test()
	case "init":
		initialize(token_name)
	case "setup":
		setup()
	case "display":
		display()
	case "add":
		add_user(*usr, plt, amt)
	default:
		fmt.Println("command not found")
	}

}

func reset_log() error {
	// Ensure the directory exists
	dir := filepath.Dir(logCSVPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// Open (or create) the file in write mode to reset it or create a new one
	file, err := os.Create(logCSVPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the column headers back to the file
	_, err = file.WriteString(defaultHeaders)
	if err != nil {
		return err
	}

	return nil
}

// cleanFileContent opens the file at the given filePath and truncates it to zero length,
// effectively cleaning its content.
func cleanFileContent(filePath string) error {
	// Try to open the file in read mode
	_, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			// If the file does not exist, use default headers
			return nil
		} else {
			return err
		}
	} else {
		err = os.Remove(filePath)
		check(err)
	}

	return nil
}

func test() {
	// _, tx, quo, _ := stable_coin.DeployStableCoin(rootT, quoClient, big.NewInt(1))
	// ecomm.WaitTx(quoClient, tx, "Deploy ERC20 Stable Coin on Quorum")
	// //log.Println(receipt)

	// valueB, err := quo.TotalSupply(&bind.CallOpts{})
	// if err != nil {
	// 	log.Printf("Error fetching balance: %s", err)
	// } else {
	// 	log.Printf("Balance: %s", valueB.String())
	// }
	hash := solsha3.SoliditySHA3(
		solsha3.Uint256(big.NewInt(0).Mul(big.NewInt(1), ecomm.DecimalB)), //8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b
	)
	fmt.Println(hex.EncodeToString(hash)) // b10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6

}

// Deploy contracts and mint enough tokens
func initialize(token_name string) {
	fabricToken := ecomm.NewErc20Client(token_name)

	fmt.Println("Initialize Log and Info files")
	err := reset_log()
	check(err)

	err = cleanFileContent(userInfoFile)
	check(err)

	err = cleanFileContent(contractInfoFile)
	check(err)

	// err = cleanFileContent(auctionInfoFile)
	// check(err)

	fmt.Println("Initialize Fabric Stable Coin: ", token_name)
	_, err = fabricToken.Initialize("Multi-Dai Stablecoin", "MDAI", "15")
	check(err)

	fmt.Printf("Mint %s on Frabic \n", token_name)
	_, err = fabricToken.Mint("10000000000")
	check(err)

	supply, _ := big.NewInt(0).SetString("1"+strings.Repeat("0", ecomm.Decimal+10), 10)

	fmt.Println("Deploy ERC20 contracts on Eth and Quorum")
	eth_MDAI_addr, tx, eth_MDAI, _ := stable_coin.DeployStableCoin(rootT, ethClient, big.NewInt(1))
	ecomm.WaitTx(ethClient, tx, "Deploy ERC20 Stable Coin on Ethereum")
	// valueB, _ := eth_MDAI.BalanceOf(&bind.CallOpts{}, common.HexToAddress("0xd0a73fe9d44184e9f1264ce2097064212e67ebfe"))
	// log.Printf("Balance: %s", valueB.String())

	eth_english_addr, tx, _, _ := english_auction.DeployEnglishAuction(rootT, ethClient, eth_MDAI_addr)
	ecomm.WaitTx(ethClient, tx, "Deploy English Auction on Ethereum")

	// eth_dutch_addr, tx, _, _ := dutch_auction.DeployDutchAuction(rootT, ethClient, eth_MDAI_addr)
	// ecomm.WaitTx(ethClient, tx, "Deploy Dutch Auction on Ethereum")
	// //_ = debugTransaction(tx)

	eth_closed_bid_addr, tx, _, _ := cb1p_auction.DeployCb1pAuction(rootT, ethClient, eth_MDAI_addr)
	ecomm.WaitTx(ethClient, tx, "Deploy Closed Bid Auction on Ethereum")

	// eth_closed_bid_2nd_addr, tx, _, _ := cb2p_auction.DeployCb2pAuction(rootT, ethClient, eth_MDAI_addr)
	// ecomm.WaitTx(ethClient, tx, "Deploy Closed Bid Auction on Ethereum")

	tx, err = eth_MDAI.Mint(rootT, rootT.From, supply)
	check(err)
	ecomm.WaitTx(ethClient, tx, "Mint ERC20 Stable Coin on Ethereum")

	quo_MDAI_addr, tx, quo_MDAI, _ := stable_coin.DeployStableCoin(rootT, quoClient, big.NewInt(1))
	ecomm.WaitTx(quoClient, tx, "Deploy ERC20 Stable Coin on Quorum")

	quo_english_addr, tx, _, _ := english_auction.DeployEnglishAuction(rootT, quoClient, quo_MDAI_addr)
	ecomm.WaitTx(quoClient, tx, "Deploy English Auction on Quorum")

	// quo_dutch_addr, tx, _, _ := dutch_auction.DeployDutchAuction(rootT, quoClient, quo_MDAI_addr)
	// ecomm.WaitTx(quoClient, tx, "Deploy Dutch Auction on Ethereum")

	quo_closed_bid_addr, tx, _, _ := cb1p_auction.DeployCb1pAuction(rootT, quoClient, quo_MDAI_addr)
	ecomm.WaitTx(quoClient, tx, "Deploy Closed Bid Auction on Quorum")

	// quo_closed_bid_2nd_addr, tx, _, _ := cb2p_auction.DeployCb2pAuction(rootT, quoClient, quo_MDAI_addr)
	// ecomm.WaitTx(quoClient, tx, "Deploy Closed Bid Auction on Ethereum")

	tx, err = quo_MDAI.Mint(rootT, rootT.From, supply)
	check(err)
	ecomm.WaitTx(quoClient, tx, "Mint ERC20 Stable Coin on Quorum")

	ecomm.WriteJsonFile(contractInfoFile, ecomm.ContractInfo{
		FabricTokenName: token_name,
		EthERC20:        eth_MDAI_addr,
		QuoERC20:        quo_MDAI_addr,
		EnglishAuction: ecomm.AucConractInfo{
			Owner:   rootT.From,
			AucType: "English",
			QuoAddr: quo_english_addr,
			EthAddr: eth_english_addr,
		},
		// DutchAuction: ecomm.AuctionInfo{
		// 	Owner:   rootT.From,
		// 	QuoAddr: quo_dutch_addr,
		// 	EthAddr: eth_dutch_addr,
		// },
		Cb1pAuction: ecomm.AucConractInfo{
			Owner:   rootT.From,
			AucType: "ClosedBid1stPrice",
			QuoAddr: quo_closed_bid_addr,
			EthAddr: eth_closed_bid_addr,
		},
		// Cb2pAuction: ecomm.AuctionInfo{
		// 	Owner:   rootT.From,
		// 	QuoAddr: quo_closed_bid_2nd_addr,
		// 	EthAddr: eth_closed_bid_2nd_addr,
		// },
	})
}

func setup() {
	aucT, err := cclib.NewTransactor(auctioneerKey, password)
	eth_ERC20, quo_ERC20, fabric_ERC20 := load_ERC20()
	check(err)

	fmt.Println("Setup account for 'auctioneer 1' on Fabirc")
	_, err = fabric_ERC20.Transfer(aucT.From.Hex(), "0")
	check(err)

	valueB_, err := fabric_ERC20.BalanceOf(aucT.From.Hex())
	check(err)
	valueB, err := strconv.Atoi(valueB_)
	check(err)
	if valueB < 200 {
		_, err = fabric_ERC20.Transfer(aucT.From.Hex(), "100")
		check(err)
	}

	ecomm.AddUserToFile(userInfoFile, ecomm.UserInfo{
		UserID:  "auctioneer 1",
		Address: aucT.From,
		KeyFile: auctioneerKey,
	})

	var bidT *bind.TransactOpts
	for i := 1; i < 9; i++ {
		bidT, err = cclib.NewTransactor(fmt.Sprintf("%skey%s", keyFolder, strconv.Itoa(i+1)), password)
		check(err)
		log.Printf("Setup account for 'Bidder%d' on Fabric", i)
		_, err = fabric_ERC20.Transfer(bidT.From.Hex(), "0")
		check(err)
		log.Printf("Receive MDAI for 'Bidder%d' on Etherrum", i)
		ecomm.TransferToken(ethClient, eth_ERC20, rootT, bidT.From, 100000)

		log.Printf("Receive MDAI for 'Bidder%d' on Quo", i)
		ecomm.TransferToken(quoClient, quo_ERC20, rootT, bidT.From, 100000)

		ecomm.AddUserToFile(userInfoFile, ecomm.UserInfo{
			UserID:  "Bidder " + strconv.Itoa(i),
			Address: bidT.From,
			KeyFile: fmt.Sprintf("%skey%s", keyFolder, strconv.Itoa(i+1)),
		})
	}
}

func display() {
	DecimalB, _ := big.NewInt(0).SetString("1"+strings.Repeat("0", 15), 10)
	// var contract_info ecomm.ConractInfo
	// ecomm.ReadJsonFile(contractInfoFile, &contract_info)

	users, err := ecomm.ReadUsersFromFile(userInfoFile)
	check(err)

	eth_ERC20, quo_ERC20, fabric_ERC20 := load_ERC20()

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"User ID", "Ethereum", "Quorum", "Fabric"})

	for _, user := range users {
		// Eth balance
		valueB, _ := eth_ERC20.BalanceOf(&bind.CallOpts{}, user.Address)
		//log.Printf("For user: %s has balance: %s", user.UserID, valueB.String())
		eth_balance := big.NewInt(0).Div(valueB, DecimalB).String()

		// Quo balance
		valueB, _ = quo_ERC20.BalanceOf(&bind.CallOpts{}, user.Address)
		//log.Printf("For user: %s has balance: %s", user.UserID, valueB.String())

		quo_balance := big.NewInt(0).Div(valueB, DecimalB).String()
		// Fabric balance
		b, _ := fabric_ERC20.BalanceOf(user.Address.Hex())

		row := []string{
			user.UserID,
			eth_balance,
			quo_balance,
			string(b),
		}
		table.Append(row)
		//fmt.Println(user.UserID, user.Address, user.KeyFile)
	}
	var contract_info ecomm.ContractInfo
	ecomm.ReadJsonFile(contractInfoFile, &contract_info)
	auction_infos := []ecomm.AucConractInfo{
		contract_info.EnglishAuction,
		//contract_info.DutchAuction,
		contract_info.Cb1pAuction,
		//contract_info.Cb2pAuction,
	}

	for _, auction_info := range auction_infos {
		// Eth balance
		valueB, _ := eth_ERC20.BalanceOf(&bind.CallOpts{}, auction_info.EthAddr)
		eth_balance := big.NewInt(0).Div(valueB, DecimalB).String()

		// Quo balance
		valueB, _ = quo_ERC20.BalanceOf(&bind.CallOpts{}, auction_info.QuoAddr)
		quo_balance := big.NewInt(0).Div(valueB, DecimalB).String()

		row := []string{
			"Auction: " + auction_info.AucType,
			eth_balance,
			quo_balance,
			"0",
		}
		table.Append(row)
	}

	table.Render()
}

func add_user(user_id string, platform string, amount string) {
	// Get contract address and corresponding client
	users, err := ecomm.ReadUsersFromFile(userInfoFile)
	check(err)

	user_key := fmt.Sprintf("%s%s%d", key_path, "key", len(users)+1)
	userT, err := cclib.NewTransactor(user_key, password)
	check(err)

	amout_, _ := strconv.ParseInt(amount, 10, 64)
	eth_ERC20, quo_ERC20, fabric_ERC20 := load_ERC20()

	if platform == "eth" {
		ecomm.TransferToken(ethClient, eth_ERC20, rootT, userT.From, amout_)
		_, err = fabric_ERC20.Transfer(userT.From.Hex(), "0")
		check(err)
	} else if platform == "quo" {
		ecomm.TransferToken(quoClient, quo_ERC20, rootT, userT.From, amout_)
		_, err = fabric_ERC20.Transfer(userT.From.Hex(), "0")
		check(err)
	} else {
		_, err = fabric_ERC20.Transfer(userT.From.Hex(), amount)
		check(err)
	}

	ecomm.AddUserToFile(userInfoFile, ecomm.UserInfo{
		UserID:  user_id,
		Address: userT.From,
		KeyFile: user_key,
	})
}

func load_ERC20() (*stable_coin.StableCoin, *stable_coin.StableCoin, *ecomm.Erc20Client) {
	var contract_info ecomm.ContractInfo
	ecomm.ReadJsonFile(contractInfoFile, &contract_info)

	eth_ERC20, err := stable_coin.NewStableCoin(contract_info.EthERC20, ethClient)
	check(err)

	quo_ERC20, err := stable_coin.NewStableCoin(contract_info.QuoERC20, quoClient)
	check(err)

	fabric_ERC20 := ecomm.NewErc20Client(contract_info.FabricTokenName)

	return eth_ERC20, quo_ERC20, fabric_ERC20
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
