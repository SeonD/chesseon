package main

import (
	"net/http"

	"github.com/SeonD/chesseon/db"
	"github.com/SeonD/chesseon/handlers/games"
	"github.com/SeonD/chesseon/handlers/moves"
	"github.com/SeonD/chesseon/handlers/players"
	"github.com/SeonD/chesseon/middlewares"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

const (
	// Port that runs the app
	Port = "3030"
)

func init() {
	db.Connect()
}

func main() {
	router := gin.Default()

	router.Use(middlewares.Connect)
	router.Use(middlewares.ErrorHandler)
	router.Use(middlewares.SetConstants)
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		player := api.Group("/players")
		{
			player.POST("/", players.Create)
			player.GET("/:_id", players.GetByID)
		}

		game := api.Group("/games")
		{
			game.POST("/", games.Create)
			game.GET("/:_id/turn/:player_id", games.GetTurnByGameAndPlayerID)
			game.GET("/:_id", games.GetByID)
		}

		move := api.Group("/moves")
		{
			move.POST("/:_id", moves.PlayMove)
		}
	}

	router.Run(":3030")
}
