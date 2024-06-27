package main

import (
	"smallBotBackend/controllers"
	"smallBotBackend/initializers"
	// "smallBotBackend/middlewares"
    ginSwagger "github.com/swaggo/gin-swagger"
    swaggerFiles "github.com/swaggo/files"
    docs "smallBotBackend/docs"
    _ "github.com/swaggo/gin-swagger/example/basic/docs"

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
    router.SetTrustedProxies([]string {"127.0.0.1", "192.168.8.180"})

    url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

    // Login godoc
    // @Summary Login admin users
    // @Description Sign in with username and password
    // @Tags admin
    // @Accept json
    // @Produce json
    // @Param input body AuthInput true "Auth Input"
    // @Success 200 {string} token
    // @Router /login [post]
    router.POST("/login", controllers.Login)

    // CreateUser godoc
    // @Summary Create User
    // @Description Create User
    // @Tags users
    // @Accept json
    // @Produce json
    // @Param input body CreateUserInput true "Create User Input"
    // @Success 200 {object} User
    // @Router /newuser [post]
    router.POST("/newuser", controllers.CreateUser)

    // GetUsers godoc
    // @Summary Get Users
    // @Description Get Users
    // @Tags users
    // @Accept json
    // @Produce json
    // @Success 200 {object} User
    // @Router /users [get]
    router.GET("/users", controllers.GetUsers)

    // GetUserByID godoc
    // @Summary Get User By ID
    // @Description Get User By ID
    // @Tags users
    // @Accept json
    // @Produce json
    // @Param id path string true "User ID"
    // @Success 200 {object} User
    // @Router /users/{id} [get]
    router.GET("/users/:id", controllers.GetUserByID)

    // CreateUserFeedback godoc
    // @Summary Create User Feedback
    // @Description Create User Feedback
    // @Tags users
    // @Accept json
    // @Produce json
    // @Param input body UserFeedbackInput true "User Feedback Input"
    // @Success 200 {object} UserFeedback
    // @Router /users/feedback [post]
    router.POST("/users/feedback", controllers.CreateUserFeedback)

    // GetUserFeedback godoc
    // @Summary Get User Feedback
    // @Description Get User Feedback
    // @Tags users
    // @Accept json
    // @Produce json
    // @Success 200 {object} UserFeedback
    // @Router /users/feedback [get]
    router.GET("/users/feedback", controllers.GetUserFeedback)

    // UpdateUserByID godoc
    // @Summary Update User By ID
    // @Description Update User By ID
    // @Tags users
    // @Accept json
    // @Produce json
    // @Param id path string true "User ID"
    // @Param input body UpdateUserInput true "Update User Input"
    // @Success 200 {object} User
    // @Router /users/{id} [patch]
    router.PATCH("/users/:id", controllers.UpdateUserByID)
    // DeleteUserByID godoc
    // @Summary Delete User By ID
    // @Description Delete User By ID
    // @Tags users
    // @Accept json
    // @Produce json
    // @Param id path string true "User ID"
    // @Success 200 {string} string
    // @Router /users/{id} [delete]
    router.DELETE("/users/:id", controllers.DeleteUserByID)

	router.NoRoute(func(c*gin.Context){
		c.JSON(404, gin.H{"message": "Page not found"})
	})

    router.Run(":8080")
}