package players

import (
	"net/http"

	"github.com/SeonD/chesseon/models"
	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Create a player
func Create(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	player := models.Player{}
	err := c.Bind(&player)
	if err != nil {
		c.Error(err)
		return
	}

	player.ID = bson.NewObjectId()
	player.PrivateKey = bson.NewObjectId()

	err = db.C(models.CollectionPlayer).Insert(player)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"playerId":   player.ID,
		"privateKey": player.PrivateKey,
	})
}

// GetByID looks for a player by id
func GetByID(c *gin.Context) {
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
