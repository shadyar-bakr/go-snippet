package models

import (
	"time"

	"gorm.io/gorm"
)

type Snippet struct {
	gorm.Model
	Title   string    `gorm:"size:100;not null"`
	Content string    `gorm:"type:text;not null"`
	Expires time.Time `gorm:"not null;index"`
}

func CreateSnippet(db *gorm.DB, snippet *Snippet) error {
	if snippet.Title == "" || snippet.Content == "" {
		return ErrInvalidInput
	}

	result := db.Create(snippet)
	if result.Error != nil {
		return ErrInternalServer
	}
	return nil
}

func GetSnippet(db *gorm.DB, id uint) (*Snippet, error) {
	var snippet Snippet
	result := db.First(&snippet, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, ErrNoRecord
		}
		return nil, ErrInternalServer
	}

	if time.Now().After(snippet.Expires) {
		return nil, ErrUnauthorized
	}

	return &snippet, nil
}

func GetAllSnippets(db *gorm.DB) ([]Snippet, error) {
	var snippets []Snippet
	result := db.Where("expires > ?", time.Now()).
		Order("created_at DESC").
		Find(&snippets)

	if result.Error != nil {
		return nil, ErrInternalServer
	}

	return snippets, nil
}

func UpdateSnippet(db *gorm.DB, snippet *Snippet) error {
	if snippet.Title == "" || snippet.Content == "" {
		return ErrInvalidInput
	}

	result := db.Save(snippet)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return ErrNoRecord
		}
		return ErrInternalServer
	}

	return nil
}

func DeleteSnippet(db *gorm.DB, id uint) error {
	result := db.Delete(&Snippet{}, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return ErrNoRecord
		}
		return ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return ErrNoRecord
	}

	return nil
}
