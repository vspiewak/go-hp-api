package models

type Message struct {
	Login     string `json:"login" binding:"alpha"`
	Tweet 		string `json:"tweet" binding:"required"`
	Comment   string `json:"comment" binding:"omitempty,alphanum"`
}
