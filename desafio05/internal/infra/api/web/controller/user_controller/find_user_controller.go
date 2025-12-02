package usercontroller

import (
	"context"
	"net/http"

	resterr "github.com/gbuenodev/fullcycle_go_expert/desafio05/configs/rest_err.go"
	userusecase "github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/usecase/user_usecase"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct {
	UserUseCase userusecase.UserUsecaseInterface
}

func NewUserController(userUseCase userusecase.UserUsecaseInterface) *UserController {
	return &UserController{
		UserUseCase: userUseCase,
	}
}

func (uc *UserController) FindUserById(c *gin.Context) {
	userId := c.Param("userId")
	if err := uuid.Validate(userId); err != nil {
		errRest := resterr.NewBadRequestError("Invalid fields", resterr.Causes{
			Field:   "userId",
			Message: "Invalid UUID value",
		})

		c.JSON(errRest.Code, errRest)
		return
	}

	user, err := uc.UserUseCase.FindUserById(context.Background(), userId)
	if err != nil {
		errRest := resterr.ConvertError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, user)
}
