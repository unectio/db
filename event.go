package db

import (
	"time"
	"github.com/unectio/api"
	"github.com/unectio/util/mongo"
	"gopkg.in/mgo.v2/bson"
)

type EventSourceDb struct {
	DbCommon				`bson:",inline"`

	Pull		*EventPullDb		`bson:"pull,omitempty"`
}

func (o *EventSourceDb)ID() bson.ObjectId { return o.Id }
func (o *EventSourceDb)Location() *mongo.Location { return LocEventSources }

func (es *EventSourceDb)DbKey() string {
	return "event::" + es.Id.Hex()
}

type EventPullDb struct {
	Host		string			`bson:"host"`
	Path		string			`bson:"path"`
}

type EventTrigDb struct {
	Src		bson.ObjectId		`bson:"src"`
}

type DeferEventDb struct {
	Id		bson.ObjectId		`bson:"_id"`
	Fn		bson.ObjectId		`bson:"fnid"`
	After		time.Time		`bson:"after"`
	Req		*api.RunRequest		`bson:"req"`
}

func (o *DeferEventDb)ID() bson.ObjectId { return o.Id }
func (o *DeferEventDb)Location() *mongo.Location { return LocDeferEvs }
