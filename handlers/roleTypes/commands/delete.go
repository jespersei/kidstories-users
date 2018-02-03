package roleTypesCommands

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
	"github.com/kidstories/users/models"
)

// Delete an roles
func Delete(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	roles := models.Roles{}
	err := c.Bind(&roles)
	if err != nil {
		c.Error(err)
		return
	}
	query := bson.M{"_id": roles.ID}
	err = db.C(models.CollectionRoles).Remove(query)
	if err != nil {
		c.Error(err)
		return
	}
}
