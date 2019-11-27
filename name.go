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
	 * It must be indexed
	 */
	Cookie		string		`bson:"cookie"`
}

const (
	NameLenMax		int    = 64
	SharedProject		string = "*"
)

var nameRe = regexp.MustCompile("^[\\p{L}\\d_]+(.[\\p{L}\\d_]+)*$")

func ValidName(n string) bool {
	if len(n) == 0 || len(n) >= NameLenMax {
		return false
	}

	return nameRe.MatchString(n)
}

func (n *Name)New(projectRef, name string) error {
	if !ValidName(name) {
		return errors.New("bad name value")
	}

	n.Make(projectRef, name)
	return nil
}

func (n *Name)Make(project, name string) *Name {
	fill(n, project, name)
	return n
}

func fill(into *Name, project, name string) {
	into.Name = name
	into.ProjectRef = project
	into.Cookie = util.Sha256(into.Str())
}

func (n *Name)Str() string {
	return n.ProjectRef + "::" + n.Name
}

func (n *Name)Q(q bson.M) bson.M {
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
