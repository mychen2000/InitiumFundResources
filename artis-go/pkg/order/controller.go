package order

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t *PlaceController) placeOrder(c *gin.Context) {
	placeOrder, err := t.alpacaClient.PlaceOrder(*t.req)

	if err != nil {
		t.logger.Error(err)

		data := map[string]interface{}{
			"err": err.Error(),
		}
		c.JSON(http.StatusServiceUnavailable, data)
		return
	}

	c.JSON(http.StatusOK, *placeOrder)
}

func (t *GetController) getOrder(c *gin.Context){
	getOrder, err := t.alpacaClient.GetOrder(*t.orderID)

	if err != nil {
		t.logger.Error(err)

		data := map[string]interface{}{
			"err": err.Error(),
		}
		c.JSON(http.StatusServiceUnavailable, data)
		return
	}

	c.JSON(http.StatusOK, *getOrder)
}