package services

import (
    "lang-portal/models"
    "gorm.io/gorm"
)

type WordService struct {
    DB *gorm.DB
}

func (s *WordService) GetAllWords() ([]models.Word, error) {
    var words []models.Word
    result := s.DB.Find(&words)
    return words, result.Error
}

func (s *WordService) CreateWord(word *models.Word) error {
    result := s.DB.Create(word)
    return result.Error
}
