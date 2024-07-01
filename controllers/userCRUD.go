package controllers

import (
	"net/http"
	"smallBotBackend/initializers"
	"smallBotBackend/models"
	"github.com/gin-gonic/gin"
)

// CreateUser godoc
// @Summary Create User
// @Description Create User
// @Tags users
// @Accept json
// @Produce json
// @Param input body CreateUserInput true "Create User Input"
// @Success 200 {object} User
// @Router /newuser [post]

func CreateUser(c *gin.Context) {

	var createUserInput models.CreateUserInput

	if err := c.ShouldBindJSON(&createUserInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFound models.User
	initializers.DB.Where("telegram_id=?", createUserInput.Telegram_ID).Find(&userFound)

	if userFound.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Telegram ID already used"})
		return
	}

	user := models.User{
		Telegram_ID: createUserInput.Telegram_ID,
		Language: createUserInput.Language,
		Email:    createUserInput.Email,
		RoleID:   createUserInput.RoleID,
	}

	initializers.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// GetUsers godoc
// @Summary Get Users
// @Description Get Users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} User
// @Router /users [get]

func GetUsers(c *gin.Context) {

	var users []models.User
	initializers.DB.Preload("Role").Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})

}

// GetUserByID godoc
// @Summary Get User By ID
// @Description Get User By ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} User
// @Router /users/{id} [get]

func GetUserByID(c *gin.Context) {

	id := c.Param("id")

	var user models.User
	initializers.DB.Preload("Role").First(&user, id)

	c.JSON(http.StatusOK, gin.H{"data": user})

}
// GetUserFeedback godoc
// @Summary Get User Feedback
// @Description Get User Feedback
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} UserFeedback
// @Router /users/feedback [get]

func GetUserFeedback(c *gin.Context) {

	var userFeedback []models.UserFeedback
	initializers.DB.Preload("User").Find(&userFeedback)

	c.JSON(http.StatusOK, gin.H{"data": userFeedback})

}

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

func UpdateUserByID(c *gin.Context) {

	id := c.Param("id")

	var user models.User
	initializers.DB.First(&user, id)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	var updateUserInput models.UpdateUserInput
	if err := c.ShouldBindJSON(&updateUserInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	initializers.DB.Model(&user).Updates(models.User{
		Telegram_ID: updateUserInput.Telegram_ID,
		Language: updateUserInput.Language,
		Email:    updateUserInput.Email,
		RoleID:   updateUserInput.RoleID,
	})

	c.JSON(http.StatusOK, gin.H{"data": user})

}

// DeleteUserByID godoc
// @Summary Delete User By ID
// @Description Delete User By ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {string} string
// @Router /users/{id} [delete]

func DeleteUserByID(c *gin.Context) {

	id := c.Param("id")

	var user models.User
	initializers.DB.First(&user, id)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	initializers.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": "user deleted"})
}