package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	jwt "gopkg.in/appleboy/gin-jwt.v2"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
	"github.com/gin-mongo-api/models"
	"github.com/kidstories/users/config"
	"github.com/kidstories/users/handlers/roleTypes/commands"
	"github.com/kidstories/users/handlers/roleTypes/queries"
	"github.com/kidstories/users/handlers/users/commands"
	"github.com/kidstories/users/handlers/users/queries"
	"github.com/kidstories/users/middlewares"
	//"github.com/fvbock/endless" //this is for production
	"github.com/itsjamie/gin-cors"
)

func main() {
	router := gin.Default()
	//gin.SetMode(gin.ReleaseMode) // this is for production

	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, DELETE, POST, PATCH, PUT, OPTIONS",
		RequestHeaders:  "Origin, Authorization, Content-Type, X-Requested-With",
		ExposedHeaders:  "",
		MaxAge:          time.Hour * 12,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	router.RedirectTrailingSlash = false
	router.RedirectFixedPath = false

	secretKey := "$2a$14$408pgmdP3nc4x2ZyLZWOFuD3r6jek9uE6I/zccwS5EDbsAmaq5sa2"
	hashSecretKey, _ := usersCommands.HashPassword(secretKey)

	authMiddleware := &jwt.GinJWTMiddleware{
		Realm:      hashSecretKey,
		Key:        []byte(hashSecretKey),
		Timeout:    time.Hour * 12,
		MaxRefresh: time.Hour * 12,
		Authenticator: func(username string, password string, c *gin.Context) (string, bool) {
			db := c.MustGet("db").(*mgo.Database)
			query := bson.M{
				"username": username,
			}
			user := models.User{}
			err := db.C(models.CollectionUser).Find(query).Sort("-_id").One(&user)
			bpassword := password
			match := usersCommands.CheckPasswordHash(bpassword, user.Password)

			credentials := user.Password + user.Username + user.Email + user.Role
			if err != nil || match != true {
				fmt.Println(err)
				c.Error(err)
				return credentials, false
			}

			return credentials, true
		},
		Authorizator: func(username string, c *gin.Context) bool {
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":         code,
				"errorMessage": message,
				"success":      false,
			})

		},
		TokenLookup: "header:Authorization",
	}

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.Connect)
	router.Use(middlewares.ErrorHandler)

	//Routes
	router.POST("/login", authMiddleware.LoginHandler)
	router.OPTIONS("/login", config.OptionsUser)

	auth := router.Group("/api")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)

		// Users
		auth.GET("/users", usersQueries.List)
		auth.POST("/users", usersCommands.Create)
		auth.PATCH("/users", usersCommands.Update)
		auth.DELETE("/users", usersCommands.Delete)

		// RoleType
		auth.GET("/roles", roleTypesQueries.List)
		auth.POST("/roles", roleTypesCommands.Create)
		auth.PATCH("/roles", roleTypesCommands.Update)
		auth.DELETE("/roles", roleTypesCommands.Delete)
	}

	// If the routes is not existing
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "ERROR"})
	})

	http.Handle("/api", router)

	// Start listening
	port := config.Port
	if len(os.Getenv("PORT")) > 0 {
		port = os.Getenv("PORT")
	}
	router.Run(":" + port)
	//endless.ListenAndServe(":"+port, router)
}
