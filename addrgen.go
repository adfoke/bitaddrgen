package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
)

func main() {
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
	var savestring [1000]string
	for idx, line := range lines {
		println("Hex:", line)
		d := new(big.Int)
		d.SetString(line, 16)
		fmt.Println("Decimal:", d)
		privKey, _ := btcec.NewPrivateKey(btcec.S256())
		privKey.D = d
		fmt.Println("privKey:", privKey.D)
		privKeyWif, _ := btcutil.NewWIF(privKey, &chaincfg.MainNetParams, true)
		savestring[idx] = privKeyWif.String()
		fmt.Println("privKeyWif:", savestring[idx])
	}
	var savefile ="save.txt"
	var f *os.File
	var err2 error
	f, err2 = os.OpenFile(savefile, os.O_WRONLY|os.O_CREATE, 0666)
	if err2 != nil {
        fmt.Println("file open error", err2)
    }
	defer f.Close()
	writer := bufio.NewWriter(f)
	 for _,line := range savestring{
		 writer.WriteString(line)		 
		 writer.WriteString("\n")
	}
	writer.WriteString("haha")
	writer.Flush() 
}

func checkFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
	   return false
	}
	return true
 }
