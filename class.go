package db

import (
	"github.com/unectio/util/mongo"
	"gopkg.in/mgo.v2/bson"
)

type ClassDb struct {
	Id		bson.ObjectId		`bson:"_id,omitempty"`
	Name		string			`bson:"name"`
}

func (k *ClassDb)ID() bson.ObjectId { return k.Id }
func (k *ClassDb)Location() * mongo.Location { return LocClasses }
