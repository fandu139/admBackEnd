package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

// type login struct {
// 	Username string `form:"username" json:"username" binding:"required"`
// 	Password string `form:"password" json:"password" binding:"required"`
// }

// var identityKey = "id"

// func helloHandler(c *gin.Context) {
// 	claims := jwt.ExtractClaims(c)
// 	user, _ := c.Get(identityKey)
// 	c.JSON(200, gin.H{
// 		"userID":   claims[identityKey],
// 		"userName": user.(*User).UserName,
// 		"text":     "Hello World.",
// 	})
// }

// User demo
// type User struct {
// 	UserName  string
// 	FirstName string
// 	LastName  string
// }

func main() {

	// Start Off in Codding if you Run Local
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	router := gin.New()
	router.Use(gin.Logger())
	// End Off in Codding if you Run Local

	// Start Off in Codding if you Run Online
	// router := gin.Default()
	// End Off in Codding if you Run Online

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Server Sudah Jalan",
		})
	})

	router.POST("/auth", func(context *gin.Context) {
		username := context.PostForm("username")
		password := context.PostForm("password")

		if username == "" && password == "" {
			context.JSON(200, gin.H{
				"status": "success",
				"token":  username,
			})
		} else {
			context.JSON(400, gin.H{
				"status": "failed",
				"token":  "",
			})
		}

		// fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})
	router.Run(":" + port) // Change Port if you run Local
	// router.Run(":8089")
}

// func main() {
// 	port := os.Getenv("PORT")
// 	// gin.SetMode(gin.ReleaseMode)
// 	router := gin.New()
// 	router.Use(gin.Logger())
// 	router.Use(gin.Recovery())

// 	if port == "" {
// 		port = "8000"
// 	}

// 	router.GET("/", func(context *gin.Context) {
// 		context.JSON(http.StatusOK, gin.H{
// 			"message": "Server Sudah Jalan",
// 		})
// 	})

// 	// the jwt middleware
// 	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
// 		Realm:       "test zone",
// 		Key:         []byte("secret key"),
// 		Timeout:     time.Hour,
// 		MaxRefresh:  time.Hour,
// 		IdentityKey: identityKey,
// 		PayloadFunc: func(data interface{}) jwt.MapClaims {
// 			if v, ok := data.(*User); ok {
// 				return jwt.MapClaims{
// 					identityKey: v.UserName,
// 				}
// 			}
// 			return jwt.MapClaims{}
// 		},
// 		IdentityHandler: func(c *gin.Context) interface{} {
// 			claims := jwt.ExtractClaims(c)
// 			return &User{
// 				UserName: claims[identityKey].(string),
// 			}
// 		},
// 		Authenticator: func(c *gin.Context) (interface{}, error) {
// 			var loginVals login
// 			if err := c.ShouldBind(&loginVals); err != nil {
// 				return "", jwt.ErrMissingLoginValues
// 			}
// 			userID := loginVals.Username
// 			password := loginVals.Password

// 			if (userID == "admin" && password == "admin") || (userID == "test" && password == "test") {
// 				return &User{
// 					UserName:  userID,
// 					LastName:  "Bo-Yi",
// 					FirstName: "Wu",
// 				}, nil
// 			}

// 			return nil, jwt.ErrFailedAuthentication
// 		},
// 		Authorizator: func(data interface{}, c *gin.Context) bool {
// 			if v, ok := data.(*User); ok && v.UserName == "admin" {
// 				return true
// 			}

// 			return false
// 		},
// 		Unauthorized: func(c *gin.Context, code int, message string) {
// 			c.JSON(code, gin.H{
// 				"code":    code,
// 				"message": message,
// 			})
// 		},
// 		// TokenLookup is a string in the form of "<source>:<name>" that is used
// 		// to extract token from the request.
// 		// Optional. Default value "header:Authorization".
// 		// Possible values:
// 		// - "header:<name>"
// 		// - "query:<name>"
// 		// - "cookie:<name>"
// 		// - "param:<name>"
// 		TokenLookup: "header: Authorization, query: token, cookie: jwt",
// 		// TokenLookup: "query:token",
// 		// TokenLookup: "cookie:token",

// 		// TokenHeadName is a string in the header. Default value is "Bearer"
// 		TokenHeadName: "Bearer",

// 		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
// 		TimeFunc: time.Now,
// 	})

// 	if err != nil {
// 		log.Fatal("JWT Error:" + err.Error())
// 	}

// 	router.POST("/login", authMiddleware.LoginHandler)

// 	router.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
// 		claims := jwt.ExtractClaims(c)
// 		log.Printf("NoRoute claims: %#v\n", claims)
// 		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
// 	})

// 	auth := router.Group("/auth")
// 	// Refresh time can be longer than token timeout
// 	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
// 	auth.Use(authMiddleware.MiddlewareFunc())
// 	{
// 		auth.GET("/hello", helloHandler)
// 	}

// 	if err := http.ListenAndServe(":"+port, router); err != nil {
// 		log.Fatal(err)
// 	}
// }
