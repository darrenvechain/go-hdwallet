package hdwallet

import (
	"github.com/ethereum/go-ethereum/crypto"
)

func init() {
	coins[VET] = newVET
}

type vet struct {
	name   string
	symbol string
	key    *Key

	// eth token
	contract string
}

func newVET(key *Key) Wallet {
	return &vet{
		name:   "Vechain",
		symbol: "VET",
		key:    key,
	}
}

func (c *vet) GetType() uint32 {
	return c.key.Opt.CoinType
}

func (c *vet) GetName() string {
	return c.name
}

func (c *vet) GetSymbol() string {
	return c.symbol
}

func (c *vet) GetKey() *Key {
	return c.key
}

func (c *vet) GetAddress() (string, error) {
	return crypto.PubkeyToAddress(*c.key.PublicECDSA).Hex(), nil
}
