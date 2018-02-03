package usersCommands

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
	"github.com/kidstories/users/models"
)

// Delete an user
func Delete(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	user := models.User{}
	err := c.Bind(&user)
	if err != nil {
		c.Error(err)
		return
	}
	query := bson.M{"_id": user.ID}
	err = db.C(models.CollectionUser).Remove(query)
	if err != nil {
		c.Error(err)
		return
	}
}
