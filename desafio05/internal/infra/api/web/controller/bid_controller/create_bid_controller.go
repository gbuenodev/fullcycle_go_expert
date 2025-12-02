package bidcontroller

import (
	"context"
	"net/http"

	resterr "github.com/gbuenodev/fullcycle_go_expert/desafio05/configs/rest_err.go"
	"github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/infra/api/web/validation"
	bidusecase "github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/usecase/bid_usecase"
	"github.com/gin-gonic/gin"
)

type BidController struct {
	BidUseCase bidusecase.BidUseCaseInterface
}

func (bc *BidController) NewBidController(bidUseCase bidusecase.BidUseCaseInterface) *BidController {
	return &BidController{
		BidUseCase: bidUseCase,
	}
}

func (bc *BidController) CreateBid(c *gin.Context) {
	var bidInputDTO bidusecase.BidInputDTO

	if err := c.ShouldBindJSON(&bidInputDTO); err != nil {
		errRest := validation.ValidateErr(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	err := bc.BidUseCase.CreateBid(context.Background(), &bidInputDTO)
	if err != nil {
		errRest := resterr.ConvertError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.Status(http.StatusCreated)
}
