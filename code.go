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

type FuncCodeDb struct {
	DbCommon				`bson:",inline"`

	State		string			`bson:"state"`

	FnId		bson.ObjectId		`bson:"fnid"`
	Lang		string			`bson:"lang"`
	Repo		*FuncRepoLink		`bson:"repo_link,omitempty"`

	Gen		int			`bson:"gen"`

	/*
	 * Key for lookup (deployment -> *code).
	 * XXX The field must be indexed.
	 */
	DepKey		string			`bson:"dep_key"`
}

func (c *FuncCodeDb)ByFn(id bson.ObjectId) bson.M {
	/* .FnId */
	return bson.M{"fnid": id}
}

func (c *FuncCodeDb)ByRepo(id bson.ObjectId) bson.M {
	/* .Repo.Id */
	return bson.M{"repo_link.id": id}
}

func (c *FuncCodeDb)ByDepKey(dk string) bson.M {
	/* .DepKey */
	return bson.M{"dep_key": dk}
}

func (c *FuncCodeDb)ID() bson.ObjectId { return c.Id }
func (c *FuncCodeDb)Location() *mongo.Location { return LocCode }

func (c *FuncCodeDb)Version() string { return c.Id.Hex() }

type FuncRepoLink struct {
	Id		bson.ObjectId		`bson:"id"`
	Path		string			`bson:"path"`
}


func (c *FuncCodeDb)UpdateQ() bson.M {
	/* .Repo & .Gen */
	return bson.M{"repo_link": c.Repo, "gen": c.Gen}
}
