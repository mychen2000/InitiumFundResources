package accountActivities

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t *Controller) getAccountActivities(c *gin.Context){
	accountActivities, err := t.alpacaClient.GetAccountActivities(t.activityType, t.opts) //这里读了Controller中的值，同时也给到了Controller没有被定义的情况

	if err != nil { //当err发生的时候执行
		t.logger.Error(err)

		data := map[string]interface{}{
			"err": err.Error(),
		}
		c.JSON(http.StatusServiceUnavailable, data)
		return
	}

	c.JSON(http.StatusOK, accountActivities)
}