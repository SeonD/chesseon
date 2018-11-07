package games

import (
	"github.com/SeonD/chesseon/models"
	"gopkg.in/mgo.v2/bson"
)

// GetNewTurn creates a turn
func GetNewTurn(Game bson.ObjectId, Player bson.ObjectId) models.Turn {
	turn := models.Turn{}

	turn.ID = bson.NewObjectId()
	turn.Game = Game
	turn.Player = Player

	return turn
}
