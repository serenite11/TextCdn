package repository

import (
	"Test-Task/TextCDN/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	TextViewer
}

type TextViewer interface {
	GetText(typeOfText string, language string) (models.Text, error)
}

func NewRepository(collection *mongo.Collection) *Repository {
	return &Repository{NewTextMongo(collection)}
}
