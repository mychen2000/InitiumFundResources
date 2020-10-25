package order

import (
	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type PlaceController struct{
	alpacaClient *alpaca.Client
	req          *alpaca.PlaceOrderRequest
	logger       *zap.SugaredLogger
}

func PlaceOrderRouter(r gin.IRouter, db *gorm.DB, alpacaClient *alpaca.Client, req *alpaca.PlaceOrderRequest, logger *zap.SugaredLogger) {
	pc := &PlaceController{
		alpacaClient: alpacaClient,
		req:          req,
		logger:       logger,
	}

	r.GET("", pc.placeOrder)
}

type GetController struct{
	alpacaClient *alpaca.Client
	orderID      *string
	logger       *zap.SugaredLogger
}

func GetOrderRouter(r gin.IRouter, db *gorm.DB, alpacaClient *alpaca.Client, orderID *string, logger *zap.SugaredLogger){
	gc := &GetController{
		alpacaClient: alpacaClient,
		orderID: orderID,
		logger: logger,
	}

	r.GET("", gc.getOrder)
}