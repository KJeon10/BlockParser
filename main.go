package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
)

func main(){
	blockFile, err := ioutil.ReadFile("/Users/jeonkangmin/Desktop/blk03063.dat") // Upload whole block file to buffer
	if err != nil {
		log.Fatal(err)
	}

	blockData := blockFile [0:]

	preAmble := blockData[0:8]

	blockHeader := blockData[8:96]

	numTx, txStartPos := VariableHandler(blockData[88:97], 88)

	transactions := blockData[txStartPos:]

	fmt.Printf("MagicNumber : %s\n", hex.EncodeToString(preAmble[0:4]))	//magicnumber
	fmt.Printf("BlockSize : %s\n", hex.EncodeToString(preAmble[4:8])) //blocksize

	fmt.Printf("Version : %s\n", hex.EncodeToString(blockHeader[0:4])) //version
	fmt.Printf("PrevBlockHash : %s\n", hex.EncodeToString(blockHeader[4:36]))	//preblockhash
	fmt.Printf("MerkleRoot :%s\n", hex.EncodeToString(blockHeader[36:68])) // merkleroot
	fmt.Printf("Time : %s\n", hex.EncodeToString(blockHeader[68:72])) // time
	fmt.Printf("Bits : %s\n", hex.EncodeToString(blockHeader[72:76])) // bits
	fmt.Printf("Nonce : %s\n", hex.EncodeToString(blockHeader[76:80])) //nonce

	fmt.Printf("numTx : %d\n", numTx) // number of Transactions 
 	for Index:=0; uint64(Index)<numTx; Index++{
	transactions = TxParser(transactions, uint64(Index))
 }

	
	}

	func VariableHandler(current []byte, currentPos uint64) (uint64, uint64){
		switch {
		case current[0] == 0:
			return VariableHandler(current[2:], currentPos+2)
		case current[0] < 253 && current[0] > 0:
			return uint64(current[0]), currentPos+1
		case current[0] == 253:
			return uint64(binary.LittleEndian.Uint16(current[1:3])), currentPos+3
		case current[0] == 254:
			return uint64(binary.LittleEndian.Uint32(current[1:5])), currentPos+5
		case current[0] == 255:
			return binary.LittleEndian.Uint64(current[1:9]), currentPos+9
		default:
			return 0, 0
		}
	}

	func TxParser(TxData []byte, Index uint64)(NewTxData []byte){
		fmt.Printf("===========START TRANSACTION INDEX %d===========\n\n\n\n\n", Index)
		fmt.Printf("Version : %s\n", hex.EncodeToString(TxData[0:4]))

		numInput, InputStartPos := VariableHandler(TxData[4:], 4)

		fmt.Printf("Number of input : %d\n", numInput)

		OutputPos := InputParser(TxData[InputStartPos:], numInput) + InputStartPos

		numOutput, OutputStartPos := VariableHandler(TxData[OutputPos:], OutputPos)
		fmt.Printf("Number of Output : %d\n", numOutput)

		NewTxData = OutputParser(TxData[OutputStartPos:], numOutput)
		fmt.Printf("\n\n\n\n\n===========END TRANSACTION INDEX %d===========\n\n\n\n\n", Index)
		return NewTxData
	}


	func InputParser(InputData []byte, numInput uint64) (OutputStartPos uint64){
		var Pos uint64 = 0
		for Index:=0; uint64(Index)<numInput; Index++{
		fmt.Printf("\n\n###INPUT INDEX %d###\n\n", Index)
		fmt.Printf("Previous Tx Hash : %s\n", hex.EncodeToString(InputData[Pos:Pos+32]))
		fmt.Printf("Previous TxOut Index : %s\n", hex.EncodeToString(InputData[Pos+32:Pos+36]))
		ScriptLength, ScriptStartPos := VariableHandler(InputData[Pos+36:Pos+45], Pos+36)
		fmt.Printf("Script Length : %d\n", ScriptLength)
		fmt.Printf("Input Script : %s\n", hex.EncodeToString(InputData[ScriptStartPos:ScriptStartPos+ScriptLength]))
		fmt.Printf("Sequence : %s\n", hex.EncodeToString(InputData[ScriptStartPos+ScriptLength:ScriptStartPos+ScriptLength+4]))
		fmt.Printf("\n\n###END OF INPUT INDEX%d###\n\n", Index)
		Pos = ScriptStartPos+ScriptLength+4
		}
		return Pos
		
	}

	func OutputParser(OutputData []byte, numOutput uint64)(NewTx []byte){ 
		var Pos uint64 = 0
		for Index:=0; uint64(Index)<numOutput; Index++{
		fmt.Printf("\n\n###OUTPUT INDEX %d###\n\n", Index)
		fmt.Printf("OUTPUT VALUE : %d\n", binary.LittleEndian.Uint64(OutputData[Pos:Pos+8]))
		ScriptLength, ScriptStartPos := VariableHandler(OutputData[Pos+8:], Pos+8)
		fmt.Printf("Script Length : %d\n", ScriptLength)
		fmt.Printf("Output Script : %s\n", hex.EncodeToString(OutputData[ScriptStartPos:ScriptStartPos+ScriptLength]))
		Pos = ScriptStartPos+ScriptLength
		}
		fmt.Printf("LockTime: %d\n",binary.LittleEndian.Uint32(OutputData[Pos:Pos+4]))
		fmt.Printf("\n\n### END OF OUTPUT ###\n\n")
		return OutputData[Pos+2:]
	}