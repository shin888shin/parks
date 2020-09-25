package services

import (
	"github.com/shin888shin/parks/domain/parks"
	"github.com/shin888shin/parks/utils/errors"
)

func CreatePark(park parks.Park) (*parks.Park, *errors.RestErr) {
	if err := park.Validate(); err != nil {
		return nil, err
	}
	if err := park.Save(); err != nil {
		return nil, err
	}

	// return either a result or an error -- not both
	return &park, nil
}
