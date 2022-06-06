package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {  

    f, err := os.Open("/Users/jeonkangmin/Desktop/blk00000.dat")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()
	
	stat, err := f.Stat()
	if err!=nil {
		return
	}

    reader := bufio.NewReader(f)
    buf := make([]byte, stat.Size())



    for {
        _, err := reader.Read(buf)

        if err != nil {
            if err != io.EOF {
                fmt.Println(err)
            }
            break
        }
        
        fmt.Printf("%s", hex.Dump(buf))
    }
}