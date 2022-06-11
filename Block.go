package main

type Block struct {
	Preamble     Preamble
	BlockHeader  BlockHeader
	TotalTxCount TotalTxCount
	Transactions []Transaction
}

type BlockInfo struct {
	HeaderHash []byte // Version + Previous Block Hash + Merkle Root Hash + Time + Bits + nonce
	// https://dlt-repo.net/how-to-calculate-a-bitcoin-block-hash-manually/ will use step 2~ from this article
}

type Preamble struct {
	MagicNumber []byte
	BlockSize   uint32
}
type BlockHeader struct {
	BlockVersion  uint32
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
	IsSegwit      bool
	TxInputCount  uint64
	TxInput       []TxInput
	TxOutputCount uint64
	TxOutput      []TxOutput
	LockTime      uint32
}

type TxInput struct {
	PrevTxHash     []byte
	PrevTxOutIndex int32 // Signed Integer(4Bytes)
	ScriptLength   uint64
	Script         []byte // Length is specified by ScriptLength
	Sequence       int32
}

type TxOutput struct {
	OutputValue  uint64
	ScriptLength uint64
	Script       []byte
}

// TXID => nVersion + txins + txouts + nLocktime
// https://github.com/bitcoin/bips/blob/master/bip-0144.mediawiki
// Weight => Version*4, Marker*1, Flag*1, InputCount*1, Input*4, OutputCount*4, Output*4, Witness*1, Locktime*4
