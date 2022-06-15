package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	HASH160()

	ExecutionTimerStart := time.Now()
	BlockFile := OpenBlockfile()
	BlockList := ReturnBlockListFromFile(BlockFile)
	for i := 0; i < len(BlockList); i++ {
		ParseResultPrinter(ParseBlock(BlockList[i]))
	}
	elapsed := time.Since(ExecutionTimerStart)
	fmt.Printf("\n\n+++Execution Time : %s+++\n\n", elapsed)
}

func OpenBlockfile() (BlockFile []byte) {
	BlockFile, err := os.ReadFile("/Users/jeonkangmin/Desktop/blk00000.dat") // Upload whole block file to memory for faster indexing
	if err != nil {
		log.Fatal(err)
	}
	return BlockFile
}

// Types used : []byte, uint32, uint64, bool, int32
func ReturnBlockListFromFile(BlockFile []byte) (BlockList [][]byte) {
	var Position uint32 = 0
	var BlockFileLength = uint32(len(BlockFile))
	for Position < BlockFileLength {
		if BlockFile[Position] == 249 {
			var BlockLength = binary.LittleEndian.Uint32((BlockFile)[Position+4 : Position+8])
			BlockList = append(BlockList, (BlockFile)[Position:Position+BlockLength+8])
			Position = Position + BlockLength + 8
		} else {
			return BlockList
		}
	}
	return BlockList
}
func ParseResultPrinter(block *Block) {

	/*fmt.Printf("====BLOCK DATA====\n")

	fmt.Printf("MagicNumber : %s\n", hex.EncodeToString(block.Preamble.MagicNumber))
	fmt.Printf("BlockSize : %d\n", block.Preamble.BlockSize)
		fmt.Printf("BlockVersion : %d\n", block.BlockHeader.BlockVersion)*/
	fmt.Printf("PrevBlockHash : %s\n", hex.EncodeToString(block.BlockHeader.PrevBlockHash))
	/*fmt.Printf("MerkleRoot : %s\n", hex.EncodeToString(block.BlockHeader.MerkleRoot))
	fmt.Printf("Time : %d\n", block.BlockHeader.Time)
	fmt.Printf("Nonce : %d\n", block.BlockHeader.Nonce)
	fmt.Printf("TotalTxCount : %d\n\n", block.TotalTxCount)
	for i := 0; i < int(block.TotalTxCount); i++ {
		fmt.Printf("\n----- Transaction INDEX %d -----\n\n", i)
		fmt.Printf("TxVersion : %d\n", block.Transactions[i].TxVersion)
		fmt.Printf("IsSegwit : %t\n", block.Transactions[i].IsSegwit)
		fmt.Printf("TxInputCount : %d\n", block.Transactions[i].TxInputCount)
		for j := 0; uint64(j) < block.Transactions[i].TxInputCount; j++ {
			fmt.Printf("\n+++Tx Input Index %d of Transactin Index %d+++\n\n", j, i)
			fmt.Printf("PrevTxHash : %s\n", hex.EncodeToString(block.Transactions[i].TxInputs[j].PrevTxHash))
			fmt.Printf("PrevTxOutIndex : %d\n", block.Transactions[i].TxInputs[j].PrevTxOutIndex)
			fmt.Printf("ScriptLength : %d\n", block.Transactions[i].TxInputs[j].ScriptLength)
			fmt.Printf("Script : %s\n", hex.EncodeToString(block.Transactions[i].TxInputs[j].Script))
			fmt.Printf("Sequence : %d\n", block.Transactions[i].TxInputs[j].Sequence)
			fmt.Printf("++++++++++++++++++++++++++++++++++++++++++++++++\n\n")
		}
		fmt.Printf("TxOuputCount : %d\n", block.Transactions[i].TxOutputCount)
		for k := 0; uint64(k) < block.Transactions[i].TxOutputCount; k++ {
			fmt.Printf("\n+++Tx Output Index %d of Transactin Index %d+++\n\n", k, i)
			fmt.Printf("OutputValue: %d\n", block.Transactions[i].TxOutputs[k].OutputValue)
			fmt.Printf("ScriptLength: %d\n", block.Transactions[i].TxOutputs[k].ScriptLength)
			fmt.Printf("Script: %s\n", hex.EncodeToString(block.Transactions[i].TxOutputs[k].Script))
			fmt.Printf("++++++++++++++++++++++++++++++++++++++++++++++++\n\n")
		}
		fmt.Printf("LockTime : %d\n", block.Transactions[i].LockTime)
		fmt.Printf("------------------------------------------------\n")
	}*/
}
