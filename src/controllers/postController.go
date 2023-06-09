package controllers

import (
	Algo "cad/algo"
	"cad/initializers"
	model "cad/models"
	// "fmt"
	"regexp"
	"strconv"
	"strings"
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
	
	initializers.DB.Find(&questions)
	
	/* Cek apakah terdapat pemisah ";" */
	regex := regexp.MustCompile(`\b[^;]+\b`)
	return_bool, _ := regexp.MatchString(`;`, body.Question);
	if(return_bool){
		matches := regex.FindAllString(body.Question, -1)
		l := 0
		for l < len(matches){
			var newQuestion model.Question
			/* Pass in to Regex Again */
			ans, ret := Algo.Regex(matches[l], questions, &newQuestion, body.SearchMethod)
			
			if ret == 2 { // add new question
				// check if the question already exists, if it does, update the question answer to the new answer
				temp := model.Question{}
				initializers.DB.Where("question = ?", newQuestion.Question).Find(&temp)
				if temp.Question != "" {
					temp.Answer = newQuestion.Answer
					resUpdate := initializers.DB.Save(&temp)
					ans = "Sukses mengupdate pertanyaan"
					if resUpdate.Error != nil {
						ans = "Gagal mengupdate pertanyaan"
					}
					} else{
						// if it doesn't, create a new question
						resCreate := initializers.DB.Create(&newQuestion)
						ans = "Sukses menambahkan pertanyaan"
						if resCreate.Error != nil {
							ans = "Gagal menambahkan pertanyaan"
						}
					}
			} else if ret == 3 { // hapus question
				// find the id of the question
				initializers.DB.Where("question = ?", newQuestion.Question).Find(&newQuestion)
				resDelete := initializers.DB.Delete(&newQuestion)
				ans = "Sukses menghapus pertanyaan"
				if resDelete.Error != nil {
					ans = "Gagal menghapus pertanyaan"
					temp,_ := Algo.Regex(newQuestion.Question, questions, &newQuestion, body.SearchMethod)
					ans += "\n" + temp
				}
				} else if ret == -1 {
					ans = "Pertanyaan tidak ditemukan."
				}
			body.Answer += strconv.Itoa(l+1) + ". Q: " + matches[l] + "\nA: " + ans
			if l != len(matches)-1 {
				body.Answer += "\n\n"
			}
			l++
		}
	} else {
		ans, ret := Algo.Regex(body.Question, questions, &newQuestion, body.SearchMethod)

		if ret == 2 { // add new question
			// check if the question already exists, if it does, update the question answer to the new answer
			temp := model.Question{}
			initializers.DB.Where("question = ?", newQuestion.Question).Find(&temp)
			if temp.Question != "" {
				temp.Answer = newQuestion.Answer
				resUpdate := initializers.DB.Save(&temp)
				ans = "Sukses mengupdate pertanyaan"
				if resUpdate.Error != nil {
					ans = "Gagal mengupdate pertanyaan"
				}
			} else{
				// if it doesn't, create a new question
				resCreate := initializers.DB.Create(&newQuestion)
				ans = "Sukses menambahkan pertanyaan"
				if resCreate.Error != nil {
					ans = "Gagal menambahkan pertanyaan"
				}
			}
		} else if ret == 3 { // hapus question
			// find the id of the question
			initializers.DB.Where("question = ?", newQuestion.Question).Find(&newQuestion)
			resDelete := initializers.DB.Delete(&newQuestion)
			ans = "Sukses menghapus pertanyaan"
			if resDelete.Error != nil {
				ans = "Gagal menghapus pertanyaan"
				temp,_ := Algo.Regex(newQuestion.Question, questions, &newQuestion, body.SearchMethod)
				ans += "\n" + temp
			}
		} else if ret == -1 {
			ans = "Pertanyaan tidak ditemukan."
		}
		
		body.Answer = ans
	}
	
	// update current conversation topic to the first 3 words of the last question
	post := model.Chat{Question: body.Question, Answer: body.Answer, IDUser: idUser, IDConversation: idConv, SearchMethod: body.SearchMethod}
	result := initializers.DB.Create(&post)
	conv := model.Conversation{}
	initializers.DB.Where("id_conversation = ?", idConv).Find(&conv)
	words := strings.Fields(body.Question)
	if len(words) > 2 {
		conv.Topic = strings.Join(words[:3], " ")
	} else {
		conv.Topic = body.Question
	}
	initializers.DB.Save(&conv)

	if result.Error != nil {
		c.Status(400)
		return
	}

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
