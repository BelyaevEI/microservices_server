package model

type Chat struct {
	ID     int64
	Name   string
	UserID []int64
}

type ChatCreate struct {
	Name   string
	UserID []int64
}
