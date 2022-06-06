package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

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
	if err != nil {
		log.Fatal(err)
	}


	
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