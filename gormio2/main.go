package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Channel struct {
	gorm.Model
	Name        string
	Description string
}

type User struct {
	gorm.Model
	Email    string
	Username string
}

type Message struct {
	gorm.Model
	Content   string
	UserID    uint
	ChannelID uint
	User      User
	Channel   Channel
}

func setup(db *gorm.DB) {
	db.AutoMigrate(&Channel{}, &User{}, &Message{})
	seed(db)
}

func seed(db *gorm.DB) {
	channels := []Channel{
		{Name: "General", Description: "General Discussions"},
		{Name: "Off-Topic", Description: "Weird stuff goes here"},
		{Name: "Suggestions", Description: "Video suggestions go here"},
	}
	for _, c := range channels {
		db.Create(&c)
	}
	users := []User{
		{Email: "test@test.com", Username: "Joe420"},
		{Email: "yes@yes.com", Username: "Bob"},
	}
	for _, u := range users {
		db.Create(&u)
	}
	var generalChat, suggestionsChat Channel
	db.First(&generalChat, "Name = ?", "General")
	db.First(&suggestionsChat, "Name = ?", "Suggestions")
	var joe, bob User
	db.First(&joe, "Username = ?", "Joe420")
	db.First(&bob, "Username = ?", "Bob")
	messages := []Message{
		{Content: "Hello!", Channel: generalChat, User: joe},
		{Content: "What up", Channel: generalChat, User: bob},
		{Content: "Make more go videos", Channel: suggestionsChat, User: joe},
	}
	for _, m := range messages {
		db.Create(&m)
	}
}

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("can't connect to database")
	}
	defer db.Close()
	db.LogMode(true)
	setup(db)
	var users []User
	db.Find(&users)
	for _, u := range users {
		fmt.Println("Email:", u.Email, "Username:", u.Username)
	}
	var messages []Message
	db.Model(users[0]).Related(&messages)
	for _, m := range messages {
		fmt.Println("Message:", m.Content, "Sender:", m.UserID)
	}
	doError(db)
}

func doError(db *gorm.DB) {
	var fred User
	if err := db.Where("username = ?", "Fred").First(&fred).Error; err != nil {
		log.Fatalf("Error when loading user: %s", err)
	}
}
