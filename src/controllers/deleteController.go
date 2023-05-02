package controllers

import (
	"cad/initializers"
	model "cad/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteConversation(c *gin.Context) {
	idUser, err := strconv.Atoi(c.Param("idUser"))
	if err != nil {
		// handle the error
		c.Status(400)
		return
	}
	idConv, err := strconv.Atoi(c.Param("idConv"))
	if err != nil {
		// handle the error
		c.Status(400)
		return
	}
	initializers.DB.Delete(&model.Conversation{}, idUser, idConv)
}

func DeleteQuestion(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("idQuestion"))
	if err != nil {
		// handle the error
		c.Status(400)
		return
	}
	initializers.DB.Delete(&model.Question{}, id)
}
