package usersQueries

import (
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/gin-gonic/gin"
	"github.com/kidstories/users/models"
)

// List all users
func List(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	users := []models.User{}
	err := db.C(models.CollectionUser).Find(nil).Sort("-_id").All(&users)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, users)
}
