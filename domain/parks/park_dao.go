package parks

import (
	"github.com/shin888shin/parks/datasources/mysql/parks"
	"github.com/shin888shin/parks/utils/date"
	"github.com/shin888shin/parks/utils/errors"
)

const (
	queryInsertPark = "INSERT INTO parks(name, description, location, created_at) VALUES(?, ?, ?, ?);"
	queryGetPark    = "SELECT id, name, description, location, created_at FROM parks WHERE id = ?;"
)

func (park *Park) Get() *errors.RestErr {
	stmt, err := parks.Client.Prepare(queryGetPark)
	if err != nil {
		return errors.NewInternalServerErr(err.Error())
	}
	defer stmt.Close()
	getResult := stmt.QueryRow(park.ID)

	if err := getResult.Scan(&park.ID, &park.Name, &park.Description, &park.Location, &park.CreatedAt); err != nil {
		return errors.ParseError(err)
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
		return errors.ParseError(err)
	}

	parkID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.ParseError(err)
	}

	park.ID = parkID
	return nil
}
