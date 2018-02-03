package roleTypesCommands

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
	"github.com/kidstories/users/models"
)

// Update an roles
func Update(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	roles := models.Roles{}
	err := c.Bind(&roles)
	if err != nil {
		c.Error(err)
		return
	}

	query := bson.M{"_id": roles.ID}
	doc := bson.M{
		"code":     roles.Code,
		"roleName": roles.RoleName,
	}
	err = db.C(models.CollectionRoles).Update(query, doc)
	if err != nil {
		c.Error(err)
		return
	}
}
