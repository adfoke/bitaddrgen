package main

import (
	"fmt"
	"math/big"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
)

func main(){
	d := new(big.Int)
	d.SetString("0x60f4d11574f5deee49961d9609ac6"[2:],16)
	fmt.Println("Decimal:",d)
	privKey, _ := btcec.NewPrivateKey(btcec.S256())
	privKey.D = d
	fmt.Println("privKey:",privKey.D)
	privKeyWif, _ := btcutil.NewWIF(privKey, &chaincfg.MainNetParams, true)
	fmt.Println("privKeyWif:",privKeyWif)
}
