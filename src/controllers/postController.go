package controllers

import (
	Algo "cad/algo"
	"cad/initializers"
	model "cad/models"
	// "fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func PostUser(c *gin.Context) {
	// get data
	var body struct {
		IDUser   int
		Username string
		Password string
	}
	c.Bind(&body)

	// create post
	post := model.User{IDUser: body.IDUser, Username: body.Username, Password: body.Password}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	// return post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostConversation(c *gin.Context) {
	// get data
	var body struct {
		IDConversation int
		Topic          string
		Date           time.Time
	}
	Date := time.Now();
	idUser, err := strconv.Atoi(c.Param("idUser"))
	if err != nil {
		// handle the error
		c.Status(400)
		return
	}
	c.Bind(&body)
	text := body.Topic
	post := model.Conversation{IDConversation: body.IDConversation, Topic: text, IDUser: idUser, Date: Date}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	// return post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostChat(c *gin.Context) {
	// get data
	var body struct {
		IDConversation int
		IDChat         int
		Question       string
		Answer         string
		SearchMethod   int
	}
	idUser, err1 := strconv.Atoi(c.Param("idUser"))
	if err1 != nil {
		c.Status(400)
		return
	}
	idConv, err2 := strconv.Atoi(c.Param("idConv"))
	if err2 != nil {
		c.Status(400)
		return
	}
	c.Bind(&body)

	// QnA
	var questions []model.Question
	var newQuestion model.Question
	var ret int

	initializers.DB.Find(&questions)
	body.Answer, ret = Algo.Regex(body.Question, questions, &newQuestion, body.SearchMethod)

	if ret == 2 {
		resCreate := initializers.DB.Create(&newQuestion)
		if resCreate.Error != nil {
			// c.Status(400)
			// return
			body.Answer = "Gagal menambahkan pertanyaan"
		}
		body.Answer = "Sukses menambahkan pertanyaan"
	} else if ret == 3 {
		// find the id of the question
		initializers.DB.Where("question = ?", newQuestion.Question).Find(&newQuestion)
		resDelete := initializers.DB.Delete(&newQuestion)
		if resDelete.Error != nil {
			// c.Status(400)
			// return
			body.Answer = "Gagal menghapus pertanyaan"
		}
		body.Answer = "Sukses menghapus pertanyaan"
	} else if ret == -1 {
		body.Answer = "Unknown Error"
	}

	post := model.Chat{Question: body.Question, Answer: body.Answer, IDUser: idUser, IDConversation: idConv, SearchMethod: body.SearchMethod}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	// return post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostQuestion(c *gin.Context) {
	// get data
	var body struct {
		IDQuestion int
		Question   string
		Answer     string
	}
	c.Bind(&body)
	// create post
	post := model.Question{IDQuestion: body.IDQuestion, Question: body.Question, Answer: body.Answer}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	// return post
	c.JSON(200, gin.H{
		"post": post,
	})
}
