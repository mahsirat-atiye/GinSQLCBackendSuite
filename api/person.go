package api

import (
	db "example/backend/db/sqlc"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createPersonRequest struct {
	Name string `json:"name" binding:"required"`
	Age  string `json:"age" binding:"required"`
}

func (server *Server) createPerson(ctx *gin.Context) {
	var req createPersonRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.CreatePersonParams{
		Name: req.Name,
		Age:  req.Age,
	}

	person, err := server.store.CreatePerson(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, person)
}
