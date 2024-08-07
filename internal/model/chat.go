package model

// Chat represents a chat
type Chat struct {
	ID     int64
	Name   string
	UserID []int64
}

// ChatCreate represents a chat to be created
type ChatCreate struct {
	Name   string
	UserID []int64
}
