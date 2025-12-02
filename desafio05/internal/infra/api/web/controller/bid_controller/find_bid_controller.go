package bidcontroller

import (
	"context"
	"net/http"

	resterr "github.com/gbuenodev/fullcycle_go_expert/desafio05/configs/rest_err.go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (bc *BidController) FindBidByAuctionId(c *gin.Context) {
	auctionId := c.Param("auctionId")
	if err := uuid.Validate(auctionId); err != nil {
		errRest := resterr.NewBadRequestError("Invalid fields", resterr.Causes{
			Field:   "auctionId",
			Message: "Invalid UUID value",
		})

		c.JSON(errRest.Code, errRest)
		return
	}

	bids, err := bc.BidUseCase.FindBidByAuctionId(context.Background(), auctionId)
	if err != nil {
		errRest := resterr.ConvertError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, bids)
}
