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
	Address string `bson:"address"`
	Domain  string `bson:"domain"`
	User    string `bson:"user"`
	Pass    string `bson:"pass"`
}

func (c *SimpleMwareCredsDb) String() string {
	return fmt.Sprintf("%s:%s@%s/%s", c.User, c.Pass, c.Address, c.Domain)
}

/*
 * Admin-managed entry describing an mware service hanging aroung.
 */
type MwareServiceDb struct {
	Id   bson.ObjectId `bson:"_id,omitempty"`
	Type string        `bson:"type"`
	Desc string        `bson:"description"`

	Mgo *MongoServiceDb `bson:"mongo"`
}

func (_ *MwareServiceDb) TypeQ(typ string) bson.M {
	return bson.M{"type": typ}
}

func (o *MwareServiceDb) ID() bson.ObjectId         { return o.Id }
func (o *MwareServiceDb) Location() *mongo.Location { return LocMwares }
