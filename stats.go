package db

import (
	"time"
	"github.com/unectio/util/mongo"
	"gopkg.in/mgo.v2/bson"
)

type FuncStatsDb struct {
	Id		bson.ObjectId		`bson:"_id,omitempty"`
	FnId		bson.ObjectId		`bson:"fnid"`

	Calls		uint64			`bson:"calls"`
	RunTime		uint64			`bson:"runtime"`
	RealTime	uint64			`bson:"realtime"`
	LastCall	time.Time		`bson:"lastcall"`

	Arch		*StatsArchDb		`bson:"arch,omitempty"`
}

func (st *FuncStatsDb)ID() bson.ObjectId { return st.Id }
func (st *FuncStatsDb)Location() *mongo.Location { return LocStats }

type ProjectStatsDb struct {
	Id		bson.ObjectId		`bson:"_id,omitempty"`
	Project		string			`bson:"project"`

	Calls		uint64			`bson:"calls"`
	RunTime		uint64			`bson:"runtime"`
	RealTime	uint64			`bson:"realtime"`

	Arch		*StatsArchDb		`bson:"arch,omitempty"`
}

func (st *ProjectStatsDb)ID() bson.ObjectId { return st.Id }
func (st *ProjectStatsDb)Location() *mongo.Location { return LocPStats }

type StatsArchDb struct {
	Ts		time.Time		`bson:"ts"`
	Label		string			`bson:"label"`
}
