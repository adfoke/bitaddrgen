package main

import (
	"fmt"
	"os"
   "bufio"
   "math/big"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
)

func main(){
	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	 }
	defer file.Close()
	var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
	for _,line := range lines{
	println("Hex:",line)
	d := new(big.Int)
	d.SetString(line[2:],16)
	fmt.Println("Decimal:",d)
	privKey, _ := btcec.NewPrivateKey(btcec.S256())
	privKey.D = d
	fmt.Println("privKey:",privKey.D)
	privKeyWif, _ := btcutil.NewWIF(privKey, &chaincfg.MainNetParams, true)
	fmt.Println("privKeyWif:",privKeyWif)
	}
	
}
