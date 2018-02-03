package config

import (
	"github.com/gin-gonic/gin"
	"github.com/kidstories/users/db"
)

// Port at which the server starts listening
const Port = "7001"

func init() {
	db.Connect()
}

// OptionsUser : this is add methods in login user
func OptionsUser(c *gin.Context) {
	c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, DELETE, POST, PATCH, PUT, OPTIONS")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With")
	c.Next()
}
