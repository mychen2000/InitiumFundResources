package position

import (
	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Position struct{
	Order   *alpaca.Order
	Time    *time.Time
	Account *alpaca.Account
}

var Positions []Position

func (pc *PositionChangeController) positionChange(c *gin.Context){
	placeOrder, err := pc.alpacaClient.PlaceOrder(*pc.req)

	if err != nil {
		pc.logger.Error(err)

		data := map[string]interface{}{
			"err": err.Error(),
		}
		c.JSON(http.StatusServiceUnavailable, data)
		return
	}

	account, errAcc := pc.alpacaClient.GetAccount()

	if errAcc != nil {
		pc.logger.Error(errAcc)

		data := map[string]interface{}{
			"err": errAcc.Error(),
		}
		c.JSON(http.StatusServiceUnavailable, data)
		return
	}

	date := time.Now()

	currentPosition := Position{
		Order:   placeOrder,
		Time:    &date,
		Account: account,
	}

	Positions[len(Positions)] = currentPosition
	c.JSON(http.StatusOK, Positions)
}

func getPosition(c *gin.Context){
	c.JSON(http.StatusOK, Positions)
}