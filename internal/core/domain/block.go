package domain

import "time"

type Block struct {
	Index        int64
	Timestamp    time.Time
	Transactions []Transaction
	PrevHash     string
	Hash         string
	Nonce        int
	Difficulty   int
	Data         string
}

type Transaction struct {
	Sender    string
	Recipient string
	Amount    float64
}
