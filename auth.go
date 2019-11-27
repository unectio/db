package db

import (
	"github.com/unectio/util/mongo"
	"gopkg.in/mgo.v2/bson"
)

type AuthMethodDb struct {
	DbCommon				`bson:",inline"`

	JWT		*AuthJWTDb		`bson:"jwt,omitempty"`
	Platform	bool			`bson:"platform,omitempty"`
}

func (am *AuthMethodDb)UpdateJWTQ() bson.M {
	/* .JWT */
	return bson.M{"jwt": am.JWT}
}

func (ac *AuthMethodDb)ID() bson.ObjectId { return ac.Id }
func (ac *AuthMethodDb)Location() *mongo.Location { return LocAuthCtx }

type AuthJWTDb struct {
	Key		[]byte			`bson:"key,omitempty"`	/* FIXME -- encrypt */
}
