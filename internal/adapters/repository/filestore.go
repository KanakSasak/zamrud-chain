package repository

import (
	"encoding/json"
	"os"
	"sync"
	"zamrud/internal/core/domain"
)

const defaultFileName = "blockchain_data.json"

type FileStore struct {
	mu sync.Mutex
}

func NewFileStore() *FileStore {
	return &FileStore{}
}

// Load loads the blockchain data from the file.
func (fs *FileStore) Load() ([]domain.Block, error) {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	file, err := os.Open(defaultFileName)
	if err != nil {
		if os.IsNotExist(err) {
			// File doesn't exist, return empty blockchain
			return []domain.Block{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var blocks []domain.Block
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&blocks)
	if err != nil {
		return nil, err
	}

	return blocks, nil
}

// Save saves the blockchain data to the file.
func (fs *FileStore) Save(blocks []domain.Block) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	file, err := os.OpenFile(defaultFileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(blocks)
	if err != nil {
		return err
	}

	return nil
}

func (fs *FileStore) GetChain() ([]domain.Block, error) {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	file, err := os.Open(defaultFileName)
	if err != nil {
		if os.IsNotExist(err) {
			// File doesn't exist, return empty blockchain
			return []domain.Block{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var blocks []domain.Block
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&blocks)
	if err != nil {
		return nil, err
	}

	return blocks, nil
}
