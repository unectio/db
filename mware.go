package db

import (
	"fmt"
	"github.com/unectio/util/mongo"
	"gopkg.in/mgo.v2/bson"
)

/*
 * Describes simple creds to access a generic mware server
 *
 * @domain -- a virtual sub-entity within the server (DB for
 *            mongo/maria, vhost for rabbitmq, etc.
 * @user/@pass -- user and, well, password
 */

type SimpleMwareCredsDb struct {
	Address		string			`bson:"address"`
	Domain		string			`bson:"domain"`
	User		string			`bson:"user"`
	Pass		string			`bson:"pass"`
}

func (c *SimpleMwareCredsDb)String() string {
	return fmt.Sprintf("%s:%s@%s/%s", c.User, c.Pass, c.Address, c.Domain)
}

/*
 * Admin-managed entry describing an mware service hanging aroung.
 */
type MwareServiceDb struct {
	Id		bson.ObjectId		`bson:"_id,omitempty"`
	Type		string			`bson:"type"`
	Desc		string			`bson:"description"`

	Mgo		*MongoServiceDb		`bson:"mongo"`
}

func (_ *MwareServiceDb)TypeQ(typ string) bson.M {
	return bson.M{"type": typ}
}

func (o *MwareServiceDb)ID() bson.ObjectId { return o.Id }
func (o *MwareServiceDb)Location() *mongo.Location { return LocMwares }
