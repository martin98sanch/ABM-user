package user

import (
	"github.com/martin98sanch/ABM-user/src/api/sql"
)

type (
	DeleteByIDFunc func(id ID) error
)

//Delete user. This is not a hard delete
func MakeDeleteByIDFunc(sqlExec sql.ExecFunc) (DeleteByIDFunc, error) {
	if sqlExec == nil {
		return nil, ErrCantMakeTheInjection
	}
	return func(id ID) error {
		query := sql.DeleteUserByIDStatement
		if _, err := sql.Exec(query, true, id); err != nil {
			return ErrCantDeleteUserByID
		}
		return nil
	}, nil
}
