package repository

import (
	"Test-Task/TextCDN/internal/models"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

const (
	type_of_text = "type"
	lang         = "language"
)

type TextMongo struct {
	collection *mongo.Collection
}

func NewTextMongo(collection *mongo.Collection) *TextMongo {
	return &TextMongo{collection: collection}
}

func (c *TextMongo) GetText(typeOfText, language *string) (models.Text, error) {
	var text models.Text
	err := c.collection.FindOne(
		context.TODO(),
		bson.D{{type_of_text, typeOfText}, {lang, language}},
	).Decode(&text)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Text{}, errors.New("Text doesnt exists")
		}
		return models.Text{}, err
	}
	return text, nil
}
