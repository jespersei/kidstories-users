package roleTypesCommands

import (
	"gopkg.in/mgo.v2"

	"github.com/gin-gonic/gin"
	"github.com/kidstories/users/models"
)

// Create an roles
func Create(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	roles := models.Roles{}
	err := c.Bind(&roles)
	if err != nil {
		c.Error(err)
		return
	}

	err = db.C(models.CollectionRoles).Insert(roles)
	if err != nil {
		c.Error(err)
		return
	}
}
