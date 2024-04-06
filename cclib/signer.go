package cclib

import (
	"io/ioutil"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Signer struct {
	privkey *keystore.Key
}

func NewSigner(keyfile, password string) (*Signer, error) {
	f, err := os.Open(keyfile)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	keyjson, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	privkey, err := keystore.DecryptKey(keyjson, password)
	if err != nil {
		return nil, err
	}
	return &Signer{
		privkey: privkey,
	}, nil
}

func (s *Signer) Address() common.Address {
	return s.privkey.Address
}

func (s *Signer) Sign(hash []byte) ([]byte, error) {
	return crypto.Sign(hash, s.privkey.PrivateKey)
}

func VerifySignature(hash, signature []byte, addr string) bool {
	pubkey, err := crypto.SigToPub(hash, signature)
	if err != nil {
		return false
	}

	if addr == crypto.PubkeyToAddress(*pubkey).Hex() {
		return true
	}
	return false
}
