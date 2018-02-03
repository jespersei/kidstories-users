package usersCommands

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
	"github.com/kidstories/users/models"
)

// HashPassword ...
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash ...
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Create an user
func Create(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	user := models.User{}
	err := c.Bind(&user)
	if err != nil {
		c.Error(err)
		return
	}

	i := bson.NewObjectId()
	user.ID = i
	password := user.Password
	hash, _ := HashPassword(password)
	user.Password = hash

	user.CreatedOn = time.Now().UnixNano() / int64(time.Millisecond)
	user.UpdatedOn = time.Now().UnixNano() / int64(time.Millisecond)

	query := bson.M{
		"username": user.Username,
	}
	//existing := true
	err = db.C(models.CollectionUser).Find(query).Sort("-_id").One(&user)
	if err == nil {
		fmt.Println("error", err)
		fmt.Println("existing", user)
		c.Error(err)
		return
	}

	//if existing == true {
	err = db.C(models.CollectionUser).Insert(user)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, user)

}
