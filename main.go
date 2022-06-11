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
	for i := 0; i < len(*BlockList); i++ {
		fmt.Printf("####### BLOCK INDEX %d ######\n", i)
		fmt.Printf("%d", len((*BlockList)[i]))
		fmt.Printf("\n####### END BLOCK INDEX %d ######\n", i)
	}
	ExecutionTimeElapsed := time.Since(ExecutionTimerStart)
	fmt.Printf("+++++EXECUTION TIME TOTAL %s+++++", ExecutionTimeElapsed)
}

func OpenBlockfile() (BlockFile *[]byte) {
	blockFile, err := os.ReadFile("/Users/jeonkangmin/Desktop/blk00000.dat") // Upload whole block file to memory for faster indexing
	BlockFile = &blockFile
	if err != nil {
		log.Fatal(err)
	}
	return BlockFile
}

// Types used : []byte, uint32, uint64, bool, int32
func ReturnBlockListFromFile(BlockFile *[]byte) (BlockList *[][]byte) {
	var Position uint32 = 0
	var BlockFileLength = uint32(len(*BlockFile))
	var blockList [][]byte
	BlockList = &blockList
	for Position < BlockFileLength {
		var BlockLength = binary.LittleEndian.Uint32((*BlockFile)[Position+4 : Position+8])
		blockList = append(blockList, (*BlockFile)[Position:Position+BlockLength])
		Position = Position + BlockLength + 8

	}
	return BlockList
}
