package main

import (
	"smallBotBackend/controllers"
	"smallBotBackend/initializers"
	// "smallBotBackend/middlewares"

    "github.com/gin-gonic/gin"
)

func init(){
    initializers.LoadEnvs()
    initializers.ConnectDB()
}

func main(){
    router := gin.Default()
    router.ForwardedByClientIP = true
    router.SetTrustedProxies([]string {"127.0.0.1", "192.168.8.180"})

    router.POST("/login", controllers.Login)
    router.POST("/newuser", controllers.CreateUser)
    router.GET("/users", controllers.GetUser)

	router.NoRoute(func(c*gin.Context){
		c.JSON(404, gin.H{"message": "Page not found"})
	})

    router.Run(":8080")
}