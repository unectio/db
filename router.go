package db

import (
	"github.com/unectio/util/mongo"
	"gopkg.in/mgo.v2/bson"
)

type RouterDb struct {
	DbCommon				`bson:",inline"`

	Compute		ComputeDb		`bson:"compute'`

	URLType		string			`bson:"url_type"`
	URL		string			`bson:"url"`
	Auth		bson.ObjectId		`bson:"auth,omitempty"`

	Mux		[]*RouteRuleDb		`bson:"mux,omitempty"`
}

func (rd *RouterDb)UpdateMuxQ() bson.M {
	/* .Mux */
	return bson.M{"mux": rd.Mux}
}

const (
	RouterURLLocal	= "local"
	RouterURLDomain	= "domain"
)

func (rt *RouterDb)ID() bson.ObjectId { return rt.Id }
func (rt *RouterDb)Location() *mongo.Location { return LocRouter }

type RouteRuleDb struct {
	Methods		string			`bson:"methods"` /* ,-separated */
	Path		string			`bson:"path"`
	Key		string			`bson:"key"`
	FnId		bson.ObjectId		`bson:"fnid"`
}
