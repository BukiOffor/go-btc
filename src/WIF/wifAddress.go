package main

import (
	"fmt"

	//"github.com/bitcoinsv/bsvd/version"
	"github.com/libsv/go-bk/base58"
	"github.com/libsv/go-bk/bec"
	"github.com/libsv/go-bk/crypto"
)

func main(){
	privKey,_ := bec.NewPrivateKey(bec.S256())
	pKey := privKey.Serialise()
	version := make([]byte, 0)
	version = append(version, 0x80)
	pKey = append(version, pKey...)
	pKey = append(pKey, 0x01) 
	chksum := crypto.Sha256d(pKey)[:4]
	pKey = append(pKey, chksum...)
	w := base58.Encode(pKey)
	fmt.Printf("WIF: %s\n", w) 

}