package position

import (
	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

type PositionChangeController struct{
	alpacaClient *alpaca.Client
	req          *alpaca.PlaceOrderRequest
	time         *time.Time
	logger       *zap.SugaredLogger
}

func PositionChangeRouter(r gin.IRouter, db *gorm.DB, alpacaClient *alpaca.Client, req *alpaca.PlaceOrderRequest, time *time.Time, logger *zap.SugaredLogger) {
	pc := &PositionChangeController{
		alpacaClient: alpacaClient,
		req:          req,
		time:         time,
		logger:       logger,
	}

	r.GET("", pc.positionChange)
}

func GetPositionRouter(r gin.IRouter){
	r.GET("", getPosition)
}