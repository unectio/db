package db

import (
	"github.com/unectio/util/mongo"
	"gopkg.in/mgo.v2/bson"
)

type ProjectDb struct {
	Id		bson.ObjectId		`bson:"_id,omitempty"`
	Name					`bson:",inline"`
	Compute		ComputeDb		`bson:"compute"`
	UserData	string			`bson:"userdata,omitempty"`
}

func (p *ProjectDb)ID() bson.ObjectId { return p.Id }
func (p *ProjectDb)Location() *mongo.Location { return LocProjects }

const (
	DefaultProjectId string		= "0"
	DefaultProjectName string	= "default"
)
