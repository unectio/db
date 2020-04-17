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
	"net/http"
	"github.com/unectio/util"
	"github.com/unectio/util/mongo"
	"github.com/unectio/util/restmux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	sc "github.com/unectio/util/context"
)

func Find(ctx context.Context, q bson.M, out DbObject) error {
	return MakeQuery(ctx, q, out.Location()).One(out)
}

func Load(ctx context.Context, id bson.ObjectId, obj DbObject) error {
	return Find(ctx, bson.M{"_id": id}, obj)
}

func Iterator(ctx context.Context, q bson.M, loc *mongo.Location) mongo.Iter {
	return MakeQuery(ctx, q, loc).Iter()
}

func MakeQuery(ctx context.Context, q bson.M, loc *mongo.Location) mongo.Query {
	return col(ctx, loc).Find(q)
}

func Upd(ctx context.Context, o DbObject, q bson.M, u bson.M) error {
	var upd interface{}

	if u == nil {
		upd = o /* full rewrite */
	} else {
		upd = bson.M{"$set": u}
	}

	if q == nil {
		q = bson.M{"_id": o.ID()}
	}

	return col(ctx, o.Location()).Update(q, upd)
}

func Update(ctx context.Context, loc *mongo.Location, q bson.M, u bson.M) error {
	return col(ctx, loc).Update(q, u)
}

func Del(ctx context.Context, o DbObject) error {
	return col(ctx, o.Location()).RemoveId(o.ID())
}

func DelAll(ctx context.Context, loc *mongo.Location, q bson.M) error {
	return col(ctx, loc).RemoveAll(q)
}

func Add(ctx context.Context, o DbObject) error {
	return col(ctx, o.Location()).Insert(o)
}

func RestError(err error) restmux.Error {
	if err == nil {
		sc.Log("DB").Errorf("DB error tries to report nil")
		return &restmux.GenError{http.StatusInternalServerError, "Error reporting error %) sorry"}
	} else if mongo.IsNotFound(err) {
		return &restmux.GenError{http.StatusNotFound, "No such object"}
	} else if mongo.IsDup(err) {
		return &restmux.GenError{http.StatusConflict, "Name already exists"}
	} else {
		sc.Log("DB").Errorf("DB connection (or query) error at %s: %s", util.Caller(), err.Error())
		return &restmux.GenError{http.StatusInternalServerError, "Error querying database"}
	}
}

func col(ctx context.Context, loc *mongo.Location) mongo.Collection {
	return sc.GetDb(ctx).Collection(loc)
}

func InitRealMongo(ctx context.Context, url string) error {
	s, err := mongo.Connect(url)
	if err == nil {
		err = setupIndexes(s)
		defer s.Close()
		sc.SetDB(ctx, s)
	}
	return err
}

func setupIndexes(s mongo.Session) error {
	var err error

	err = setupCookieIndex(s, LocFunc)	;if err != nil { return err }
	err = setupCookieIndex(s, LocTrigger)	;if err != nil { return err }
	err = setupCookieIndex(s, LocCode)	;if err != nil { return err }
	err = setupCookieIndex(s, LocRepo)	;if err != nil { return err }
	err = setupCookieIndex(s, LocRouter)	;if err != nil { return err }
	err = setupCookieIndex(s, LocAuthCtx)	;if err != nil { return err }
	err = setupCookieIndex(s, LocApps)	;if err != nil { return err }
	err = setupCookieIndex(s, LocSecret)	;if err != nil { return err }
	err = setupCookieIndex(s, LocWebsock)	;if err != nil { return err }
	err = setupCookieIndex(s, LocMongo)	;if err != nil { return err }

	err = setupFieldIndex(s, LocTrigger, "key", false); if err != nil { return err }

	return nil
}

func setupCookieIndex(s mongo.Session, loc *mongo.Location) error {
	return setupFieldIndex(s, loc, "cookie", true)
}

func setupFieldIndex(s mongo.Session, loc *mongo.Location, field string, unique bool) error {
	index := mgo.Index {
		Unique:		unique,
		Background:	true,
		Sparse:		true,
		Key:		[]string{ field },
	}

	return s.Collection(loc).EnsureIndex(&index)
}

func Q() bson.M { return bson.M{} }
