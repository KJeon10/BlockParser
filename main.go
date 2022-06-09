package main

import (
	"log"
	"os"
)

func main() {

}

func OpenBlockfile() (BlockFile []byte) {
	BlockFile, err := os.ReadFile("/Users/jeonkangmin/Desktop/blk03063.dat") // Upload whole block file to buffer
	if err != nil {
		log.Fatal(err)
	}
	return BlockFile
}

// Types used : []byte, uint32, uint64, bool, int32
func RetrnNbytes() {

}
