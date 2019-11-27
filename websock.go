package db

import (
	"github.com/unectio/util/mongo"
	"gopkg.in/mgo.v2/bson"
)

type WebsockDb struct {
	DbCommon				`bson:",inline"`

	AccToken	string			`bson:"acc_token"`
	Auth		bson.ObjectId		`bson:"auth,omitempty"`

	Compute		ComputeDb		`bson:"compute"`
}

func (o *WebsockDb)ID() bson.ObjectId { return o.Id }
func (o *WebsockDb)Location() *mongo.Location { return LocWebsock }

type WsTrigDb struct {
	Id		bson.ObjectId		`bson:"id"`
}

func (ws *WebsockDb)DbKey() string {
	return "websock::" + ws.Id.Hex()
}
