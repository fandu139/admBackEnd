// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/dgrijalva/jwt-go"
// )

// var mySigningKey = []byte("captainjacksparrowsayshi")

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello World")
// 	fmt.Println("Endpoint Hit: homePage")
// }

// func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 		if r.Header["Token"] != nil {

// 			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
// 				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 					return nil, fmt.Errorf("There was an error")
// 				}
// 				return mySigningKey, nil
// 			})

// 			if err != nil {
// 				fmt.Fprintf(w, err.Error())
// 			}

// 			if token.Valid {
// 				endpoint(w, r)
// 			}
// 		} else {
// 			fmt.Fprintf(w, "Not Authorized")
// 		}
// 	})
// }

// func main() {

// 	// Start Off in Codding if you Run Local
// 	// port := os.Getenv("PORT")
// 	// if port == "" {
// 	// 	log.Fatal("$PORT must be set")
// 	// }
// 	// router := gin.New()
// 	// router.Use(gin.Logger())
// 	// End Off in Codding if you Run Local

// 	// router := gin.Default()

// 	http.Handle("/", isAuthorized(homePage))
// 	log.Fatal(http.ListenAndServe(":9000", nil))

// 	// router.GET("/", func(context *gin.Context) {
// 	// 	context.JSON(http.StatusOK, gin.H{
// 	// 		"message": "Server Sudah Jalan",
// 	// 	})
// 	// })

// 	// router.POST("/auth", func(context *gin.Context) {
// 	// 	username := context.PostForm("username")
// 	// 	password := context.PostForm("password")

// 	// 	if username == "" && password == "" {
// 	// 		context.JSON(200, gin.H{
// 	// 			"status": "success",
// 	// 			"token":  username,
// 	// 		})
// 	// 	} else {
// 	// 		context.JSON(400, gin.H{
// 	// 			"status": "failed",
// 	// 			"token":  "",
// 	// 		})
// 	// 	}

// 	// fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
// 	// })
// 	// router.Run(":" + port) // Change Port if you run Local
// 	// router.Run(":8089")
// }
