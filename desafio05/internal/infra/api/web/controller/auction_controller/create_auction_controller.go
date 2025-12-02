package auctioncontroller

import (
	"context"
	"net/http"

	resterr "github.com/gbuenodev/fullcycle_go_expert/desafio05/configs/rest_err.go"
	"github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/infra/api/web/validation"
	auctionusecase "github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/usecase/auction_usecase"
	"github.com/gin-gonic/gin"
)

type AuctionController struct {
	AuctionUseCase auctionusecase.AuctionUseCaseInterface
}

func NewAuctionController(auctionUseCase auctionusecase.AuctionUseCaseInterface) *AuctionController {
	return &AuctionController{
		AuctionUseCase: auctionUseCase,
	}
}

func (ac *AuctionController) CreateAuction(c *gin.Context) {
	var auctionInputDTO auctionusecase.AuctionInputDTO

	if err := c.ShouldBindJSON(&auctionInputDTO); err != nil {
		errRest := validation.ValidateErr(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	err := ac.AuctionUseCase.CreateAuction(context.Background(), &auctionInputDTO)
	if err != nil {
		errRest := resterr.ConvertError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.Status(http.StatusCreated)
}
