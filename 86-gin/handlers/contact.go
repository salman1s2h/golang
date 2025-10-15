package handlers

import (
	"demo/database"
	"demo/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type IContactHandler interface {
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	DeleteByID(ctx *gin.Context)
}

type Contacthandler struct {
	IContactDB database.IContactDB
}

func NewContacthandler(icontactdb database.IContactDB) IContactHandler {
	return &Contacthandler{IContactDB: icontactdb}
}

func (ch *Contacthandler) Create(ctx *gin.Context) {
	contact := new(models.Contact)
	err := ctx.Bind(contact)
	if err != nil {
		log.Println(err.Error())
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	contact.Lastmodified = time.Now().Unix()
	contact.Status = "active"
	err = contact.Validate()
	if err != nil {
		log.Println(err.Error())
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	contact, err = ch.IContactDB.Create(contact)
	if err != nil {
		log.Println(err.Error())
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, contact)

}
func (ch *Contacthandler) GetAll(ctx *gin.Context) {
	contacts, err := ch.IContactDB.GetAll()
	if err != nil {
		log.Println(err.Error())
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(200, contacts)

}

func (ch *Contacthandler) DeleteByID(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		log.Println("id not provided")
		ctx.String(http.StatusBadRequest, "id not provided or invalid id")
		return
	}

	err := ch.IContactDB.DeleteByID(id)
	if err != nil {
		log.Println(err.Error())
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	ctx.String(http.StatusNoContent, "record successfully deleted")

}

func (ch *Contacthandler) GetByID(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		log.Println("id not provided")
		ctx.String(http.StatusBadRequest, "id not provided or invalid id")
		return
	}
	contact, err := ch.IContactDB.GetByID(id)
	if err != nil {
		log.Println(err.Error())
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, contact)
}
