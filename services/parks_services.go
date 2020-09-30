package services

import (
	"github.com/shin888shin/parks/domain/parks"
	"github.com/shin888shin/parks/utils/date"
	"github.com/shin888shin/parks/utils/errors"
)

var (
	ParksService parksServiceInterface = &parksService{}
)

type parksService struct {
}

type parksServiceInterface interface {
	GetPark(int64) (*parks.Park, *errors.RestErr)
	CreatePark(parks.Park) (*parks.Park, *errors.RestErr)
	UpdatePark(parks.Park) (*parks.Park, *errors.RestErr)
	GetAllParks() ([]parks.Park, *errors.RestErr)
}

func (ps *parksService) GetPark(parkID int64) (*parks.Park, *errors.RestErr) {
	result := &parks.Park{ID: parkID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func (ps *parksService) CreatePark(park parks.Park) (*parks.Park, *errors.RestErr) {
	if err := park.Validate(); err != nil {
		return nil, err
	}
	park.CreatedAt = date.GetNowDatetime()
	if err := park.Save(); err != nil {
		return nil, err
	}

	// return either a result or an error -- not both
	return &park, nil
}

func (ps *parksService) UpdatePark(park parks.Park) (*parks.Park, *errors.RestErr) {
	current, err := ps.GetPark(park.ID)
	if err != nil {
		return nil, err
	}

	current.Name = park.Name
	current.Description = park.Description
	current.Location = park.Location

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func (ps *parksService) GetAllParks() ([]parks.Park, *errors.RestErr) {
	dao := &parks.Park{}
	return dao.GetAllParks()
}
