package main

func main(){

	type block struct {
		magicNumber [4]byte
		blockSize [4]byte
		blockVersion [4]byte
		prevBlockHash [32]byte
		merkleRoot [32]byte
		time [4]byte
		bits [4]byte
		nonce [4]byte
		totalTxNum int64
		version [4]byte
		numInput int64
	}
	type txInput struct {
		preTxhash [32]byte
		prevTxoutIndex int32
	}
}