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
