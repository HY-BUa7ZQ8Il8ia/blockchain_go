package main

import (
	"flag"
	"log"
)

//ログ出力設定
func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	port := flag.Uint("port", 5000, "TCP Port Number for Blockchain Server")
	flag.Parse()
	// portはアドレスで返るため、*portに変換、unit16で引数にする
	app := NewBlockchainServer(uint16(*port))
	app.Run()
}
