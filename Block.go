package main

type BlockInfo struct {
	HeaderHash string // Version + Previous Block Hash + Merkle Root Hash + Time + Bits + nonce
}
type BlockHeader struct {
	BlockVersion  int32
	PrevBlockHash string
	MerKleRoot    string
	Time          uint32
	Bits          uint32
	Nonce         uint32
}
