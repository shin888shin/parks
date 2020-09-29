package parks

import (
	"github.com/shin888shin/parks/datasources/mysql/parks"
	"github.com/shin888shin/parks/utils/errors"
)

const (
	queryInsertPark  = "INSERT INTO parks(name, description, location, created_at) VALUES(?, ?, ?, ?);"
	queryGetPark     = "SELECT id, name, description, location, created_at FROM parks WHERE id = ?;"
	queryGetAllParks = "SELECT id, name, description, location, created_at FROM parks;"
	queryUpdatePark  = "UPDATE parks SET name = ?, description = ?, location = ? WHERE id = ?;"
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

func (park *Park) Update() *errors.RestErr {
	stmt, err := parks.Client.Prepare(queryUpdatePark)
	if err != nil {
		return errors.NewInternalServerErr(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(park.Name, park.Description, park.Location, park.ID)

	if err != nil {
		return errors.ParseError(err)
	}

	return nil
}

func (park *Park) GetAllParks() ([]Park, *errors.RestErr) {
	stmt, err := parks.Client.Prepare(queryGetAllParks)
	if err != nil {
		return nil, errors.NewInternalServerErr(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, errors.NewInternalServerErr(err.Error())
	}
	defer rows.Close()

	results := make([]Park, 0)
	for rows.Next() {
		var park Park
		if err := rows.Scan(&park.ID, &park.Name, &park.Description, &park.Location, &park.CreatedAt); err != nil {
			return nil, errors.ParseError(err)
		}
		results = append(results, park)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundErr("no records")
	}
	return results, nil
}
