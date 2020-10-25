package main

import (
	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/gin-gonic/gin"
	"github.com/initiumfund/artis-go/config"
	accountService "github.com/initiumfund/artis-go/pkg/account"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"
	"io"
	"net/http"
	"os"
	"time"
)

func SetupRouter(cfg *config.Config, db *gorm.DB, alpacaClient *alpaca.Client, log *zap.SugaredLogger) (*gin.Engine, error) {
	r := gin.New()

	//r.Use(gin.Logger())
	r.Use(Logger(log))

	stdLog, err := zap.NewStdLogAt(log.Desugar(), zapcore.ErrorLevel)
	if err != nil {
		return nil, err
	}
	r.Use(gin.RecoveryWithWriter(io.MultiWriter(os.Stderr, stdLog.Writer())))

	// TODO: CORS

	// TODO: Auth

	r.GET("/health-check", HealthCheck)

	rApi := r.Group("/api")
	// rApi.Use(<authentication middleware>)
	rApiV2 := rApi.Group("/v2")

	// the following variables are within the scope of this block
	{
		account := rApiV2.Group("/account")
		accountService.SetupRouter(account, db, alpacaClient, log.Named("Account"))
	}
	return r, nil
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"OK":   true,
		"time": time.Now().UTC().Format(time.RFC3339),
	})
}
