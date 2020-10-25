package account

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t *Controller) getAccount(c *gin.Context) { //在这个情况下，Controller中的alpacaClient和Logger的值就会被读进来
	account, err := t.alpacaClient.GetAccount() //这里读了Controller中的值，同时也给到了Controller没有被定义的情况

	if err != nil { //当err发生的时候执行
		t.logger.Error(err)

		data := map[string]interface{}{
			"err": err.Error(), //把Error()的返回值填到data这个map里面，做一个记录
		}
		c.JSON(http.StatusServiceUnavailable, data) //这里JSON的返回值是StatusServiceUnavailable，明确赋值data这个map
		return //直接返回不会进行下一步
	}

	c.JSON(http.StatusOK, *account) //当没有err的时候，会跳过上一步，直接返回account的信息并形成一个JSON文档
									//这里JSON的返回值是StatusOK，并且赋值account的内容
}
