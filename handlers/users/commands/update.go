package usersCommands

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
	"github.com/kidstories/users/models"
)

// Update an user
func Update(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	user := models.User{}
	err := c.Bind(&user)
	if err != nil {
		c.Error(err)
		return
	}

	password := user.Password
	hash, _ := HashPassword(password)
	user.Password = hash

	query := bson.M{"_id": user.ID}
	doc := bson.M{
		"username":  user.Username,
		"password":  user.Password,
		"email":     user.Email,
		"role":      user.Role,
		"updatedOn": time.Now().UnixNano() / int64(time.Millisecond),
	}
	err = db.C(models.CollectionUser).Update(query, doc)
	if err != nil {
		c.Error(err)
		return
	}
}
