package usercontroller

import (
	"context"
	"net/http"

	resterr "github.com/gbuenodev/fullcycle_go_expert/desafio05/configs/rest_err.go"
	"github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/infra/api/web/validation"
	userusecase "github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/usecase/user_usecase"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUseCase userusecase.UserUsecaseInterface
}

func NewUserController(userUseCase userusecase.UserUsecaseInterface) *UserController {
	return &UserController{
		UserUseCase: userUseCase,
	}
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var userInputDTO userusecase.UserInputDTO

	if err := c.ShouldBindJSON(&userInputDTO); err != nil {
		errRest := validation.ValidateErr(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	output, err := uc.UserUseCase.CreateUser(context.Background(), &userInputDTO)
	if err != nil {
		errRest := resterr.ConvertError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusCreated, output)
}
