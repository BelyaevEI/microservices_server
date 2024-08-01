package model

// Message represents a chat message
type Message struct {
	ID   int64
	Info MessageInfo
}

// MessageInfo represents a chat message info
type MessageInfo struct {
	ChatID int64
	UserID int64
	Text   string
}

// MessageCreate represents a chat message to be created
type MessageCreate struct {
	Info MessageInfo
}
