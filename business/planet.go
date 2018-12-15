package business

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

var PlanetHasNoMoviesError error

func init() {
	PlanetHasNoMoviesError = errors.New("This planet has no movies. Try to add a planet that has movie.")
}

type Planet struct {
	Name        string `bson: "name", json: "name"`
	Climate     string `bson: "climate", json: "climate"`
	Terrain     string `bson: "terrain", json: "terrain"`
	MoviesCount int    `json: "moviesCount,omitempty"`
}

type SwapiPlanetResponse struct {
	Count    int         `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		Name           string    `json:"name"`
		RotationPeriod string    `json:"rotation_period"`
		OrbitalPeriod  string    `json:"orbital_period"`
		Diameter       string    `json:"diameter"`
		Climate        string    `json:"climate"`
		Gravity        string    `json:"gravity"`
		Terrain        string    `json:"terrain"`
		SurfaceWater   string    `json:"surface_water"`
		Population     string    `json:"population"`
		Residents      []string  `json:"residents"`
		Films          []string  `json:"films"`
		Created        time.Time `json:"created"`
		Edited         time.Time `json:"edited"`
		URL            string    `json:"url"`
	} `json:"results"`
}

func GetPlanetsBusiness(db DatabaseInterface) ([]Planet, error) {

	planetsList, err := db.GetPlanetFromDatabase()
	if err != nil {
		return planetsList, nil
	}

	for i, e := range planetsList {
		moviesCount, err := getMoviesCount(e)

		if err != nil {
			return nil, err
		}

		planetsList[i].MoviesCount = moviesCount
	}

	return planetsList, nil
}

func AddPlanetBusiness(planet Planet, db DatabaseInterface) error {

	_, err := getMoviesCount(planet)

	if err != nil {
		return err
	}

	err = db.AddPlanetToDatabase(planet)
	return err
}

func GetPlanetsByName(name string, db DatabaseInterface) ([]Planet, error) {
	planetsList, err := db.GetPlanetsByName(name)

	for i, e := range planetsList {
		moviesCount, err := getMoviesCount(e)

		if err != nil {
			return nil, err
		}

		planetsList[i].MoviesCount = moviesCount
	}

	return planetsList, err
}

func GetPlanetByID(ID string, db DatabaseInterface) (*Planet, error) {
	planet, err := db.GetPlanetByID(ID)

	if err != nil {
		return nil, err
	}

	if planet == nil {
		return nil, nil
	}

	moviesCount, err := getMoviesCount(*planet)

	planet.MoviesCount = moviesCount

	return planet, err
}

func DeletePlanet(ID string, db DatabaseInterface) (bool, error) {
	planetDeleted, err := db.DeletePlanet(ID)

	return planetDeleted, err
}

func getMoviesCount(planet Planet) (int, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	var swapiResponse = SwapiPlanetResponse{}

	var planetName = planet.Name

	response, err := client.Get("https://swapi.co/api/planets?search=" + url.QueryEscape(planetName))
	if err != nil {
		return 0, err
	}

	buf, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}

	json.Unmarshal(buf, &swapiResponse)
	if err != nil {
		return 0, err
	}

	if len(swapiResponse.Results) == 0 {
		return 0, PlanetHasNoMoviesError
	} else if len(swapiResponse.Results[0].Films) == 0 {
		return 0, PlanetHasNoMoviesError
	} else {
		return len(swapiResponse.Results[0].Films), err
	}

}
