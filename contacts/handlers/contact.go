package handlers

import (
	"contacts/interfaces"
	"contacts/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Contact struct {
	IContact interfaces.IContact
}

func (c *Contact) Get(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if id == "" || !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		ctx.Abort()
	}

	contact, err := c.IContact.Get(id)
	if err != nil {
		ctx.String(400, err.Error()) // dont directly send db error in production systems
		ctx.Abort()
	}

	ctx.JSON(200, contact)
	ctx.Abort()

}

func (c *Contact) Create(ctx *gin.Context) {
	contact := new(models.Contact)

	if err := ctx.Bind(contact); err != nil { // automatically data is read from the request body and bind to the contact
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "wrong data given:" + err.Error(),
		})
	}
	//contact.ID = 1
	contact.Status = "active"
	contact.LastModified = time.Now().Unix()
	contact, err := c.IContact.Add(contact) // interface driven;kind of dependency injection
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
	}
	ctx.JSON(201, contact) // automatically data is returned to the response
	ctx.Abort()
}
