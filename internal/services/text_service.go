package services

import (
	"Test-Task/TextCDN/internal/models"
	"Test-Task/TextCDN/internal/repository"
)

type TextService struct {
	repo repository.TextViewer
}

func NewTextService(repos repository.TextViewer) *TextService {
	return &TextService{repo: repos}
}

func (s *TextService) GetText(typeOfText string, language string) (models.Text, error) {
	return s.repo.GetText(typeOfText, language)
}
