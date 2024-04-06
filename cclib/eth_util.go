package cclib

import (
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

// note: Not the best way to 'wait' transaction
// func WaitTx(client *ethclient.Client, hash common.Hash) (bool, error) {
// 	for {
// 		r, err := client.TransactionReceipt(context.Background(), hash)
// 		if err == ethereum.NotFound {
// 			time.Sleep(500 * time.Millisecond)
// 			continue
// 		}
// 		if err != nil {
// 			return false, err
// 		}

// 		// print the entire receipt
// 		if r.Status != types.ReceiptStatusSuccessful {
// 			fmt.Printf("Receipt: %+v\n", r)
// 		}

// 		return r.Status == types.ReceiptStatusSuccessful, nil
// 	}
// }

// Create a Eth User Object
func NewTransactor(keyfile, password string) (*bind.TransactOpts, error) {
	f, err := os.Open(keyfile)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	auth, err := bind.NewTransactorWithChainID(f, password, big.NewInt(15))
	if err != nil {
		return nil, err
	}
	auth.GasLimit = 100000000
	return auth, nil
}
