package services

import (
	"Test-Task/TextCDN/internal/models"
	"Test-Task/TextCDN/internal/repository"
)

type Service struct {
	TextViewer
}

type TextViewer interface {
	GetText(typeOfText string, language string) (models.Text, error)
}

func NewService(r *repository.Repository) *Service {
	return &Service{NewTextService(r)}
}
