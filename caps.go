package db

import (
	"github.com/unectio/util/mongo"
	"gopkg.in/mgo.v2/bson"
)

type CapsDb struct {
	Id		bson.ObjectId		`bson:"_id,omitempty"`
	Role		string			`bson:"role"`
	Caps		[]string		`bson:"caps"`
}

func (r *CapsDb)ID() bson.ObjectId { return r.Id }
func (r *CapsDb)Location() *mongo.Location { return LocCaps }
