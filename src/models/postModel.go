package model

type User struct {
	IDUser   int    `gorm:"primaryKey"`
	Username string `gorm:"not null"`
	Password string `gorm:"not null"`
}

type Conversation struct {
	IDConversation int    `gorm:"primaryKey"`
	IDUser         int    `gorm:"not null"`
	Topic          string `gorm:"not null"`
	Chats          []Chat `gorm:"foreignKey:IDConversation"`
}

type Chat struct {
	IDConversation int    `gorm:"primaryKey"`
	IDChat         int    `gorm:"primaryKey"`
	IDUser         int    `gorm:"not null"`	
	Question       string `gorm:"not null"`
	Answer         string `gorm:"not null"`
}

type Question struct {
	IDQuestion int    `gorm:"primaryKey"`
	Question   string `gorm:"not null"`
	Answer     string `gorm:"not null"`
}
