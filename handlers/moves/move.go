package moves

import (
	"net/http"

	"github.com/SeonD/chesseon/handlers/games"

	"github.com/SeonD/chesseon/models"
	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// PlayMove a player
func PlayMove(c *gin.Context) {
	oID := bson.ObjectIdHex(c.Param("_id"))

	db := c.MustGet("db").(*mgo.Database)

	turn := models.Turn{}
	err := db.C(models.CollectionTurn).FindId(oID).One(&turn)
	if err != nil {
		c.Error(err)
		return
	}

	// Validate player
	move := models.Move{}
	err = c.Bind(&move)
	if err != nil {
		c.Error(err)
		return
	}
	if move.Player != turn.Player {
		c.Status(http.StatusUnauthorized)
		return
	}
	player := models.Player{}
	err = db.C(models.CollectionPlayer).FindId(move.Player).One(&player)
	if err != nil {
		c.Error(err)
		return
	}
	if move.PrivateKey != player.PrivateKey {
		c.Status(http.StatusUnauthorized)
		return
	}

	game := models.Game{}
	err = db.C(models.CollectionGame).FindId(turn.Game).One(&game)
	if err != nil {
		c.Error(err)
		return
	}

	// TODO: validate move

	err = db.C(models.CollectionTurn).UpdateId(oID, bson.M{
		"$set": bson.M{
			"move": move.Move,
		},
	})
	if err != nil {
		c.Error(err)
		return
	}

	// TODO : Process new board state
	// TODO : Process game result if the game ends

	otherPlayerID := game.Players.Black
	if player.ID == otherPlayerID {
		otherPlayerID = game.Players.White
	}

	newTurn := games.GetNewTurn(game.ID, otherPlayerID)
	newTurn.Board.Pieces = turn.Board.Pieces
	err = db.C(models.CollectionTurn).Insert(&newTurn)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"board": turn.Board.Pieces,
	})
}

// GetByID looks for a player by id
/*func GetByID(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	player := models.Player{}
	oID := bson.ObjectIdHex(c.Param("_id"))
	err := db.C(models.CollectionPlayer).FindId(oID).One(&player)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"_id":  player.ID,
		"name": player.Name,
	})
}
*/
