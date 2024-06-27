package controllers

import(
	"net/http"
	"os"
	"time"
	"smallBotBackend/models"
	"smallBotBackend/initializers"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// Login godoc
// @Summary Login admin users
// @Description Sign in with username and password
// @Tags admin
// @Accept json
// @Produce json
// @Param input body AuthInput true "Auth Input"
// @Success 200 {string} token
// @Router /login [post]

func Login (c*gin.Context){
	
	var authInput models.AuthInput

	if err := c.ShouldBindJSON(&authInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFound models.AdminUser
	initializers.DB.Where("username=?", authInput.Username).Find(&userFound)

	if userFound.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(authInput.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
		return
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userFound.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := generateToken.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to generate token"})
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}

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
	initializers.DB.Where("username=?", createUserInput.Username).Find(&userFound)

	if userFound.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username already used"})
		return
	}

	user := models.User{
		Username: createUserInput.Username,
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

// CreateUserFeedback godoc
// @Summary Create User Feedback
// @Description Create User Feedback
// @Tags users
// @Accept json
// @Produce json
// @Param input body UserFeedbackInput true "User Feedback Input"
// @Success 200 {object} UserFeedback
// @Router /users/feedback [post]

func CreateUserFeedback(c *gin.Context) {
	
	var userFeedbackInput models.UserFeedbackInput

	if err := c.ShouldBindJSON(&userFeedbackInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFound models.User
	initializers.DB.Where("id=?", userFeedbackInput.UserID).Find(&userFound)

	if userFound.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	userFeedback := models.UserFeedback{
		UserID:      userFeedbackInput.UserID,
		Comments:    userFeedbackInput.Comments,
		BotFeedback: userFeedbackInput.BotFeedback,
	}

	initializers.DB.Create(&userFeedback)

	c.JSON(http.StatusOK, gin.H{"data": userFeedback})


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
		Username: updateUserInput.Username,
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