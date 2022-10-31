package api

import (
	"database/sql"
	"net/http"

	db "github.com/enevarez1/go-exercise/db/sqlc"
	"github.com/enevarez1/go-exercise/util"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createUserRequest struct {
	UserName string `json:"UserName" binding:"required"`
	FullName string `json:"FullName" binding:"required"`
	Email    string `json:"Email" binding:"required"`
	Password string `json:"Password" binding:"required,min=6"`
}

type createUserResponse struct {
	ID          int32        `json:"ID"`
	UserName    string       `json:"UserName"`
	FullName    string       `json:"FullName"`
	Email       string       `json:"Email"`
	CreatedAt   sql.NullTime `json:"Created_At"`
	LastUpdated sql.NullTime `json:"Last_Updated"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		UserName: req.UserName,
		FullName: req.FullName,
		Email: req.Email,
		Password: hashedPassword,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := createUserResponse {
		UserName: user.UserName,
		FullName: user.FullName,
		Email: user.Email,
		CreatedAt: user.CreatedAt,
		LastUpdated: user.LastUpdated,
	}

	ctx.JSON(http.StatusCreated, rsp)
}

type getUserRequest struct {
	ID *int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}


	user, err := server.store.GetUser(ctx, *req.ID)
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
	UserName string `json:"UserName" binding:"required"`
	FullName string `json:"FullName" binding:"required"`
	Password string `json:"Password" binding:"required"`
	ID int32 `uri:"id" binding:"required,min=1"`
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
		UserName: req.UserName,
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
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteUser(ctx *gin.Context) {
	var req deleteUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteUser(ctx, int32(req.ID))
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