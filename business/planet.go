package business

import (
	"log"
)

type Planet struct {
	Name    string `bson: "name", json: "name"`
	Climate string `bson: "climate", json: "climate"`
	Terrain string `bson: "terrain", json: "terrain"`
}

func GetPlanetsBusiness(db DatabaseInterface) ([]Planet, error) {

	planetList, err := db.GetPlanetFromDatabase()
	if err != nil {
		log.Println(err)
	}
	return planetList, nil
}

func AddPlanetBusiness(planet Planet, db DatabaseInterface) error {
	err := db.AddPlanetToDatabase(planet)

	if err != nil {
		log.Print(err)
	}

	return err
}

func GetPlanetsByName(name string, db DatabaseInterface) ([]Planet, error) {
	planetsList, err := db.GetPlanetsByName(name)
	return planetsList, err
}

func GetPlanetByID(ID string, db DatabaseInterface) (*Planet, error) {
	planet, err := db.GetPlanetByID(ID)

	return planet, err
}

func DeletePlanet(ID string, db DatabaseInterface) (bool, error) {
	planetDeleted, err := db.DeletePlanet(ID)

	return planetDeleted, err
}
