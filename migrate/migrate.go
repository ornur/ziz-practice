package main

import (
	"smallBotBackend/initializers"
	"smallBotBackend/models"
)

func init() {
	initializers.LoadEnvs()
	initializers.ConnectDB()

}

func main() {

	initializers.DB.AutoMigrate(&models.AdminUser{}, &models.User{}, &models.Message{}, &models.MessageTranslation{}, &models.HelpMessage{}, &models.HelpMessageTranslation{}, &models.UserFeedback{})
}
