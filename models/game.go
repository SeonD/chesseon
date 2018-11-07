package models

import "gopkg.in/mgo.v2/bson"

const (
	// CollectionGame : Collection name for games
	CollectionGame = "games"
)

// Game model
type Game struct {
	ID        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Players   Players       `json:"players" bson:"players"`
	CreatedOn int64         `json:"created_on" bson:"created_on"`
	EndedOn   int64         `json:"ended_on" bson:"ended_on"`
}

// Players in a game
type Players struct {
	Black bson.ObjectId `json:"black" binding:"required" bson:"black"`
	White bson.ObjectId `json:"white" binding:"required" bson:"white"`
}
