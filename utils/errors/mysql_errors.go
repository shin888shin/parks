package errors

import (
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return NewNotFoundErr(fmt.Sprintf("no record found"))
		}
		return NewInternalServerErr("error parsing error")
	}

	switch sqlErr.Number {
	case 1062:
		// error: name_UNIQUE
		return NewBadRequestErr(fmt.Sprintf("invalid data"))
	}
	return NewInternalServerErr("error processing request")
}
