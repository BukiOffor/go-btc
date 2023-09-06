package main

import (
	"fmt"
	"github.com/libsv/go-bk/bec"
	"github.com/libsv/go-bk/crypto"
	"github.com/libsv/go-bt/v2/bscript"
	
)
func main(){
	privKey, _ := bec.NewPrivateKey(bec.S256())
	pubKey := privKey.PubKey().SerialiseUncompressed()
	//first sha256 hash of public key
	pubKey = crypto.Sha256(pubKey)
	// second ripemd160 hash of the sha256 hash to generate the public address
	pubKey = crypto.Ripemd160(pubKey)
	// make a slice of byte that holds 0 elements, elements can be added using the append keyword
	version := make([]byte, 0)
	version = append(version, 0x00) //0x4d for testnet
	pubKey = append(version, pubKey...)
	address := bscript.Base58EncodeMissingChecksum(pubKey)
	fmt.Printf("Address: %s\n", address)	
	
}