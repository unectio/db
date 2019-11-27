package db

import (
	"github.com/unectio/util/mongo"
	"gopkg.in/mgo.v2/bson"
)

type AppDb struct {
	Id		bson.ObjectId		`bson:"_id,omitempty"`
	Name					`bson:",inline"`
	Compute		ComputeDb		`bson:"compute"`

	Funcs		[]bson.ObjectId		`bson:"funcs"`
	Router		*bson.ObjectId		`bson:"router,omitempty"`
}

func (a *AppDb)ID() bson.ObjectId { return a.Id }
func (a *AppDb)Location() *mongo.Location { return LocApps }

