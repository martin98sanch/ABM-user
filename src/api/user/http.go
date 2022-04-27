package user

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/martin98sanch/ABM-user/src/api/server/request"
	"github.com/martin98sanch/ABM-user/src/api/server/response"
)

type Handler struct {
	UserCreator CreatorFunc
	GetUserList GetListFunc
	GetByID     GetByIDFunc
	DeleteByID  DeleteByIDFunc
	UpdateByID  UpdateByIDFunc
}

var (
	ParamBody     = "body"
	ParamUserName = "username"
	ParamUserID   = "user_id"
)

// ValidateCreate func make the validation for create a user
func (handler Handler) ValidateCreate(ctx *gin.Context) {
	var body DTO
	err := request.GetJsonBody(ctx, &body)
	if err != nil {
		response.Make(ctx, http.StatusBadRequest, ErrInvalidBody)
		return
	}
	if err := body.Validate(); err != nil {
		response.Make(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.Set(ParamBody, body)
}

// Create func persists in database user data
func (handler Handler) Create(ctx *gin.Context) {
	user := ctx.MustGet(ParamBody).(DTO)

	if err := handler.UserCreator(&user); err != nil {
		if errors.Is(err, ErrUserAlreadyExist) {
			response.Make(ctx, http.StatusNotAcceptable, err)
			return
		}
		response.Make(ctx, http.StatusInternalServerError, ErrCantCreateUser)
		return
	}
	response.Make(ctx, http.StatusNoContent)
	return
}

// GetList func get a list of all users in the database
func (handler Handler) GetList(ctx *gin.Context) {
	userList, err := handler.GetUserList()
	if err != nil {
		response.Make(ctx, http.StatusInternalServerError, ErrCantGetUserList)
		return
	}

	response.Make(ctx, http.StatusOK, userList)
}

//ValidateGet func get a user by a given ID
func (handler Handler) ValidateGet(ctx *gin.Context) {
	userIDParam := ctx.Param(ParamUserID)
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		response.Make(ctx, http.StatusBadRequest, ErrInvalidUserID)
	}
	ctx.Set(ParamUserID, ID(userID))
}

// Get func returns user info from a given ID
func (handler Handler) Get(ctx *gin.Context) {
	userID := ctx.MustGet(ParamUserID).(ID)

	user, err := handler.GetByID(userID)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			response.Make(ctx, http.StatusNotFound, err)
			return
		}
		response.Make(ctx, http.StatusInternalServerError, ErrCantGetUserByID)
		return
	}
	response.Make(ctx, http.StatusOK, user)
	return
}

// ValidateDelete func mark a user as deleted by a given ID
func (handler Handler) ValidateDelete(ctx *gin.Context) {
	userIDParam := ctx.Param(ParamUserID)
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		response.Make(ctx, http.StatusBadRequest, ErrInvalidUserID)
	}
	ctx.Set(ParamUserID, ID(userID))
}

// Delete func returns user info from a given ID
func (handler Handler) Delete(ctx *gin.Context) {
	userID := ctx.MustGet(ParamUserID).(ID)

	if err := handler.DeleteByID(userID); err != nil {
		response.Make(ctx, http.StatusInternalServerError, ErrCantDeleteUserByID)
		return
	}
	response.Make(ctx, http.StatusOK)
	return
}

// ValidatePut func update user info by a given ID
func (handler Handler) ValidatePut(ctx *gin.Context) {
	userIDParam := ctx.Param(ParamUserID)
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		response.Make(ctx, http.StatusBadRequest, ErrInvalidUserID)
	}
	ctx.Set(ParamUserID, ID(userID))

	var body DTO
	err = request.GetJsonBody(ctx, &body)
	if err != nil {
		response.Make(ctx, http.StatusBadRequest, ErrInvalidBody)
		return
	}
	if err := body.Validate(); err != nil {
		response.Make(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.Set(ParamBody, body)

}

// Put func update user info by a given ID
func (handler Handler) Put(ctx *gin.Context) {
	userID := ctx.MustGet(ParamUserID).(ID)
	newUserInfo := ctx.MustGet(ParamBody).(DTO)
	newUserInfo.ID = userID

	if err := handler.UpdateByID(&newUserInfo); err != nil {
		if errors.Is(err, ErrUserAlreadyExist) {
			response.Make(ctx, http.StatusNotAcceptable, err)
			return
		}
		response.Make(ctx, http.StatusInternalServerError, ErrCantUpdateUserByID)
		return
	}
	response.Make(ctx, http.StatusOK)
	return
}

// TODO: Dar cobertura a todo
// TODO: Consumir info de los yml (en chrome del trabajo deje un proyecto que tiene todo armado)
// TODO: Hacer el controlador y las vistas
