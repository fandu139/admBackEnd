package main

import (
	// "log"
	// "net/http"
	// "os"

	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// func hello(w http.ResponseWriter, r *http.Request) {
// 	io.WriteString(w, "Hello Word")
// }

func main() {
	// route := gin.Default()
	// route.GET("/", func(context *gin.Context) {
	// 	context.JSON(http.StatusOK, gin.H{
	// 		"message": "Server Sudah Jalan",
	// 	})
	// })

	// route.POST("/auth", func(context *gin.Context) {
	// 	// context.JSON(http.StatusOK, gin.H{
	// 	// 	"message": "Server Sudah Jalan",
	// 	// })
	// 	// id := context.Query("id")
	// 	// page := context.DefaultQuery("page", "0")
	// 	name := context.PostForm("name")
	// 	message := context.PostForm("message")

	// 	context.JSON(200, gin.H{
	// 		"status":  "posted",
	// 		"message": message,
	// 		"name":    name,
	// 	})

	// 	// fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	// })
	// route.Run(":8089") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	// port := os.Getenv("PORT")
	// http.HandleFunc("/", hello)
	// http.ListenAndServe(":"+port, nil)

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.Run(":" + port)
}
