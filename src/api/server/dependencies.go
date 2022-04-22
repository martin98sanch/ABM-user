package server

import (
	"github.com/martin98sanch/ABM-user/src/api/sql"
	"github.com/martin98sanch/ABM-user/src/api/user"
)

func resolveUserHandler() *user.Handler {
	return &user.Handler{
		UserCreator: resolveUserCreatorFunc(),
	}
}

func resolveUserCreatorFunc() user.CreatorFunc {
	f, err := user.MakeCreatorFunc(sql.Exec)
	if err != nil {
		panic(err)
	}

	return f
}
