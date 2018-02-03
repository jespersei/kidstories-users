package roleTypesQueries

import (
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/gin-gonic/gin"
	"github.com/kidstories/users/models"
)

// List all roles
func List(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	roles := []models.Roles{}
	err := db.C(models.CollectionRoles).Find(nil).Sort("-_id").All(&roles)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, roles)
}
