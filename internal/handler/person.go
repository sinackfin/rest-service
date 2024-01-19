package handler

import (
	"github.com/gin-gonic/gin"
	"api/internal/models"
	log "github.com/sirupsen/logrus" 
	"errors"
	"net/http"
)

func (h *Handler) GetPersonByID(c *gin.Context){
	id, ok := c.GetQuery("id")
	if !ok {
		ResponseErrorJSON(c,errors.New("Incorrect ID"))
		return
	}
	person,err := h.PersonService.GetPerson(id)
	if err != nil {
		log.Error(err)
		ResponseErrorJSON(c,err)
		return
	}
	c.JSON(http.StatusOK, person)
}


func (h *Handler) CreatePerson(c *gin.Context){
	person 	:= models.Person{}
	if err := c.BindJSON(&person); err != nil {
		log.Error(err)
		ResponseErrorJSON(c,err)
		return
	}
	if err := h.PersonService.CreatePerson(&person); err != nil {
		log.Error(err)
		ResponseErrorJSON(c,err)
		return
	}
	resp := ResponseJSON{
		true,
		"Person has been created",
	}
	c.JSON(http.StatusCreated, resp)
}

func (h *Handler) DeletePersonByID(c *gin.Context){
	person_id := c.Param("id")
	if err := h.PersonService.DeletePerson(person_id); err != nil {
		log.Error(err)
		ResponseErrorJSON(c,err)
		return
	}
	resp := ResponseJSON{
		true,
		"Person has been deleted",
	}
	c.JSON(http.StatusAccepted, resp)
}

func (h *Handler) UpdatePerson(c *gin.Context){
	updated_person 	:= models.Person{}
	if err := c.BindJSON(&updated_person); err != nil {
		log.Error(err)
		ResponseErrorJSON(c,err)
		return
	}
	updated_person.ID = c.Param("id")
	if err := h.PersonService.UpdatePerson(&updated_person); err != nil {
		log.Error(err)
		ResponseErrorJSON(c,err)
		return
	}
	resp := ResponseJSON{
		true,
		"Person has been Updated",
	}
	c.JSON(http.StatusOK, resp)
}


