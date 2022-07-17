package api

import (
	"context"
	"database/sql"
	db "github.com/freedommmoto/metamaskonline_backend/model/sqlc"
	"github.com/freedommmoto/metamaskonline_backend/tool"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type UserResponse struct {
	ID         int32  `json:"id"`
	Username   string `json:"username"`
	Validation bool   `json:"validation"`
	Code       string `json:"code"`
}

type AddUseInput struct {
	Username string `json:"username" binding:"required,min=4"`
	Password string `json:"password" binding:"required,min=6"`
	Wallet   string `json:"wallet" binding:"required,min=38"`
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
	LineOwnerValidation, errSelectLineOwner := server.store.SelectLineOwnerValidationByID(context.Background(), request.Id)
	if errSelectLineOwner != nil {
		ctx.JSON(http.StatusInternalServerError, makeErrorReturnFormat(err))
		return
	}

	ctx.JSON(http.StatusOK, tranFerUserToUserResponse(user, LineOwnerValidation))
}
func (server *Server) addNewUser(ctx *gin.Context) {
	var request AddUseInput
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, makeErrorReturnFormat(err))
		return
	}

	hashPassword, err := tool.HashPassword(request.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, makeErrorReturnFormat(err))
		return
	}

	//insert user
	insertParam := db.InsertNewUserParams{
		Username: request.Username,
		Password: hashPassword,
	}
	user, errInsertUser := server.store.InsertNewUser(context.Background(), insertParam)
	if errInsertUser != nil {
		log.Println("error InsertNewUser", errInsertUser)
		ctx.JSON(http.StatusInternalServerError, makeErrorReturnFormat(errInsertUser))
		return
	}

	//insert code
	insertCodeParam := db.InsertLineOwnerValidationParams{
		IDUser:    user.IDUser,
		Code:      tool.RandomCodeNumber(4, server.store),
		CreatedAt: time.Now(),
	}
	LineOwnerValidation, errInsertCode := server.store.InsertLineOwnerValidation(context.Background(), insertCodeParam)
	if errInsertCode != nil {
		log.Println("error InsertLineOwnerValidation", errInsertCode)
		ctx.JSON(http.StatusInternalServerError, makeErrorReturnFormat(err))
		return
	}

	insertWalletParam := db.AddNewWalletParams{
		MetamaskWalletID: request.Wallet,
		IDUser:           user.IDUser,
	}
	//insert Wallet
	_, errInsertWallet := server.store.AddNewWallet(context.Background(), insertWalletParam)
	if errInsertWallet != nil {
		log.Println("error AddNewWallet", errInsertWallet)
		ctx.JSON(http.StatusInternalServerError, makeErrorReturnFormat(err))
		return
	}

	ctx.JSON(http.StatusOK, tranFerUserToUserResponse(user, LineOwnerValidation))
}

func tranFerUserToUserResponse(user db.User, code db.LineOwnerValidation) UserResponse {
	newUser := UserResponse{
		ID:         user.IDUser,
		Username:   user.Username,
		Validation: user.OwnerValidation,
		Code:       code.Code,
	}
	return newUser
}

func makeErrorReturnFormat(err error) gin.H {
	return gin.H{"error": err.Error()}
}
