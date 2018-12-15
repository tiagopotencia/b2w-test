package business

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type SwapiInterface interface {
	GetMoviesCount(planet Planet) (int, error)
}

type Swapi struct{}

func (s *Swapi) GetMoviesCount(planet Planet) (int, error) {
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
