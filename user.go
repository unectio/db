package db

import (
	"time"
	"github.com/unectio/util/mongo"
	"gopkg.in/mgo.v2/bson"
)

type UserDb struct {
	Id		bson.ObjectId		`bson:"_id,omitempty"`
	Email		string			`bson:"email"`
	Name		string			`bson:"name"`
	Role		string			`bson:"role"`
	Source		string			`bson:"source"`
	Created		time.Time		`bson:"created"`
}

func (u *UserDb)ID() bson.ObjectId { return u.Id }
func (u *UserDb)Location() *mongo.Location { return LocUsers }
