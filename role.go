package db

import (
	"github.com/unectio/util/mongo"
	"gopkg.in/mgo.v2/bson"
)

type RoleDb struct {
	Id		bson.ObjectId		`bson:"_id,omitempty"`
	ProjectId	bson.ObjectId		`bson:"project"`
	UserId		bson.ObjectId		`bson:"user"`
	Role		string			`bson:"role"`
}

func (r *RoleDb)ID() bson.ObjectId { return r.Id }
func (r *RoleDb)Location() *mongo.Location { return LocRoles }
