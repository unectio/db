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

	"github.com/unectio/util/mongo"
	"gopkg.in/mgo.v2/bson"
)

type LQ struct {
	q     bson.M
	sort  string
	limit int
}

func ListQ(q bson.M) *LQ {
	if q == nil {
		q = bson.M{}
	}

	return &LQ{q: q}
}

func (lq *LQ) Page(since bson.ObjectId, lim int) *LQ {
	lq.q["_id"] = bson.M{"$gt": since}
	lq.sort = "_id"
	lq.limit = lim
	return lq
}

func (lq *LQ) Tags(t []string) *LQ {
	/*
	 * Tags containining all the given values
	 * without regard to order or other values in it
	 */
	lq.q["tags"] = bson.M{"$all": t}
	return lq
}

func (lq *LQ) Q(ctx context.Context, loc *mongo.Location) mongo.Query {
	mq := MakeQuery(ctx, lq.q, loc)
	if lq.sort != "" {
		mq = mq.Sort(lq.sort)
		if lq.limit != 0 {
			mq = mq.Limit(lq.limit)
		}
	}
	return mq
}

func (lq *LQ) I(ctx context.Context, loc *mongo.Location) mongo.Iter {
	return lq.Q(ctx, loc).Iter()
}
