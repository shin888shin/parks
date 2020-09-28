package parks

import (
	"fmt"

	"github.com/shin888shin/parks/datasources/mysql/parks"
	"github.com/shin888shin/parks/utils/date"
	"github.com/shin888shin/parks/utils/errors"
)

var (
	parksDB = make(map[int64]*Park)
)

const (
	queryInsertPark = "INSERT INTO parks(name, description, location, created_at) VALUES(?, ?, ?, ?);"
)

func (park *Park) Get() *errors.RestErr {
	if err := parks.Client.Ping(); err != nil {
		panic(err)
	}
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
	stmt, err := parks.Client.Prepare(queryInsertPark)
	if err != nil {
		return errors.NewInternalServerErr(err.Error())
	}
	defer stmt.Close()

	park.CreatedAt = date.GetNowDatetime()

	insertResult, err := stmt.Exec(park.Name, park.Description, park.Location, park.CreatedAt)
	if err != nil {
		// TODO: maybe turn on name uniqueness
		// if strings.Contains(err.Error(), "name_UNIQUE") {
		// 	return errors.NewBadRequestErr(fmt.Sprintf("name %s must be unique ", park.Name))
		// }
		return errors.NewInternalServerErr(fmt.Sprintf("error on save: %s", err.Error()))
	}

	parkID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerErr(fmt.Sprintf("error on save: %s", err.Error()))
	}

	park.ID = parkID
	return nil
}
