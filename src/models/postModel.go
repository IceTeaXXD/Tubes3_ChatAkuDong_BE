package model

import "time"

type User struct {
	IDUser   int    `gorm:"primaryKey"`
	Username string `gorm:"not null"`
	Password string `gorm:"not null"`
}

type Conversation struct {
	IDConversation int       `gorm:"primaryKey"`
	IDUser         int       `gorm:"not null"`
	Topic          string    `gorm:"not null"`
	Chats          []Chat    `gorm:"foreignKey:IDConversation;constraint:OnDelete:CASCADE"`
	Date           time.Time `gorm:"not null"`
}

type Chat struct {
	IDChat         int    `gorm:"type:serial;primaryKey;not null"`
	IDConversation int    `gorm:"not null"`
	IDUser         int    `gorm:"not null"`
	Question       string `gorm:"not null"`
	Answer         string `gorm:"not null"`
}

type Question struct {
	IDQuestion int    `gorm:"primaryKey"`
	Question   string `gorm:"not null"`
	Answer     string `gorm:"not null"`
}
