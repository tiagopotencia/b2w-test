package business

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

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

func GetPlanetsBusiness(db DatabaseInterface, swapi SwapiInterface) ([]Planet, error) {

	planetsList, err := db.GetPlanetFromDatabase()
	if err != nil {
		return planetsList, nil
	}

	for i, e := range planetsList {
		moviesCount, err := swapi.GetMoviesCount(e)

		if err != nil {
			return nil, err
		}

		planetsList[i].MoviesCount = moviesCount
	}

	return planetsList, nil
}

func AddPlanetBusiness(planet Planet, db DatabaseInterface) error {
	err := db.AddPlanetToDatabase(planet)
	return err
}

func GetPlanetsByName(name string, db DatabaseInterface, swapi SwapiInterface) ([]Planet, error) {
	planetsList, err := db.GetPlanetsByName(name)

	for i, e := range planetsList {
		moviesCount, err := swapi.GetMoviesCount(e)

		if err != nil {
			return nil, err
		}

		planetsList[i].MoviesCount = moviesCount
	}

	return planetsList, err
}

func GetPlanetByID(ID string, db DatabaseInterface, swapi SwapiInterface) (*Planet, error) {
	planet, err := db.GetPlanetByID(ID)

	if err != nil {
		return nil, err
	}

	if planet == nil {
		return nil, nil
	}

	moviesCount, err := swapi.GetMoviesCount(*planet)

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
		return 0, nil
	} else if swapiResponse.Results[0].Name != planetName {
		return 0, nil
	} else if len(swapiResponse.Results[0].Films) == 0 {
		return 0, nil
	} else {
		return len(swapiResponse.Results[0].Films), err
	}

}
