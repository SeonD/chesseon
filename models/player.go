package models

import "gopkg.in/mgo.v2/bson"

const (
	// CollectionPlayer : Collection name for players
	CollectionPlayer = "players"
)

// Player model
type Player struct {
	ID   bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string        `json:"name" form:"name" binding:"required" bson:"name"`
}
