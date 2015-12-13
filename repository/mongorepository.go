package repository

import (
	"fmt"

	"gopkg.in/mgo.v2"
//	"gopkg.in/mgo.v2/bson"
	"reflect"
	"strings"
)

type MongoRepository struct {
	database string
	url string
	session *mgo.Session
}

func (this *MongoRepository) GetSession() *mgo.Session {
	if this.session == nil {
		session, err := mgo.Dial(this.url)
		if err != nil {
			panic(err)
		}

		this.session = session

		// Optional. Switch the session to a monotonic behavior.
		this.session.SetMode(mgo.Monotonic, true)
	}
	return this.session.Copy()
}

func (this *MongoRepository) Get(i *interface{}) error {
	session := this.GetSession()
	defer session.Close()

	domain := reflect.TypeOf(i).String()
	tableName := strings.TrimLeft(domain, "*model.")

	err := session.DB(this.database).C(tableName).Find(i).One(i)
	return err

}

func (this *MongoRepository) Save(i interface{}) error {
	session := this.GetSession()
	defer session.Close()

	domain := reflect.TypeOf(i).String()
	tableName := strings.TrimLeft(domain, "*.model.")

	err := session.DB(this.database).C(tableName).Insert(i)
	return err
}

func (this *MongoRepository) Find(i interface{}) interface{} {
	fmt.Println("Finding an entity")
	return true
}
