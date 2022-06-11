package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	ExecutionTimerStart := time.Now()

	BlockFile := OpenBlockfile()
	BlockList := ReturnBlockListFromFile(BlockFile)
	for i := 0; i < len(BlockList); i++ {
		fmt.Printf("####### BLOCK INDEX %d ######\n\n\n", i)
		fmt.Printf("%d", len(BlockList[i]))
		fmt.Printf("\n\n\n####### END BLOCK INDEX %d ######\n\n\n", i)
	}

	ExecutionTimeElapsed := time.Since(ExecutionTimerStart)
	fmt.Printf("+++++EXECUTION TIME TOTAL %s+++++", ExecutionTimeElapsed)
}

func OpenBlockfile() (BlockFile []byte) {
	BlockFile, err := os.ReadFile("/Users/jeonkangmin/Desktop/blk00000.dat") // Upload whole block file to memory for faster indexing
	if err != nil {
		log.Fatal(err)
	}
	return BlockFile
}

// Types used : []byte, uint32, uint64, bool, int32
func ReturnBlockListFromFile(BlockFile []byte) [][]byte {
	var Position uint32 = 0
	var BlockFileLength = uint32(len(BlockFile))
	var BlockList [][]byte
	for Position < BlockFileLength {
		var BlockLength = binary.LittleEndian.Uint32(BlockFile[Position+4 : Position+8])
		BlockList = append(BlockList, BlockFile[Position:Position+BlockLength])
		Position = Position + BlockLength + 8

	}
	return BlockList
}
