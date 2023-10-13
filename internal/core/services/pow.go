package services

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"strings"
	"zamrud/internal/core/domain"
)

func calculateHash(block domain.Block) string {
	record := fmt.Sprintf("%d%s%d", block.Index, block.Timestamp.String(), block.Nonce)
	for _, tx := range block.Transactions {
		record += tx.Sender + tx.Recipient + fmt.Sprintf("%f", tx.Amount)
	}
	hash := crypto.Keccak256([]byte(record))
	return hex.EncodeToString(hash)
}

func isValidProof(hash string, difficulty int) bool {
	prefix := strings.Repeat("0", difficulty)
	return strings.HasPrefix(hash, prefix)
}
