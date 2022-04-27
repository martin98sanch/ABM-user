package user

import (
	"github.com/martin98sanch/ABM-user/src/api/sql"
)

type (
	CreatorFunc func(user *DTO) error
)

//Create a user
func MakeCreatorFunc(sqlExec sql.ExecFunc, sqlQuery sql.QueryFunc) (CreatorFunc, error) {
	if sqlExec == nil {
		return nil, ErrCantMakeTheInjection
	}
	return func(user *DTO) error {
		if err := validateData(sqlQuery, user); err != nil {
			return err
		}

		query := sql.InsertUserStatement
		if _, err := sqlExec(query, user.Username, user.Password, user.Name, user.Age); err != nil {
			return ErrCantCreateUser
		}
		return nil
	}, nil

}

//ValidateData checks if username is already in use
func validateData(sqlQuery sql.QueryFunc, user *DTO) error {
	query := sql.FindUserByUsernameStatement
	rows, err := sql.Query(query, user.Username)
	if err != nil {
		return ErrCantCheckIfUserAlreadyExist
	}

	for rows.Next() {
		return ErrUserAlreadyExist
	}

	return nil
}
