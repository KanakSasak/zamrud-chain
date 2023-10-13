package services

import (
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"
	"zamrud/internal/core/domain"
	"zamrud/internal/core/ports"
	//"time"
)

const targetBits = 3
const maxNonce = 1<<63 - 1

type blockchainService struct {
	chain           []domain.Block
	transactionPool []domain.Transaction
}

func NewBlockchainService(chain []domain.Block) ports.BlockchainService {
	return &blockchainService{
		chain:           chain,
		transactionPool: []domain.Transaction{},
	}
}

func (b *blockchainService) CreateGenesisBlock() domain.Block {
	return domain.Block{
		Index:        0,
		Timestamp:    time.Now(),
		Transactions: []domain.Transaction{},
		PrevHash:     "",
		Hash:         "",
		// Here, you may hard-code some initial data or leave it empty
		Data: "Genesis Block",
	}
}

func (b *blockchainService) AddTransaction(sender, recipient string, amount float64) error {
	// Validate Ethereum address format for sender and recipient
	if !b.IsValidEthereumAddress(sender) || !b.IsValidEthereumAddress(recipient) {
		return errors.New("either sender or recipient address is not a valid Ethereum address")
	}

	transaction := domain.Transaction{
		Sender:    sender,
		Recipient: recipient,
		Amount:    amount,
	}

	// Add the transaction to the transaction pool
	b.transactionPool = append(b.transactionPool, transaction)

	// Notify the miner of the new transaction
	//fmt.Println(miner)
	//if miner != nil {
	//	miner.NotifyNewTransaction()
	//} else {
	//	// Log or handle the error situation where miner is not initialized
	//	log.Fatalln("error on miner")
	//}

	return nil
}

func (b *blockchainService) Mine() error {

	lastBlock := b.chain[len(b.chain)-1]
	totalTransactions := len(lastBlock.Transactions)

	// Base difficulty
	difficulty := targetBits

	// Increase difficulty every 1000 transactions by adding an extra leading zero
	additionalDifficulty := totalTransactions / 1000
	difficulty += additionalDifficulty

	newBlock := domain.Block{
		Index:     lastBlock.Index + 1,
		Timestamp: time.Now(),
		PrevHash:  lastBlock.Hash,
		// You can adjust this part to gather transactions from a pool or some other source
		Transactions: lastBlock.Transactions,
		Difficulty:   difficulty,
		Data:         "Some data for this block",
	}

	newBlock.Transactions = b.transactionPool
	b.transactionPool = []domain.Transaction{} // Clear the pool
	//log.Println("masuk 4")
	//Apply Proof of Work
	for nonce := 0; nonce < maxNonce; nonce++ {
		//log.Println("masuk loop ", nonce)
		newBlock.Nonce = nonce
		if !isValidProof(calculateHash(newBlock), newBlock.Difficulty) {
			continue
			//log.Println("masuk x")
		}
		newBlock.Hash = calculateHash(newBlock)
		//log.Println("masuk xx")
		break

	}

	//log.Println("masuk 5")

	// Clear the transactions in the last block and add the new block to the chain
	lastBlock.Transactions = []domain.Transaction{}
	b.chain = append(b.chain, newBlock)
	//log.Println("masuk 6")
	return nil
}

func (b *blockchainService) GetChain() []domain.Block {
	return b.chain
}

func (b *blockchainService) SetChain(chain []domain.Block) {
	b.chain = chain
}

func (b *blockchainService) GetTransactionPool() []domain.Transaction {
	return b.transactionPool
}

func (b *blockchainService) IsValidChain(chain []domain.Block) bool {
	for i := 1; i < len(chain); i++ {
		currentBlock := chain[i]
		prevBlock := chain[i-1]

		// Validate block hash integrity
		if currentBlock.PrevHash != prevBlock.Hash {
			return false
		}

		// Validate current block's hash
		if calculateHash(currentBlock) != currentBlock.Hash {
			return false
		}
	}
	return true
}

func (b *blockchainService) ResolveConflicts(peers []string) bool {
	newChain := make([]domain.Block, 0)
	maxLength := len(b.chain)

	for _, node := range peers {
		log.Println(node)
		// Fetch the blockchain from peer node
		peerChain := []domain.Block{} // This would involve an HTTP GET request to the peer

		if len(peerChain) > maxLength && b.IsValidChain(b.GetChain()) {
			maxLength = len(peerChain)
			newChain = peerChain
		}
	}

	if len(newChain) != 0 {
		b.chain = newChain
		return true
	}

	return false
}

func (b *blockchainService) AddToTransactionPool(tx domain.Transaction) error {
	if !b.IsValidEthereumAddress(tx.Sender) || !b.IsValidEthereumAddress(tx.Recipient) {
		return errors.New("either sender or recipient address is not a valid Ethereum address")
	}

	b.transactionPool = append(b.transactionPool, tx)
	return nil
}

func (b *blockchainService) SyncWithPeers() {

	peers := []string{"http://localhost:8081", "http://localhost:8082"}

	for _, peer := range peers {
		resp, err := http.Get(peer + "/blockchain")
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		var peerChain []domain.Block
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&peerChain)
		if err != nil {
			continue
		}

		// Replace the chain if peer's chain is longer and valid
		if len(peerChain) > len(b.chain) && b.IsValidChain(peerChain) {
			b.chain = peerChain
		}
	}
}

func (b *blockchainService) IsValidEthereumAddress(address string) bool {
	// Check basic format using regular expression
	matched, err := regexp.MatchString("^0x[0-9a-fA-F]{40}$", address)
	if err != nil || !matched {
		return false
	}

	// Convert to all lowercase and see if it remains the same (all small letters or all caps are valid)
	lowerAddress := strings.ToLower(address[2:])
	if address[2:] == lowerAddress || strings.ToUpper(address[2:]) == lowerAddress {
		return true
	}

	// If it's a mixed-case address, it should be a valid EIP-55 address
	return common.IsHexAddress(address)
}

func (b *blockchainService) GenerateAddress() (string, error) {
	address, err := ethereumAddress()
	if err != nil {
		return "", err
	}
	return address, nil
}
