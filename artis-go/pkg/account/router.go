package account

import (
	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Controller struct {
	alpacaClient *alpaca.Client
	logger       *zap.SugaredLogger
}

func SetupRouter(r gin.IRouter, db *gorm.DB, alpacaClient *alpaca.Client, logger *zap.SugaredLogger) {
	c := &Controller{
		alpacaClient: alpacaClient,
		logger:       logger,
	}

	r.GET("", c.getAccount)
}