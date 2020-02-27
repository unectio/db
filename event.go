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
