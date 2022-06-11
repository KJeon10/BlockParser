package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

<<<<<<< Updated upstream
func main(){
	type PreAmble struct {
		magicNumber []byte
		blockSize []byte
	}
	type BlockHeader struct {
		blockVersion []byte
		prevBlockHash []byte
		merkleRoot []byte
		time []byte
		bits []byte
		nonce []byte
		totalTxNum []byte
		version []byte
		numInput []byte
		numOutput []byte 
		lockTime []byte
	}

	 type TxData struct{
		 numTx uint32
	 }

	type TxInput struct {
		prevTxhash []byte
		prevTxoutIndex []byte
		scriptLength []byte
		script []byte
		sequence []byte
	}

	type TxOutput struct {
		outputValue []byte
		scriptLength []byte
		script []byte
	}
	type Block struct {
		PreAmble PreAmble
		BlockHeader BlockHeader
		TxData TxData
	}

	blockFile, err := ioutil.ReadFile("/Users/jeonkangmin/Desktop/blk00000.dat") // Upload whole block file to buffer
=======
func main() {
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
>>>>>>> Stashed changes
	if err != nil {
		log.Fatal(err)
	}

<<<<<<< Updated upstream
=======
// Types used : []byte, uint32, uint64, bool, int32
func ReturnBlockListFromFile(BlockFile []byte) (BlockList [][]byte) {
	var Position uint32 = 0
	var BlockFileLength = uint32(len(BlockFile))
	for Position < BlockFileLength {
		var BlockLength = binary.LittleEndian.Uint32((BlockFile)[Position+4 : Position+8])
		BlockList = append(BlockList, (BlockFile)[Position:Position+BlockLength+8])
		Position = Position + BlockLength + 8
>>>>>>> Stashed changes

	
	//x:=0
	fmt.Printf("%X\n\n\n\n\n", blockFile[89])
	
	var P uint32 = 0
	/*fmt.Print(blockFile[88:97])
	a, b := VariableHandler(blockFile[88:97], howBytes)
	fmt.Printf("%d and %d", a, b)*/

	magicNumber := blockFile[0:4]
	blockSize := blockFile[4:8]
	blockVersion:= blockFile[8:12]
	prevBlockHash:= blockFile[12:44]
	merkleRoot:= blockFile[44:76]
	time:= blockFile[76:80]
	bits:= blockFile[80:84]
	nonce:= blockFile[84:88]
	numTx, P := VariableHandler(blockFile[88:97], 88)
	fmt.Println(P)
	prevTxhash := blockFile[P:P+32]
	prevTxoutIndex := blockFile[P+32:P+36]
	inputScriptLength, P := VariableHandler((blockFile[P+36:P+40]), P+36)
	fmt.Println(P)
	var L uint32
	L = 5
	inputScript := blockFile[P:P+L]
	fmt.Printf("magicNumber : %X\nblockSize : %X\nblockVersion : %X\nprevBlockHash : %X\nmerkleRoot : %X\ntime : %X\nbits : %X\nnonce : %X\nnumTx : %X\nprevTxHash : %X\npreTxOutIndex : %X\ninputScriptLength : %X\ninputScript : %X\n", magicNumber, blockSize, blockVersion, prevBlockHash, merkleRoot, time, bits, nonce, numTx, prevTxhash, prevTxoutIndex, inputScriptLength, inputScript)
	

/*	blockData := &Block{
		PreAmble : PreAmble{
			magicNumber: blockFile[0:4],
			blockSize: blockFile[5:9],
		},
		BlockHeader : BlockHeader{
			blockVersion: blockFile[8:12],
			prevBlockHash: blockFile[12:44],
			merkleRoot: blockFile[44:76],
			time: blockFile[76:80],
			bits: blockFile[80:84],
			nonce: blockFile[84:88],
		},
		TxData : TxData {
			numTx: ,
		},
	}
	*/



}

<<<<<<< Updated upstream
	func VariableHandler(current []byte, newPointer uint32) ([]byte, uint32){
		switch {
		case current[0] < 253:
			newPointer += 1
			return current[0:1], newPointer
		case current[0] == 253:
			newPointer += 2
			return current[1:2], newPointer
		case current[0] == 254:
			newPointer+= 5
			return current[1:5], newPointer
		case current[0] == 255:
			newPointer+=8
			return current[1:9], newPointer
		default:
			return current[0:1], newPointer
		}
	}
=======
func ParseResultPrinter(block *Block) {
	fmt.Printf("\n====BLOCK DATA====\n")
	fmt.Printf("MagicNumber : %s\n", hex.EncodeToString(block.Preamble.MagicNumber))
	fmt.Printf("BlockSize : %d\n", block.Preamble.BlockSize)
	fmt.Printf("BlockVersion : %d\n", block.BlockHeader.BlockVersion)
	fmt.Printf("PrevBlockHash : %s\n", hex.EncodeToString(block.BlockHeader.PrevBlockHash))
	fmt.Printf("MerkleRoot : %s\n", hex.EncodeToString(block.BlockHeader.MerkleRoot))
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
	}
}
>>>>>>> Stashed changes
