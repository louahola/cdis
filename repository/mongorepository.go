package repository

import (
	"fmt"

	"gopkg.in/mgo.v2"
//	"gopkg.in/mgo.v2/bson"
	"reflect"
	"strings"
)

type MongoRepository struct {
	_session *mgo.Session
}

func (this *MongoRepository) session() *mgo.Session {
	if this._session == nil {
		session, err := mgo.Dial("localhost")
		if err != nil {
			panic(err)
		}

		this._session = session

		// Optional. Switch the session to a monotonic behavior.
		this._session.SetMode(mgo.Monotonic, true)
	}
	return this._session.Copy()
}

func (this *MongoRepository) Save(i interface{}) error {
	session := this.session()
	defer session.Close()

	domain := reflect.TypeOf(i).String()
	tableName := strings.TrimLeft(domain, "*.model.")

	err := session.DB("cdis").C(tableName).Insert(i)
	return err
}

func (this *MongoRepository) Find(i interface{}) interface{} {
	fmt.Println("Finding an entity")
	return true
}
