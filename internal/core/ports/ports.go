package ports

import "zamrud/internal/core/domain"

type BlockchainService interface {
	AddTransaction(sender, recipient string, amount float64) error
	Mine() error
	GetChain() []domain.Block
	IsValidChain(chain []domain.Block) bool
	ResolveConflicts(peers []string) bool
	AddToTransactionPool(tx domain.Transaction) error
	SyncWithPeers()
	IsValidEthereumAddress(address string) bool
	GenerateAddress() (string, error)
	GetTransactionPool() []domain.Transaction
	SetChain(chain []domain.Block)
	CreateGenesisBlock() domain.Block
}

type StorageRepository interface {
	SaveChain(chain []domain.Block) error
	LoadChain() ([]domain.Block, error)
	GetChain() ([]domain.Block, error)
}

type ConfigRepository interface {
	LoadConfig() (Config, error)
}

type Config struct {
	NodeAddress string   `yaml:"node_address"`
	Peers       []string `yaml:"peers"`
}
