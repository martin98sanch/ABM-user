package user

import "errors"

var (
	ErrCantCreateUser              = errors.New("can't create user")
	ErrCantMakeTheInjection        = errors.New("can't make the injection")
	ErrCantGetUserList             = errors.New("can't get user list")
	ErrCantGetUserByID             = errors.New("can't get user by id")
	ErrCantUpdateUserByID          = errors.New("can't update user by id")
	ErrCantDeleteUserByID          = errors.New("can't delete user by id")
	ErrInvalidBody                 = errors.New("invalid body")
	ErrInvalidUserID               = errors.New("can't get user info because ID is invalid")
	ErrCantCheckIfUserAlreadyExist = errors.New("can't check if user already exist")
	ErrUserAlreadyExist            = errors.New("user already exist")
	ErrNotFound                    = errors.New("user not found")
)
