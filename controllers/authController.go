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
		RoleID:   createUserInput.RoleID,
	}

	initializers.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})

}

func GetUsers(c *gin.Context) {

	var users []models.User
	initializers.DB.Preload("Role").Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})

}

func GetUserByID(c *gin.Context) {

	id := c.Param("id")

	var user models.User
	initializers.DB.Preload("Role").First(&user, id)

	c.JSON(http.StatusOK, gin.H{"data": user})

}

func GetUserFeedback(c *gin.Context) {

	var userFeedback []models.UserFeedback
	initializers.DB.Preload("User").Find(&userFeedback)

	c.JSON(http.StatusOK, gin.H{"data": userFeedback})

}

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
		RoleID:   updateUserInput.RoleID,
	})

	c.JSON(http.StatusOK, gin.H{"data": user})

}

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