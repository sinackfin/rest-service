package handler

import (
	"github.com/gin-gonic/gin"
	"api/internal/middleware"
	"api/internal/service"
)

type ResponseJSON struct {
	Success		bool	`json:"success"`
	Msg			string	`json:"msg"`
}

type Handler struct{
	HTTPHandler *gin.Engine
	PersonService	*service.PersonService
}


func New(ps *service.PersonService) *Handler{
	var h Handler 
	router := gin.Default()
	router.Use(middleware.JSONLogMiddleware())
	api := router.Group("/api")
	{
		api.GET("/person",h.GetPersonByID)
		api.PATCH("/person/:id",h.UpdatePerson)
		api.DELETE("/person/:id",h.DeletePersonByID)
		api.POST("/person",h.CreatePerson)
	}
	h.HTTPHandler 	= router
	h.PersonService = ps
	return &h
}


func ResponseErrorJSON(c *gin.Context,err error) *ResponseJSON {
	errResp := ResponseJSON{
		false,
		err.Error(),
	}
	c.JSON(503,errResp)
	c.Abort()
	return &errResp
}