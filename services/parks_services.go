package services

import "github.com/shin888shin/parks/domain/parks"

func CreatePark(park parks.Park) (*parks.Park, error) {
	return &park, nil
}
