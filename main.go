package main

import (
    "golang.org/x/crypto/ssh/terminal"
    "github.com/ethereum/go-ethereum/crypto"
    "fmt"
    "io/ioutil"
    "os"
)

func main() {
    if ! terminal.IsTerminal(0) {
        fmt.Print("Nodekey address: ")
        b, _ := ioutil.ReadAll(os.Stdin)
        key, _ := crypto.HexToECDSA(string(b))
        addr := crypto.PubkeyToAddress(key.PublicKey)
        fmt.Printf("{%X}\n", addr)
    } else {
        fmt.Println("no piped data\n")
    }
}