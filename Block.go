package main

type BlockInfo struct {
	HeaderHash []byte // Version + Previous Block Hash + Merkle Root Hash + Time + Bits + nonce
	// https://dlt-repo.net/how-to-calculate-a-bitcoin-block-hash-manually/ will use step 2~ from this article
}
type BlockHeader struct {
	BlockVersion  []byte
	PrevBlockHash []byte
	MerKleRoot    []byte
	Time          uint32
	Bits          uint32
	Nonce         uint32
}

type TotalTxCount struct {
	TotalTxCount uint64
}

type Transaction struct {
	TxVersion     uint32
	TxInputCount  uint64
	TxInput       []TxInput
	TxOutputCount uint64
	TxOutput      []TxOutput
	LockTime      uint32
}
