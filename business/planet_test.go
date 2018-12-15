package business

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tiagopotencia/i-am-back/mocks"
)

func TestAddPlanetBusiness(t *testing.T) {
	mockPlanet := Planet{
		Name:    "a",
		Climate: "a",
		Terrain: "a",
	}

	mockDB := &mocks.DatabaseInterface{}
	mockDB.On("AddPlanetToDatabase", mockPlanet).Return(nil)

	actual := AddPlanetBusiness(mockPlanet, mockDB)
	mockDB.AssertExpectations(t)
	assert.Equal(t, nil, actual, "msgAndArgs")
}

func TestGetAllPlanets(t *testing.T) {

	planetsList := []Planet{
		Planet{
			Name:    "a",
			Climate: "a",
			Terrain: "b",
		},
	}

	mockDB := &mocks.DatabaseInterface{}
	mockDB.On("GetPlanetFromDatabase").Return(planetsList)

	actual := GetPlanetsBusiness(mockDB)
	mockDB.AssertExpectations(t)
	assert.Equal(t, planetsList, actual, msgAndArgs)

}
