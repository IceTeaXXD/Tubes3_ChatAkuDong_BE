package main

import (
	"cad/initializers"
	model "cad/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&model.Conversation{}, &model.Chat{}, &model.Question{}, &model.User{})
}
