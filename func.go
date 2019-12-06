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
