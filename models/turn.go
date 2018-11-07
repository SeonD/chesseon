package models

import "gopkg.in/mgo.v2/bson"

const (
	// CollectionTurn : Collection name for turns
	CollectionTurn = "turns"
)

// Turn model
type Turn struct {
	ID        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Game      bson.ObjectId `json:"game" bson:"game"`
	Player    bson.ObjectId `json:"player" bson:"player"`
	Move      string        `json:"move" bson:"move"`
	TurnStart int64         `json:"turn_start" bson:"turn_start"`
	TurnEnd   int64         `json:"turn_end" bson:"turn_end"`
	Board     Board         `json:"board" bson:"board"`
}

// Board model
type Board struct {
	Pieces map[string]string `json:"pieces" bson:"pieces"`
}

// Move model
type Move struct {
	Player     bson.ObjectId `json:"player_id" bson:"player_id"`
	PrivateKey bson.ObjectId `json:"private_key" bson:"private_key"`
	Move       string        `json:"move" bson:"move"`
}
