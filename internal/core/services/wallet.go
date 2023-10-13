package services

import (
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/crypto"
)

func ethereumAddress() (string, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return address, nil
}
