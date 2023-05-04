package main

import (
	"cad/controllers"
	"cad/initializers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	// Add CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://chatakudong.vercel.app","http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type"},
	}))

	r.POST("/register", controllers.PostUser)
	r.POST("/:idUser/conversation", controllers.PostConversation)
	r.POST("/:idUser/:idConv/chat", controllers.PostChat)
	r.POST("/question", controllers.PostQuestion)
	r.GET("/login", controllers.GetUsers)
	r.GET("/:idUser", controllers.GetConversationFromUser)
	r.GET("/:idUser/:idConv", controllers.GetChatFromConversation)
	r.GET("/question", controllers.GetQuestions)
	r.DELETE("/:idUser/:idConv", controllers.DeleteConversation)
	r.DELETE("/question/:idQuestion", controllers.DeleteQuestion)
	r.Run()
}
