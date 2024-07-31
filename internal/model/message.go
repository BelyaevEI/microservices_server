package model

type Message struct {
	ID   int64
	Info MessageInfo
}

type MessageInfo struct {
	ChatID int64
	UserID int64
	Text   string
}

type MessageCreate struct {
	Info MessageInfo
}
