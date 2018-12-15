package business

import (
	"log"

	"gopkg.in/mgo.v2/bson"

	mgo "gopkg.in/mgo.v2"
)

//var db *mgo.Database
var dbName string
var COLLECTION string

type DatabaseInterface interface {
	AddPlanetToDatabase(interface{}) error
	ConnectToDB(string, string) error
	GetPlanetFromDatabase() ([]Planet, error)
	GetPlanetsByName(name string) ([]Planet, error)
	GetPlanetByID(ID string) (*Planet, error)
	DeletePlanet(ID string) (bool, error)
}

type Database struct {
	Session *mgo.Session
	DbName  string
}

func (d *Database) ConnectToDB(dbUri, dbName string) error {

	// dbName = os.Getenv("dbName")
	if dbName == "" {
		dbName = "b2w-test"
	}
	session, err := mgo.Dial(dbUri + dbName)

	if err != nil {
		log.Fatal(err)
	}

	mongo := Database{
		Session: session,
		DbName:  dbName,
	}

	d.DbName = mongo.DbName
	d.Session = mongo.Session

	COLLECTION = "planets"
	//db = session.DB(dbName)

	return nil
}

func (d Database) AddPlanetToDatabase(planet interface{}) error {
	session := d.Session.Copy()
	defer session.Close()

	c := session.DB(dbName).C(COLLECTION)
	err := c.Insert(planet)

	if err != nil {
		return err
	}

	return nil
}

func (d *Database) GetPlanetFromDatabase() ([]Planet, error) {
	session := d.Session.Copy()
	defer session.Close()

	var result []Planet

	c := session.DB(d.DbName).C(COLLECTION)
	err := c.Find(bson.M{}).All(&result)

	return result, err
}

func (d *Database) GetPlanetsByName(name string) ([]Planet, error) {
	session := d.Session.Copy()
	defer session.Close()

	var result []Planet

	c := session.DB(d.DbName).C(COLLECTION)
	err := c.Find(bson.M{"name": name}).All(&result)

	if err == mgo.ErrNotFound {
		return result, nil
	}

	return result, err
}

func (d *Database) GetPlanetByID(ID string) (*Planet, error) {
	session := d.Session.Copy()
	defer session.Close()

	var result Planet

	if bson.IsObjectIdHex(ID) == false {
		return nil, nil
	}

	c := session.DB(d.DbName).C(COLLECTION)
	err := c.FindId(bson.ObjectIdHex(ID)).One(&result)

	if err == mgo.ErrNotFound {
		return nil, nil
	}

	return &result, err
}

func (d *Database) DeletePlanet(ID string) (bool, error) {
	session := d.Session.Copy()
	defer session.Close()

	c := session.DB(d.DbName).C(COLLECTION)
	err := c.RemoveId(bson.ObjectIdHex(ID))

	if err == mgo.ErrNotFound {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil

}
