package user

import "errors"

var (
	ErrCantCreateUser       = errors.New("can't create user")
	ErrCantMakeTheInjection = errors.New("can't make the injection")
	ErrCantGetUserList      = errors.New("can't get user list")
	ErrCantGetUserByID      = errors.New("can't get user by id")
	ErrCantUpdateUserByID   = errors.New("can't get user by id")
	ErrInvalidBody          = errors.New("invalid body")
)
