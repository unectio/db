package db

import (
	"context"
	"net/http"
	"github.com/unectio/util"
	"github.com/unectio/util/mongo"
	"github.com/unectio/util/restmux"
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
	if mongo.IsNotFound(err) {
		return &restmux.GenError{http.StatusNotFound, "No such object"}
	} else {
		sc.Log("DB").Errorf("DB connection (or query) error at %s: %s", util.Caller(), err.Error())
		return &restmux.GenError{http.StatusInternalServerError, "Error looking up object"}
	}
}

func col(ctx context.Context, loc *mongo.Location) mongo.Collection {
	return sc.GetDb(ctx).Collection(loc)
}

func InitRealMongo(ctx context.Context, url string) error {
	s, err := mongo.Connect(url)
	if err == nil {
		defer s.Close()
		sc.SetDB(ctx, s)
	}
	return err
}

func Q() bson.M { return bson.M{} }
