package user

import "github.com/martin98sanch/ABM-user/src/api/sql"

type (
	UpdateByIDFunc func(user *DTO) error
)

//Update user info
func MakeUpdateByIDFunc(sqlExec sql.ExecFunc, sqlQuery sql.QueryFunc) (UpdateByIDFunc, error) {
	if sqlExec == nil {
		return nil, ErrCantMakeTheInjection
	}
	return func(user *DTO) error {
		if err := validateData(sqlQuery, user); err != nil {
			return err
		}

		query := sql.UpdateUserByIDStatement
		if _, err := sql.Exec(query, user.Username, user.Password, user.Name, user.Age, user.ID); err != nil {
			return ErrCantUpdateUserByID
		}
		return nil
	}, nil
}
