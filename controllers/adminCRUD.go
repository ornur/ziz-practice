package controllers

import (
	"net/http"
	"os"
	"smallBotBackend/initializers"
	"smallBotBackend/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func CreateUserAdmin(c *gin.Context) {

	var createUserInput models.AuthInput

	if err := c.ShouldBindJSON(&createUserInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFound models.AdminUser
	initializers.DB.Where("username=?", createUserInput.Username).Find(&userFound)

	if userFound.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username already used"})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(createUserInput.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to hash password"})
		return
	}

	user := models.AdminUser{
		Username: createUserInput.Username,
		Password: string(hashedPassword),
	}

	initializers.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// Login godoc
// @Summary Login admin users
// @Description Sign in with username and password
// @Tags admin
// @Accept json
// @Produce json
// @Param input body AuthInput true "Auth Input"
// @Success 200 {string} token
// @Router /login [post]
func AdminLogin(c *gin.Context) {

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

func GetAdminUsers(c *gin.Context) {

	var users []models.AdminUser
	initializers.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})

}

func GetAdminUserByID(c *gin.Context) {
	
	id := c.Param("id")

	var user models.AdminUser
	initializers.DB.Where("id=?", id).Find(&user)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})

}

func DeleteAdminUserByID(c *gin.Context) {

	id := c.Param("id")

	var user models.AdminUser
	initializers.DB.Where("id=?", id).Find(&user)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	initializers.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})

}