package parks

import (
	"strings"

	"github.com/shin888shin/parks/utils/errors"
)

type Park struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Location    string `json:"location"`
	CreatedAt   string `json:"created_at"`
}

func (park *Park) Validate() *errors.RestErr {
	park.Name = strings.TrimSpace(strings.ToLower(park.Name))
	if park.Name == "" {
		return errors.NewBadRequestErr("invalid name")
	}
	return nil
}
