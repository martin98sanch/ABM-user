package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/martin98sanch/ABM-user/src/api/server/request"
	"github.com/martin98sanch/ABM-user/src/api/server/response"
)

type Handler struct {
	UserCreator CreatorFunc
}

var (
	ParamBody     = "body"
	ParamUserName = "username"
)

//ValidateCreate func make the validation for create a user
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
		response.Make(ctx, http.StatusInternalServerError, ErrCantCreateUser)
		return
	}
	response.Make(ctx, http.StatusNoContent)
	return
}

//TODO: Ya esta en teoria todo menos el eliminado de un usuario en la base de datos
// Queda pendiente crear en http.go para usar las demas funciones y probar
// Dar cobertura desp
