package games

import (
	"net/http"
	"time"

	"github.com/SeonD/chesseon/models"
	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Create a game
func Create(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	game := models.Game{}
	err := c.Bind(&game)
	if err != nil {
		c.Error(err)
		return
	}

	game.ID = bson.NewObjectId()
	game.CreatedOn = time.Now().UnixNano() / int64(time.Millisecond)

	err = db.C(models.CollectionGame).Insert(&game)
	if err != nil {
		c.Error(err)
		return
	}

	turn := GetNewTurn(game.ID, game.Players.Black)
	turn.Board.Pieces = c.MustGet("boardSetup").(map[string]string)

	err = db.C(models.CollectionTurn).Insert(&turn)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"gameId":      game.ID.Hex(),
		"firstMoveId": turn.ID.Hex(),
	})
}

// GetByID gets a game by object id
func GetByID(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	game := models.Game{}
	oID := bson.ObjectIdHex(c.Param("_id"))
	err := db.C(models.CollectionGame).FindId(oID).One(&game)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"_id":        game.ID,
		"players":    game.Players,
		"created_on": game.CreatedOn,
		"ended":      game.EndedOn > 0,
	})
}

// GetTurnByGameAndPlayerID gets a game by object id
func GetTurnByGameAndPlayerID(c *gin.Context) {
	gameObjectID := bson.ObjectIdHex(c.Param("_id"))
	playerObjectID := bson.ObjectIdHex(c.Param("player_id"))

	db := c.MustGet("db").(*mgo.Database)

	turn := models.Turn{}
	err := db.C(models.CollectionTurn).Find(bson.M{
		"game": gameObjectID,
	}).Sort("-_id").One(&turn)
	if err != nil {
		c.Error(err)
		return
	}

	if turn.Player != playerObjectID {
		c.Status(http.StatusLocked)
		return
	}

	// update timestamps only when it's unset
	timeNow := time.Now().UnixNano() / int64(time.Millisecond)

	err = db.C(models.CollectionTurn).UpdateId(turn.ID, bson.M{
		"$set": bson.M{
			"turn_start": timeNow,
			"turn_end":   timeNow + 5000,
		},
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"board":      turn.Board.Pieces,
		"turnStart":  timeNow,
		"turnEnd":    timeNow + 5000,
		"nextMoveId": turn.ID,
	})
}
