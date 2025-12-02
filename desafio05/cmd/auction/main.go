package main

import (
	"context"
	"log"

	"github.com/gbuenodev/fullcycle_go_expert/desafio05/configs/database/mongodb"
	auctioncontroller "github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/infra/api/web/controller/auction_controller"
	bidcontroller "github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/infra/api/web/controller/bid_controller"
	usercontroller "github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/infra/api/web/controller/user_controller"
	"github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/infra/repository/auction"
	"github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/infra/repository/bid"
	"github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/infra/repository/user"
	auctionusecase "github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/usecase/auction_usecase"
	bidusecase "github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/usecase/bid_usecase"
	userusecase "github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/usecase/user_usecase"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	ctx := context.Background()

	if err := godotenv.Load("cmd/auction/.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
		return
	}

	db, err := mongodb.NewMongoDBConnection(ctx)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer db.Client().Disconnect(ctx)

	router := gin.Default()

	userController, bidController, auctionController := initDependencies(db)

	router.GET("/auctions", auctionController.FindAuctions)
	router.GET("/auctions/:auctionId", auctionController.FindAuctionById)
	router.POST("/auctions", auctionController.CreateAuction)
	router.GET("/auctions/winner/:auctionId", auctionController.FindWinningBidByAuctionId)
	router.POST("/bid", bidController.CreateBid)
	router.GET("/bid/:auctionId", bidController.FindBidByAuctionId)
	router.GET("/user/:userId", userController.FindUserById)
	router.POST("/user", userController.CreateUser)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	router.Run(":8080")
}

func initDependencies(db *mongo.Database) (
	userController *usercontroller.UserController,
	bidController *bidcontroller.BidController,
	auctionController *auctioncontroller.AuctionController,
) {
	auctionRepository := auction.NewAuctionRepository(db)
	bidRepository := bid.NewBidRepository(db, auctionRepository)
	userRepository := user.NewUserRepository(db)

	userController = usercontroller.NewUserController(userusecase.NewUserUseCase(userRepository))
	auctionController = auctioncontroller.NewAuctionController(auctionusecase.NewAuctionUseCase(auctionRepository, bidRepository))
	bidController = bidController.NewBidController(bidusecase.NewBidUseCase(bidRepository))

	return userController, bidController, auctionController
}
