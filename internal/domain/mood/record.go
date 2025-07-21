package moodDmn

import "go.mongodb.org/mongo-driver/v2/bson"

type Mood struct {
	Id   bson.ObjectID `json:"-" bson:"_id,omitempty"`
	Name string        `json:"name" bson:"name"`
}
