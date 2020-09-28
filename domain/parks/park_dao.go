package parks

import (
	"fmt"
	"strings"

	"github.com/shin888shin/parks/datasources/mysql/parks"
	"github.com/shin888shin/parks/utils/date"
	"github.com/shin888shin/parks/utils/errors"
)

const (
	queryInsertPark = "INSERT INTO parks(name, description, location, created_at) VALUES(?, ?, ?, ?);"
	queryGetPark    = "SELECT id, name, description, location, created_at FROM parks WHERE id = ?;"

	errorNoRows = "no rows in result set"
)

func (park *Park) Get() *errors.RestErr {
	stmt, err := parks.Client.Prepare(queryGetPark)
	if err != nil {
		return errors.NewInternalServerErr(err.Error())
	}
	defer stmt.Close()
	getResult := stmt.QueryRow(park.ID)
	if err := getResult.Scan(&park.ID, &park.Name, &park.Description, &park.Location, &park.CreatedAt); err != nil {
		fmt.Println(err)
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundErr(fmt.Sprintf("park %d not found", park.ID))
		}
		return errors.NewInternalServerErr(fmt.Sprintf("error getting park %d: %s", park.ID, err.Error()))
	}

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
