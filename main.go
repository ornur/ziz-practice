package main

import (
    "smallBotBackend/controllers"
	"smallBotBackend/initializers"
	// "smallBotBackend/middlewares"

    ginSwagger "github.com/swaggo/gin-swagger"
    swaggerFiles "github.com/swaggo/files"
    docs "smallBotBackend/docs"

    "github.com/gin-gonic/gin"
)

// @title Small Bot Backend API
// @version 1.0
// @description This is a simple API for a small bot backend
// @BasePath /api/v1
// @host localhost:8080
// @schemes http
// @produce json
// @consumes json
// @paths /login, /newuser, /users, /users/{id}, /users/feedback, /users/{id}, /users/{id}, /users/{id}

func init(){
    initializers.LoadEnvs()
    initializers.ConnectDB()
}

func main(){
    router := gin.Default()

    docs.SwaggerInfo.BasePath = "/api/v1"
    router.ForwardedByClientIP = true
    router.SetTrustedProxies([]string {"127.0.0.1", "192.168.8.180", "192.168.8.228"})

    url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

    router.POST("/admin/createuseradmin", controllers.CreateUserAdmin)
    router.POST("/admin/login", controllers.AdminLogin)
    router.GET("/admin/users", controllers.GetAdminUsers)
    router.GET("/admin/users/:id", controllers.GetAdminUserByID)
    router.DELETE("/admin/users/:id", controllers.DeleteAdminUserByID)

    router.POST("/newuser", controllers.CreateUser)
    router.GET("/users", controllers.GetUsers)
    router.GET("/users/:id", controllers.GetUserByID)
    router.PATCH("/users/:id", controllers.UpdateUserByID)
    router.DELETE("/users/:id", controllers.DeleteUserByID)
    router.GET("/users/feedback", controllers.GetUserFeedback)

    router.GET("/checkuser/:telegram_id", controllers.CheckUser)

	router.NoRoute(func(c*gin.Context){
		c.JSON(404, gin.H{"message": "Page not found"})
	})

    router.Run(":8080")
}