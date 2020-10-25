package weeklyReport

import (
	"github.com/gin-gonic/gin"
	"github.com/initiumfund/artis-go/pkg/position"
	"net/http"
	"time"
)

var WeeklyReport []position.Position

func (wrc *WeeklyReportController) getWeeklyReport(c *gin.Context){
	currentTime := wrc.today
	for i:= 0; i < len(position.Positions); i++{
		date := position.Positions[i].Time
		if currentTime.Sub(*date) < 24*7*time.Hour{
			WeeklyReport[len(WeeklyReport)] = position.Positions[i]
		}
	}

	c.JSON(http.StatusOK, WeeklyReport)
}