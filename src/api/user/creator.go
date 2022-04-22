package user

import "github.com/martin98sanch/ABM-user/src/api/sql"

type (
	CreatorFunc    func(user *DTO) error
	GetListFunc    func() (Users, error)
	GetByIDFunc    func(id int) (DTO, error)
	UpdateByIDFunc func(id DTO) error
)

//Create a user
func MakeCreatorFunc(sqlExec sql.ExecFunc) (CreatorFunc, error) {
	if sqlExec == nil {
		return nil, ErrCantMakeTheInjection
	}
	return func(user *DTO) error {
		query := "INSERT INTO users SET username=?, password=?,Name=?, Age=?"
		if _, err := sqlExec(query, user.Username, user.Password, user.Name, user.Age); err != nil {
			return ErrCantCreateUser
		}
		return nil
	}, nil

}

//Get a list of users
func MakeGetList(sqlQuery sql.QueryFunc) (GetListFunc, error) {
	if sqlQuery == nil {
		return nil, ErrCantMakeTheInjection
	}
	return func() (Users, error) {
		query := "SELECT id, username, password, email FROM users"
		users := Users{}
		rows, err := sql.Query(query)
		if err != nil {
			return users, ErrCantGetUserList
		}

		for rows.Next() {
			user := DTO{}
			rows.Scan(&user.ID, &user.Username, &user.Password, &user.Name, &user.Age)
			users = append(users, user)
		}

		return users, nil
	}, nil
}

//Get a user by a given ID
func MakeGetByIDFunc(sqlQuery sql.QueryFunc) (GetByIDFunc, error) {
	if sqlQuery == nil {
		return nil, ErrCantMakeTheInjection
	}
	return func(id int) (DTO, error) {
		user := DTO{}
		query := "SELECT * FROM users WHERE id=?"
		rows, err := sql.Query(query, id)
		if err != nil {
			return user, ErrCantGetUserByID
		}
		for rows.Next() {
			rows.Scan(&user.ID, &user.Username, &user.Password, &user.Name, &user.Age)
		}
		return user, nil
	}, nil
}

//Update user info
func MakeUpdate(sqlExec sql.ExecFunc) (UpdateByIDFunc, error) {
	if sqlExec == nil {
		return nil, ErrCantMakeTheInjection
	}
	return func(user DTO) error {
		query := "UPDATE users SET username=?, password=?, name=?, age=? WHERE id=?"
		if _, err := sql.Exec(query, user.Username, user.Password, user.Name, user.Age, user.ID); err != nil {
			return ErrCantUpdateUserByID
		}
		return nil
	}, nil
}
