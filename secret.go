package db

import (
	"context"
	"github.com/unectio/util"
	"encoding/json"
	"github.com/unectio/util/mongo"
	"gopkg.in/mgo.v2/bson"
)

type SecretDb struct {
	DbCommon				`bson:",inline"`

	KeyId		bson.ObjectId		`bson:"keyid"`
	Data		[]byte			`bson:"data"`
}

func (s *SecretDb)ID() bson.ObjectId { return s.Id }
func (s *SecretDb)Location() *mongo.Location { return LocSecret }

func (s *SecretDb)PayloadUpdReq() bson.M {
	/* .KeyId and .Data */
	return bson.M{"keyid": s.KeyId, "data": s.Data}
}

func (s *SecretDb)Save(ctx context.Context, pl map[string]string) error {
	data, err := json.Marshal(pl)
	if err != nil {
		return err
	}

	var key KeyDb

	err = Find(ctx, key.ByKind(KeyKindSecret), &key)
	if err != nil {
		return util.Error("cannot find key", err)
	}

	cdata, err := util.Encrypt(key.Value, data)
	if err != nil {
		return util.Error("cannot encrypt", err)
	}

	s.KeyId = key.Id
	s.Data = cdata

	return nil
}

func (s *SecretDb)Load(ctx context.Context) (map[string]string, error) {
	var key KeyDb

	err := Load(ctx, s.KeyId, &key)
	if err != nil {
		return nil, util.Error("cannot find key", err)
	}

	data, err := util.Decrypt(key.Value, s.Data)
	if err != nil {
		return nil, util.Error("cannot decrypt", err)
	}

	var kv map[string]string
	err = json.Unmarshal(data, &kv)
	if err != nil {
		return nil, util.Error("cannot unmarshal", err)
	}

	return kv, nil
}

