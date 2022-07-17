package api

import (
	"context"
	"database/sql"
	db "github.com/freedommmoto/metamaskonline_backend/model/sqlc"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserResponse struct {
	Username   string `json:"username"`
	Validation bool   `json:"validation"`
}

type GetUserInput struct {
	//make sure you start with uppercase
	Id int32 `json:"name" uri:"id" binding:"required,min=1""`
}

func (server *Server) getUserByID(ctx *gin.Context) {
	var request GetUserInput
	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, makeErrorReturnFormat(err))
		return
	}
	//id := ctx.Query("id")
	//log.Println(request)

	user, err := server.store.SelectUserID(context.Background(), request.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, makeErrorReturnFormat(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, makeErrorReturnFormat(err))
		return
	}

	ctx.JSON(http.StatusOK, tranFerUserToUserResponse(user))
}

func tranFerUserToUserResponse(user db.User) UserResponse {
	newUser := UserResponse{
		Username:   user.Username,
		Validation: user.OwnerValidation,
	}
	return newUser
}

func makeErrorReturnFormat(err error) gin.H {
	return gin.H{"error": err.Error()}
}
