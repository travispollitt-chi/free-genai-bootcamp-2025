package models

import (
	"database/sql"
	"gorm.io/gorm"
	// ...other imports...
)

// Define models for each table
type Word struct {
	gorm.Model
	Spanish string
	English string
	Parts   string
}

type WordGroup struct {
	ID      int
	WordID  int
	GroupID int
}

type Group struct {
	gorm.Model
	Name       string
	WordsCount int
}

type StudySession struct {
	ID             int
	GroupID        int
	StudyActivityID int
	CreatedAt      string
}

type StudyActivity struct {
	ID   int
	Name string
	URL  string
}

type WordReviewItem struct {
	ID            int
	WordID        int
	StudySessionID int
	Correct       bool
	CreatedAt     string
}

// ...other models...
