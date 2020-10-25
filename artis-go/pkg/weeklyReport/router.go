package weeklyReport

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

type WeeklyReportController struct {
	today  *time.Time
	logger *zap.SugaredLogger
}

func ActivitiesRouter(r gin.IRouter, db *gorm.DB, today *time.Time,logger *zap.SugaredLogger) {
	c := &WeeklyReportController{
		today:  today,
		logger: logger,
	}

	r.GET("", c.getWeeklyReport)
}