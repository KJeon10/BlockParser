package main

import (
	"encoding/binary"
)

func ParseBlock(BlockData []byte) *Block {
	var UsedBytes uint8
	block := &Block{}
	block.Preamble = *PreambleParser(BlockData[0:8])
	block.BlockHeader = *BlockHeaderParser(BlockData[8:88])
	block.TotalTxCount, UsedBytes = VarintHandler(BlockData[88:97])
	block.Transactions = *TransactionsParser(BlockData[88+UsedBytes:], block.TotalTxCount)
	return block
}

func VarintHandler(varint []byte) (uint64, uint8) {
	switch {
	case varint[0] == 253:
		return uint64(binary.LittleEndian.Uint16(varint[1:3])), 3
	case varint[0] == 254:
		return uint64(binary.LittleEndian.Uint32(varint[1:5])), 5
	case varint[0] == 255:
		return binary.LittleEndian.Uint64(varint[1:9]), 9
	default:
		return uint64(varint[0]), 1
	}
}
func PreambleParser(Data []byte) *Preamble {
	preamble := &Preamble{}
	preamble.MagicNumber = Data[0:4]
	preamble.BlockSize = binary.LittleEndian.Uint32(Data[4:8])
	return preamble
}

func BlockHeaderParser(Data []byte) *BlockHeader {
	blockheader := &BlockHeader{}
	blockheader.BlockVersion = binary.LittleEndian.Uint32(Data[0:4])
	blockheader.PrevBlockHash = Data[4:36]
	blockheader.MerkleRoot = Data[36:68]
	blockheader.Time = binary.LittleEndian.Uint32(Data[68:72])
	blockheader.Bits = binary.LittleEndian.Uint32(Data[72:76])
	blockheader.Nonce = binary.LittleEndian.Uint32(Data[76:80])
	return blockheader
}

func TransactionsParser(Data []byte, numTx uint64) *[]Transaction {
	var UsedBytes uint8
	var transactions []Transaction
	Pos := 0
	for i := 0; uint64(i) < numTx; i++ {
		transaction := &Transaction{}
		transaction.TxVersion = binary.LittleEndian.Uint32(Data[Pos : Pos+4])
		transaction.TxInputCount, UsedBytes = VarintHandler(Data[Pos+4 : Pos+13])

		Pos += 4 + int(UsedBytes)

		for i := 0; i < int(transaction.TxInputCount); i++ {
			txInput := &TxInput{}
			txInput.PrevTxHash = Data[Pos : Pos+32]
			txInput.PrevTxOutIndex = int32(binary.LittleEndian.Uint32(Data[Pos+32 : Pos+36]))
			txInput.ScriptLength, UsedBytes = VarintHandler(Data[Pos+36 : Pos+45])
			txInput.Script = Data[Pos+36+int(UsedBytes) : Pos+36+int(UsedBytes)+int(txInput.ScriptLength)]
			txInput.Sequence = int32(binary.LittleEndian.Uint32(Data[Pos+36+int(UsedBytes)+int(txInput.ScriptLength) : 40+int(UsedBytes)+int(txInput.ScriptLength)+Pos]))
			Pos += 40 + int(UsedBytes) + int(txInput.ScriptLength)
			transaction.TxInputs = append(transaction.TxInputs, *txInput)
		}

		transaction.TxOutputCount, UsedBytes = VarintHandler(Data[Pos : Pos+9])

		Pos = Pos + int(UsedBytes)

		for i := 0; i < int(transaction.TxOutputCount); i++ {
			txOutput := &TxOutput{}
			txOutput.OutputValue = binary.LittleEndian.Uint64(Data[Pos : Pos+8])
			txOutput.ScriptLength, UsedBytes = VarintHandler(Data[Pos+8 : Pos+17])
			txOutput.Script = Data[Pos+8+int(UsedBytes) : Pos+8+int(UsedBytes)+int(txOutput.ScriptLength)]
			Pos = Pos + 8 + int(UsedBytes) + int(txOutput.ScriptLength)
			transaction.TxOutputs = append(transaction.TxOutputs, *txOutput)
		}
		transaction.LockTime = binary.LittleEndian.Uint32(Data[Pos : Pos+4])
		transactions = append(transactions, *transaction)
		Pos = Pos + 4
	}
	return &transactions
}
