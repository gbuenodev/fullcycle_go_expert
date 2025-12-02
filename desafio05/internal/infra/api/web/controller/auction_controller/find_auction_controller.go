package auctioncontroller

import (
	"context"
	"net/http"
	"strconv"

	resterr "github.com/gbuenodev/fullcycle_go_expert/desafio05/configs/rest_err.go"
	auctionusecase "github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/usecase/auction_usecase"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (ac *AuctionController) FindAuctionById(c *gin.Context) {
	auctionId := c.Param("auctionId")
	if err := uuid.Validate(auctionId); err != nil {
		errRest := resterr.NewBadRequestError("Invalid fields", resterr.Causes{
			Field:   "auctionId",
			Message: "Invalid UUID value",
		})

		c.JSON(errRest.Code, errRest)
		return
	}

	auction, err := ac.AuctionUseCase.FindAuctionById(context.Background(), auctionId)
	if err != nil {
		errRest := resterr.ConvertError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, auction)
}

func (ac *AuctionController) FindAuctions(c *gin.Context) {
	status := c.Query("status")
	category := c.Query("category")
	productName := c.Query("productName")

	statusEnum, errConv := strconv.Atoi(status)
	if errConv != nil {
		errRest := resterr.NewBadRequestError("Error trying to validate auction status param")
		c.JSON(errRest.Code, errRest)
		return
	}

	auctions, err := ac.AuctionUseCase.FindAuctions(context.Background(), auctionusecase.AuctionStatus(statusEnum), category, productName)
	if err != nil {
		errRest := resterr.ConvertError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, auctions)
}

func (ac *AuctionController) FindWinningBidByAuctionId(c *gin.Context) {
	auctionId := c.Param("auctionId")
	if err := uuid.Validate(auctionId); err != nil {
		errRest := resterr.NewBadRequestError("Invalid fields", resterr.Causes{
			Field:   "auctionId",
			Message: "Invalid UUID value",
		})

		c.JSON(errRest.Code, errRest)
		return
	}

	auction, err := ac.AuctionUseCase.FindWinningBidByAuctionId(context.Background(), auctionId)
	if err != nil {
		errRest := resterr.ConvertError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, auction)
}
