/////////////////////////////////////////////////////////////////////////////////
//
// Copyright (C) 2019-2020, Unectio Inc, All Right Reserved.
//
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
// ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
/////////////////////////////////////////////////////////////////////////////////

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

	Reveal		map[string]string	`bson:"reveal"`
}

func (s *SecretDb)ID() bson.ObjectId { return s.Id }
func (s *SecretDb)Location() *mongo.Location { return LocSecret }

func (s *SecretDb)PayloadUpdReq() bson.M {
	/* .KeyId and .Data */
	return bson.M{"keyid": s.KeyId, "data": s.Data}
}

func (s *SecretDb)SavePayload(ctx context.Context, pl map[string]string) error {
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

func (s *SecretDb)LoadPayload(ctx context.Context) (map[string]string, error) {
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

