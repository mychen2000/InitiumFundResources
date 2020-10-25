package main

import (
	"fmt"
	"github.com/initiumfund/artis-go/config"
)

func main() {
	// Note: 这边的 relative path 是 relative to 你运行时的目录 (而不是 ./config/alpaca.go) 的
	cfg := config.LoadConfig("./config/config.yaml")
	alpacaClient := config.AlpacaClient(cfg)

	acct, err := alpacaClient.GetAccount()
	if err != nil {
		panic(err)
	}

	fmt.Println(*acct)
}
