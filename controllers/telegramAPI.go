package controllers

import(
	"net/http"
	"smallBotBackend/initializers"
	"smallBotBackend/models"
	"github.com/gin-gonic/gin"
)

// check user by telegram id from users table if not existing user send sorry message
// if user exists send welcome message by role and ask for language
// all messages by language and role of user from messages table => message_translations'

func CheckUser(c *gin.Context) {
	
	telegramID := c.Param("telegram_id")
	
	var user models.User
	initializers.DB.Where("telegram_id=?", telegramID).Find(&user)
	
	if user.ID == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Sorry, you are not registered in our system"})
		return
	}
	
	// if user exits take roleID and language from `users` table.
	// then roleID goes to `messages` table and get message_id by role
	// then language and message_id goes to `message_translations` table and get `content`
	// then send `content` saves in `message` variable

	var message string
	var messageTranslation models.MessageTranslation
	initializers.DB.Where("roleID=?", user.RoleID).Find(&messageTranslation)
	initializers.DB.Where("language=? AND message_id=?", user.Language, messageTranslation.MessageID).Find(&messageTranslation)
	message = messageTranslation.Content

	
	c.JSON(http.StatusOK, gin.H{"message": message})
	
}