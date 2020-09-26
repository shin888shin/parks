package parks

import (
	"fmt"
	"strings"

	"github.com/shin888shin/parks/utils/date"
	"github.com/shin888shin/parks/utils/errors"
)

var (
	parksDB = make(map[int64]*Park)
)

func (park *Park) Get() *errors.RestErr {
	result := parksDB[park.ID]
	if result == nil {
		return errors.NewNotFoundErr(fmt.Sprintf("park %d not found", park.ID))
	}
	park.ID = result.ID
	park.Name = result.Name
	park.Description = result.Description
	park.Location = result.Location
	park.CreatedAt = result.CreatedAt
	return nil
}

func (park *Park) Save() *errors.RestErr {
	current := parksDB[park.ID]
	if current != nil {
		if strings.TrimSpace(strings.ToLower(current.Name)) == strings.TrimSpace(strings.ToLower(park.Name)) {
			return errors.NewBadRequestErr(fmt.Sprintf("park with name %s already exists", park.Name))
		}
		return errors.NewBadRequestErr(fmt.Sprintf("park %d already exists", park.ID))
	}

	park.CreatedAt = date.GetNowString()
	parksDB[park.ID] = park
	return nil
}
