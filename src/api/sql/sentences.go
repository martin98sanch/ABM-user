package sql

const (
	InsertUserStatement         = "INSERT INTO users SET username=?, password=?,Name=?, Age=?"
	SelectAllUsersStatement     = "SELECT * FROM users"
	SelectUserByIDStatement     = "SELECT id, username, password, name, age, deleted FROM users WHERE id=?"
	UpdateUserByIDStatement     = "UPDATE users SET username=?, password=?, name=?, age=? WHERE id=?"
	DeleteUserByIDStatement     = "UPDATE users SET deleted=? WHERE id=?"
	FindUserByUsernameStatement = "SELECT id FROM users WHERE username=?"
)
