package business

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Results struct {
	Films []string
}

func TestAddPlanetBusiness(t *testing.T) {
	mockPlanet := Planet{
		Name:    "a",
		Climate: "a",
		Terrain: "a",
	}

	mockDB := &DatabaseInterfaceMock{}
	mockDB.On("AddPlanetToDatabase", mockPlanet).Return(nil)

	actual := AddPlanetBusiness(mockPlanet, mockDB)
	mockDB.AssertExpectations(t)
	assert.Equal(t, nil, actual, "Must Add Planet")
}

func TestGetAllPlanets(t *testing.T) {

	planetsList := []Planet{
		Planet{
			Name:        "a",
			Climate:     "a",
			Terrain:     "b",
			MoviesCount: 0,
		},
	}

	mockPlanet := Planet{
		Name:    "a",
		Climate: "a",
		Terrain: "b",
	}

	mockDB := &DatabaseInterfaceMock{}
	mockDB.On("GetPlanetFromDatabase").Return(planetsList, nil)

	mockSwapi := &SwapiInterfaceMock{}
	mockSwapi.On("GetMoviesCount", mockPlanet).Return(1, nil)

	actual, _ := GetPlanetsBusiness(mockDB, mockSwapi)
	mockDB.AssertExpectations(t)
	assert.Equal(t, planetsList, actual, "Must return planets")

}

func TestGetPlanetsByName(t *testing.T) {
	planetsList := []Planet{
		Planet{
			Name:        "a",
			Climate:     "a",
			Terrain:     "b",
			MoviesCount: 0,
		},
	}

	mockPlanet := Planet{
		Name:    "a",
		Climate: "a",
		Terrain: "b",
	}

	mockDB := &DatabaseInterfaceMock{}
	mockDB.On("GetPlanetsByName", "a").Return(planetsList, nil)

	mockSwapi := &SwapiInterfaceMock{}
	mockSwapi.On("GetMoviesCount", mockPlanet).Return(1, nil)

	actual, _ := GetPlanetsByName("a", mockDB, mockSwapi)
	mockDB.AssertExpectations(t)
	assert.Equal(t, planetsList, actual, "Must return planets")
}

func TestGetPlanetsByID(t *testing.T) {
	planet := Planet{
		Name:        "a",
		Climate:     "a",
		Terrain:     "b",
		MoviesCount: 0,
	}

	mockPlanet := Planet{
		Name:    "a",
		Climate: "a",
		Terrain: "b",
	}

	mockDB := &DatabaseInterfaceMock{}
	mockDB.On("GetPlanetByID", "a").Return(&planet, nil)

	mockSwapi := &SwapiInterfaceMock{}
	mockSwapi.On("GetMoviesCount", mockPlanet).Return(1, nil)

	actual, _ := GetPlanetByID("a", mockDB, mockSwapi)
	mockDB.AssertExpectations(t)
	assert.Equal(t, planet, *actual, "Must return planets")
}

func TestDetelePlanet(t *testing.T) {

	mockDB := &DatabaseInterfaceMock{}
	mockDB.On("DeletePlanet", "a").Return(true, nil)

	actual, _ := DeletePlanet("a", mockDB)
	mockDB.AssertExpectations(t)
	assert.Equal(t, true, actual, "Must delete planet")
}
