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
	"github.com/unectio/util/mongo"
	"gopkg.in/mgo.v2/bson"
)

type FnTargetDb struct {
	DbCommon				`bson:",inline"`

	FnId		bson.ObjectId		`bson:"fnid"`
	Fn		*NextFunctionDb		`bson:"function,omitempty"`
	Rest		*NextRestDb		`bson:"rest,omitempty"`
	Mware		*NextMwareDb		`bson:"mware,omitempty"`

	SuccFn		bson.ObjectId		`bson:"on_success"`
	FailFn		bson.ObjectId		`bson:"on_failure"`
}

func (tg *FnTargetDb)ByFn(fnid bson.ObjectId) bson.M {
	/* .FnId */
	return bson.M{"fnid": fnid}
}

func (tg *FnTargetDb)ID() bson.ObjectId { return tg.Id }
func (tg *FnTargetDb)Location() *mongo.Location { return LocTarget }

type NextFunctionDb struct {
	FnId		bson.ObjectId		`bson:"fnid"`
}

type NextRestDb struct {
	URL		string			`bson:"url"`
}

type NextMwareDb struct {
}
