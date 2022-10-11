package api

import (
	"database/sql"
	"net/http"

	db "github.com/enevarez1/go-exercise/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	UserName string `json:"user_name" binding:required`
	FullName string `json:"full_name" binding:required`
	Email    string `json:"email" binding:required`
	Password string `json:"password" binding:required`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		UserName: req.UserName,
		FullName: req.FullName,
		Email: req.Email,
		Password: req.Password,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

type getUserRequest struct {
	ID string `json: "user_name"`
}

func (server *Server) getUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, user)
}

type updateUserRequest struct {
	FullName string `json:"full_name" binding:required`
	Password string `json:"password" binding:required`
	ID string `json: "user_name"`
}

func (server *Server) updateUser(ctx *gin.Context) {
	var req updateUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateUserParams{
		UserName: req.ID,
		FullName: req.FullName,
		Password: req.Password,
	}

	err := server.store.UpdateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

type deleteUserRequest struct {
	ID string `json: "user_name"`
}

func (server *Server) deleteUser(ctx *gin.Context) {
	var req deleteUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteUser(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, nil)
}