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

	"github.com/unectio/util/mongo"
	"gopkg.in/mgo.v2/bson"
)

type RepoDb struct {
	DbCommon `bson:",inline"`

	State  string    `bson:"state"`
	Synced time.Time `bson:"synced"`

	Sync       string `bson:"sync"`
	SyncDelayM int    `bson:"sync_delay_m,omitempty"`

	Type string `bson:"type"`
	URL  string `bson:"url"`
}

func (rp *RepoDb) UpdateSyncedRq(q bson.M) bson.M {
	if q == nil {
		q = bson.M{}
	}

	/* .Synced */
	q["synced"] = time.Now()

	return q
}

func (rp *RepoDb) BySyncType(typ string) bson.M {
	/* .Sync */
	return bson.M{"sync": typ}
}

func (rp *RepoDb) ID() bson.ObjectId         { return rp.Id }
func (rp *RepoDb) Location() *mongo.Location { return LocRepo }

const (
	SyncTimer = "timer"
)

func (rp *RepoDb) SyncDelay() time.Duration {
	return time.Duration(rp.SyncDelayM) * time.Minute
}
