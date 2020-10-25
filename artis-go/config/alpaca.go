package config

import (
	"fmt"
	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/common"
)

func AlpacaKey(cfg *Config) *common.APIKey {
	// Instead of calling on Alpaca's built-in common.Credentials() function, that reads sensitive information from
	// environment variables, I'm manually crafting the struct for minimal permission

	credentials := &common.APIKey{
		ID:           cfg.Alpaca.ApiKeyID,
		Secret:       cfg.Alpaca.ApiSecretKey,
		OAuth:        "",
		PolygonKeyID: cfg.Alpaca.ApiKeyID,
	}

	return credentials
}

func AlpacaClient(cfg *Config) *alpaca.Client {
	credentials := AlpacaKey(cfg)

	alpaca.SetBaseUrl("https://paper-api.alpaca.markets")
	fmt.Printf("Running w/ credentials [%v %v]\n", credentials.ID, credentials.Secret)

	return alpaca.NewClient(credentials)
}
