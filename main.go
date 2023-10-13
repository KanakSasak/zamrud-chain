package main

import (
	"fmt"
	"log"
	"zamrud/cmd"
	"zamrud/internal/adapters/repository"
	"zamrud/internal/core/services"
)

func main() {
	cmd.Execute()

	config := repository.NewConfigRepository()
	conf, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	storage := repository.NewFileStore()
	chainfile, err := storage.GetChain()
	if err != nil {
		panic(err)
	}
	blockchain := services.NewBlockchainService(chainfile)
	nodeManager := services.NewNodeManager()
	syncService := services.NewSyncService(blockchain, nodeManager)

	// Add known peers or let nodes add peers dynamically
	for _, peer := range conf.Peers {
		log.Println(peer)
		nodeManager.AddPeer(peer)
	}

	syncService.StartSyncing()

	Alice, err := blockchain.GenerateAddress()
	if err != nil {
		panic(err)
	}

	Bob, err := blockchain.GenerateAddress()
	if err != nil {
		panic(err)
	}

	log.Println(Alice)
	log.Println(Bob)

	// Example usage
	err = blockchain.AddTransaction(Alice, Bob, 50.0)
	if err != nil {
		panic(err)
	}

	err = blockchain.Mine()
	if err != nil {
		panic(err)
	}

	chain := blockchain.GetChain()
	log.Println(chain)
	// Check if the chain is empty and add the genesis block
	if len(chain) == 0 {
		genesisBlock := blockchain.CreateGenesisBlock()
		chain = append(chain, genesisBlock)
	}
	for _, block := range chain {
		fmt.Printf("Block: %+v\n", block)
	}

	err = storage.Save(chain)
	if err != nil {
		panic(err)
	}

	loadedChain, err := storage.Load()
	if err != nil {
		panic(err)
	}

	for _, block := range loadedChain {
		fmt.Printf("Loaded Block: %+v\n", block)
	}

}
