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

type FunctionDb struct {
	DbCommon				`bson:",inline"`

	State		string			`bson:"state"`

	Compute		ComputeDb		`bson:"compute"`
	Limits		FuncLimitsDb		`bson:"limits"`
	Env		[]*EnvValDb		`bson:"env"`
	Targets		map[string]*FnTargetDb	`bson:"chain_targets"`
	CodeBalancer	string			`bson:"code_balancer"`

	Gen		int			`bson:"gen"`
}

func (f *FunctionDb)UpdateEnvQ() bson.M {
	/* .Env */
	return bson.M{"env": f.Env, "gen": f.Gen}
}

func (f *FunctionDb)UpdateLimQ() bson.M {
	/* .Limits */
	return bson.M{"limits": &f.Limits, "gen": f.Gen}
}

func (f *FunctionDb)UpdateTargetsQ() bson.M {
	/* .Targets */
	return bson.M{"chain_targets": f.Targets, "gen": f.Gen}
}

func (f *FunctionDb)ID() bson.ObjectId { return f.Id }
func (f *FunctionDb)Location() *mongo.Location { return LocFunc }

type EnvValDb struct {
	Name		string			`bson:"name"`
	Value		string			`bson:"value"`

	/*
	 * This is the resolved reference, which is not to be
	 * kept in DB.
	 */
	resolved	string			`bson:"-"`
}

func (ev *EnvValDb)Resolve(v string) {
	ev.resolved = v
}

func (ev *EnvValDb)RealValue() string {
	if ev.resolved != "" {
		return ev.resolved
	} else {
		return ev.Value
	}
}

type FuncLimitsDb struct {
	TmoMsec		int			`bson:"tmo_msec"`
	Burst		int			`bson:"burst"`
	Rate		int			`bson:"rate"`
	Class		string			`bson:"class"`
}

func (l *FuncLimitsDb)Timeout() time.Duration {
	return time.Duration(l.TmoMsec) * time.Millisecond
}

func (l *FuncLimitsDb)RL() (uint, uint) {
	return uint(l.Rate), uint(l.Burst)
}

func FnLogKey(fnid bson.ObjectId) string {
	return "fn." + fnid.Hex()
}

func (f *FunctionDb)LogKey() string {
	return FnLogKey(f.Id)
}

func (f *FunctionDb)FnId() string {
	return f.Id.Hex()
}

type FnTargetDb struct {
	Id		bson.ObjectId		`bson:"id"`
}
