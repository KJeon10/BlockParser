package main

import (
	"encoding/binary"
)

func ParseBlockStructure(BlockData *[]byte) (Preamble []byte, BlockHeader []byte, TotalTxCount uint64, Transactions []byte) {
	var blockData = *BlockData
	Preamble = blockData[0:8]
	BlockHeader = blockData[8:88]
	TotalTxCount, UsedBytes := VarintHandler(blockData[88:97])
	Transactions = blockData[88+UsedBytes:]
	return
}

func VarintHandler(varint []byte) (uint64, uint8) {
	switch {
	case varint[0] == 253:
		return binary.LittleEndian.Uint64(varint[1:3]), 3
	case varint[0] == 254:
		return binary.LittleEndian.Uint64(varint[1:5]), 5
	case varint[0] == 255:
		return binary.LittleEndian.Uint64(varint[1:9]), 9
	default:
		return uint64(varint[0]), 1
	}
}
