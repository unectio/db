package db

import (
	"time"
	"github.com/unectio/util/mongo"
	"gopkg.in/mgo.v2/bson"
)

type LogEntryDb struct {
	Key	string			`bson:"key"`
	Event	string			`bson:"event"`
	Time	time.Time		`bson:"time"`
	Text	string			`bson:"text"`
}

func (le *LogEntryDb)ID() bson.ObjectId { return bson.NewObjectId() }
func (le *LogEntryDb)Location() *mongo.Location { return LocLogs }

func FnLogs(f *FunctionDb) bson.M {
	/* .Key */
	return bson.M{"key": f.LogKey()}
}

func LogsSince(q bson.M, since time.Time) {
	/* .Time */
	q["time"] = bson.M{"$gt": since}
}
