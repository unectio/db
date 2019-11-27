package db

import (
	"github.com/unectio/util/mongo"
	"gopkg.in/mgo.v2/bson"
)

type MongoDb struct {
	DbCommon				`bson:",inline"`

	State		string			`bson:"state"`
	Tier		string			`bson:"tier"`

	Creds		*SimpleMwareCredsDb	`bson:"creds"`
	Instance	bson.ObjectId		`bson:"instance"`
}

func (o *MongoDb)ID() bson.ObjectId { return o.Id }
func (o *MongoDb)Location() *mongo.Location { return LocMongo }

func MgoInstanceQ(tier string) bson.M {
	/* MwareServiceDb.Type & MwareServiceDb.Mgo.Tier */
	return bson.M{"type": "mongo", "mgo.tier": tier}
}

type MongoServiceDb struct {
	Tier		string			`bson:"tier"`

	Admin		*SimpleMwareCredsDb	`bson:"admin_creds"`
}
