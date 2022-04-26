package server

import (
	"github.com/martin98sanch/ABM-user/src/api/sql"
	"github.com/martin98sanch/ABM-user/src/api/user"
)

func resolveUserHandler() *user.Handler {
	return &user.Handler{
		UserCreator: resolveUserCreatorFunc(),
		GetUserList: resolveUserListGetterFunc(),
		GetByID:     resolveUserByIDFunc(),
	}
}

func resolveUserListGetterFunc() user.GetListFunc {
	f, err := user.MakeGetListFunc(sql.Query)
	if err != nil {
		panic(err)
	}

	return f
}

func resolveUserByIDFunc() user.GetByIDFunc {
	f, err := user.MakeGetByIDFunc(sql.Query)
	if err != nil {
		panic(err)
	}

	return f
}

func resolveUserCreatorFunc() user.CreatorFunc {
	f, err := user.MakeCreatorFunc(sql.Exec, sql.Query)
	if err != nil {
		panic(err)
	}

	return f
}
