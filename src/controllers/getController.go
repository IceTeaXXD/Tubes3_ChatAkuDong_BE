package controllers

import (
	"cad/initializers"
	models "cad/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	initializers.DB.Find(&users)

	c.JSON(200, gin.H{
		"users": users,
	})
}

func GetConversationFromUser(c *gin.Context) {
	id := c.Param("idUser")

	var conversation []models.Conversation
	initializers.DB.Where("id_user = ?", id).Find(&conversation)

	c.JSON(200, gin.H{
		"conversation": conversation,
	})
}

func GetChatFromConversation(c *gin.Context) {
	id := c.Param("idConv")

	var chat []models.Chat
	initializers.DB.Where("id_conversation = ?", id).Find(&chat)

	c.JSON(200, gin.H{
		"chat": chat,
	})
}
func GetQuestions(c *gin.Context) {
	var questions []models.Question
	initializers.DB.Find(&questions)

	c.JSON(200, gin.H{
		"questions": questions,
	})
}
