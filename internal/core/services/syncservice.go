package services

import (
	"encoding/json"
	"net/http"
	"time"
	"zamrud/internal/core/domain"
	"zamrud/internal/core/ports"
)

type SyncService struct {
	blockchain  ports.BlockchainService
	nodeManager *NodeManager
	ticker      *time.Ticker
}

func NewSyncService(blockchain ports.BlockchainService, nodeManager *NodeManager) *SyncService {
	return &SyncService{
		blockchain:  blockchain,
		nodeManager: nodeManager,
		ticker:      time.NewTicker(5 * time.Minute),
	}
}

func (ss *SyncService) StartSyncing() {
	go func() {
		for {
			<-ss.ticker.C
			ss.SyncWithPeers()
		}
	}()
}

func (ss *SyncService) SyncWithPeers() {
	for _, peer := range ss.nodeManager.GetPeers() {
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
		if len(peerChain) > len(ss.blockchain.GetChain()) && ss.blockchain.IsValidChain(peerChain) {
			ss.blockchain.SetChain(peerChain)
		}
	}
}
