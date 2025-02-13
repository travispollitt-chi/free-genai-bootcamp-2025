package models

import (
    "gorm.io/gorm"
)

type Word struct {
    gorm.Model
    Spanish string `json:"spanish"`
    English string `json:"english"`
    Parts   []string `json:"parts"`
}
