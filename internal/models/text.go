package models

type Text struct {
	Id       string `json:"-" bson:"id"`
	Type     string `json:"type" bson:"type" binding:"required"`
	Language string `json:"language" bson:"language" binding:"required"`
	Text     string `json:"text" bson:"text"`
}
