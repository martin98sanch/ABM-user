package user

import (
	"log"

	"github.com/martin98sanch/ABM-user/src/api/sql"
)

type (
	GetListFunc func() (Users, error)
	GetByIDFunc func(id ID) (DTO, error)
)

//Get a list of users
func MakeGetListFunc(sqlQuery sql.QueryFunc) (GetListFunc, error) {
	if sqlQuery == nil {
		return nil, ErrCantMakeTheInjection
	}
	return func() (Users, error) {
		query := sql.SelectAllUsersStatement
		users := Users{}
		rows, err := sql.Query(query)
		if err != nil {
			return users, ErrCantGetUserList
		}

		for rows.Next() {
			user := DTO{}
			rows.Scan(&user.ID, &user.Username, &user.Password, &user.Name, &user.Age, &user.Deleted)
			users = append(users, user)
			log.Printf("User: %+v", user)
		}

		return users, nil
	}, nil
}

//Get a user by a given ID
func MakeGetByIDFunc(sqlQuery sql.QueryFunc) (GetByIDFunc, error) {
	if sqlQuery == nil {
		return nil, ErrCantMakeTheInjection
	}
	return func(id ID) (DTO, error) {
		user := DTO{}
		query := sql.SelectUserByIDStatement
		rows, err := sql.Query(query, id)
		if err != nil {
			return user, ErrCantGetUserByID
		}

		for rows.Next() {
			rows.Scan(&user.ID, &user.Username, &user.Password, &user.Name, &user.Age, &user.Deleted)
		}

		if err := user.Validate(); err != nil {
			return user, ErrNotFound
		}

		return user, nil
	}, nil
}
