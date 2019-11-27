package db

import (
	"github.com/unectio/util/mongo"
	"gopkg.in/mgo.v2/bson"
)

type KeyDb struct {
	Id		bson.ObjectId		`bson:"_id,omitempty"`
	Kind		string			`bson:"kind"`
	Value		[]byte			`bson:"value"`

	User		bson.ObjectId		`bson:"user,omitempty"`
	Scope		*KeyScopeDb		`bson:"scope,omitempty"`
}

type KeyScopeDb struct {
	Project		bson.ObjectId		`bson:"project,omitempty"`
	RemoteAddr	string			`bson:"remote_addr,omitempty"`
}

func (k *KeyDb)ByKind(kind string) bson.M {
	/* .Kind */
	return bson.M{"kind": kind}
}

func (k *KeyDb)ID() bson.ObjectId { return k.Id }
func (k *KeyDb)Location() * mongo.Location { return LocKeys }

const (
	/*
	 * Self-signed key for apilet
	 */
	KeyKindSelf string		= "jwt.self"
	/*
	 * Key signed by managlet
	 */
	KeyKindServer string		= "jwt.server"
	/*
	 * Key used to encrypt secrets in DB
	 */
	KeyKindSecret string		= "apilet.secrets"
	/*
	 * Key to hmac-sign tokens by apilet
	 */
	KeyKindAccess string		= "apilet.access"
)
