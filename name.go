package db

import (
	"regexp"
	"errors"
	"github.com/unectio/util"
	"gopkg.in/mgo.v2/bson"
)

type Name struct {
	Name		string		`bson:"name"`
	ProjectRef	string		`bson:"project"`
	/*
	 * Search by name (a.k.a. lookup) uses this field.
	 * Field is indexed in setupIndexes
	 */
	Cookie		string		`bson:"cookie"`
}

const (
	NameLenMax		int    = 64
	SharedProject		string = "*"
)

var nameRe = regexp.MustCompile("^[\\p{L}\\d_]+(.[\\p{L}\\d_]+)*$")

func ValidName(n string) error {
	if len(n) == 0 {
		return errors.New("empty")
	}
	if len(n) >= NameLenMax {
		return errors.New("too long")
	}

	if !nameRe.MatchString(n) {
		return errors.New("invalid symbols")
	}

	return nil
}

func (n *Name)Fill(projectRef, name, extra string) error {
	err := ValidName(name)
	if err != nil {
		return errors.New("bad name value: " + err.Error())
	}

	n.Name = name
	n.ProjectRef = projectRef
	n.Cookie = util.Sha256(n.Str() + "::" + extra)

	return nil
}

func (n *Name)Str() string {
	return n.ProjectRef + "::" + n.Name
}

func (n *Name)Q() bson.M {
	q := Q()
	q["cookie"] = n.Cookie
	return q
}

type IdNameMap map[bson.ObjectId]string

func NewIdNameMap() IdNameMap {
	return make(map[bson.ObjectId]string)
}

func ProjQ(q bson.M, project string) bson.M {
	q["project"] = project
	return q
}
