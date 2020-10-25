package accountActivities

import (
	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Controller struct {
	alpacaClient *alpaca.Client
	activityType *string
	opts         *alpaca.AccountActivitiesRequest
	logger       *zap.SugaredLogger
}

func ActivitiesRouter(r gin.IRouter, db *gorm.DB, alpacaClient *alpaca.Client, activityType *string, opts *alpaca.AccountActivitiesRequest, logger *zap.SugaredLogger) {
	c := &Controller{
		alpacaClient: alpacaClient,
		activityType: activityType,
		opts:         opts,
		logger:       logger,
	}

	r.GET("", c.getAccountActivities)
}