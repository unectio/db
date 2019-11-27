package db

import (
	"context"
	"github.com/unectio/util/mongo"
	"gopkg.in/mgo.v2/bson"
)

type LQ struct {
	q	bson.M
	sort	string
	limit	int
}

func ListQ(q bson.M) *LQ {
	if q == nil {
		q = bson.M{}
	}

	return &LQ{q: q}
}

func (lq *LQ)Page(since bson.ObjectId, lim int) *LQ {
	lq.q["_id"] = bson.M{"$gt": since}
	lq.sort = "_id"
	lq.limit = lim
	return lq
}

func (lq *LQ)Tags(t []string) *LQ {
	/*
	 * Tags containining all the given values
	 * without regard to order or other values in it
	 */
	lq.q["tags"] = bson.M{"$all": t}
	return lq
}

func (lq *LQ)Q(ctx context.Context, loc *mongo.Location) mongo.Query {
	mq := MakeQuery(ctx, lq.q, loc)
	if lq.sort != "" {
		mq = mq.Sort(lq.sort)
		if lq.limit != 0 {
			mq = mq.Limit(lq.limit)
		}
	}
	return mq
}

func (lq *LQ)I(ctx context.Context, loc *mongo.Location) mongo.Iter {
	return lq.Q(ctx, loc).Iter()
}
