package parks

import (
	"strings"
	"time"

	"github.com/shin888shin/parks/utils/errors"
)

type Park struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	CreatedAt   time.Time `json:"created_at"`
}

func (park *Park) Validate() *errors.RestErr {
	park.Description = strings.TrimSpace(park.Description)
	park.Location = strings.TrimSpace(park.Location)

	park.Name = strings.TrimSpace(strings.ToLower(park.Name))
	if park.Name == "" {
		return errors.NewBadRequestErr("invalid name")
	}
	return nil
}
