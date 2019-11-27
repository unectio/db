package db

import (
	"time"
	"github.com/unectio/util/mongo"
	"gopkg.in/mgo.v2/bson"
)

type RepoDb struct {
	DbCommon				`bson:",inline"`

	State		string			`bson:"state"`
	Synced		time.Time		`bson:"synced"`

	Sync		string			`bson:"sync"`
	SyncDelayM	int			`bson:"sync_delay_m,omitempty"`

	Type		string			`bson:"type"`
	URL		string			`bson:"url"`
}

func (rp *RepoDb)UpdateSyncedRq(q bson.M) bson.M {
	if q == nil {
		q = bson.M{}
	}

	/* .Synced */
	q["synced"] = time.Now()

	return q
}

func (rp *RepoDb)BySyncType(typ string) bson.M {
	/* .Sync */
	return bson.M{"sync": typ}
}

func (rp *RepoDb)ID() bson.ObjectId { return rp.Id }
func (rp *RepoDb)Location() *mongo.Location { return LocRepo }

const (
	SyncTimer	= "timer"
)

func (rp *RepoDb)SyncDelay() time.Duration {
	return time.Duration(rp.SyncDelayM) * time.Minute
}
