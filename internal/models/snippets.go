package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Snippet struct {
	gorm.Model
	Title     string    `gorm:"size:100;not null"`
	Content   string    `gorm:"type:text;not null"`
	Expires   time.Time `gorm:"not null;index"`
	IsExpired bool      `gorm:"default:false"`
}

type SnippetModel struct {
	DB *gorm.DB
}

func NewSnippetModel(db *gorm.DB) *SnippetModel {
	return &SnippetModel{DB: db}
}

func (m *SnippetModel) Insert(snippet *Snippet) error {
	return m.DB.Create(snippet).Error
}

func (m *SnippetModel) Get(id uint) (*Snippet, error) {
	var snippet Snippet
	err := m.DB.Where("expires > ? AND deleted_at IS NULL", time.Now()).First(&snippet, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("snippet not found")
		}
		return nil, err
	}
	return &snippet, nil
}

func (m *SnippetModel) Latest() ([]Snippet, error) {
	var snippets []Snippet
	err := m.DB.Where("expires > ? AND deleted_at IS NULL", time.Now()).
		Order("created_at desc").
		Limit(10).
		Find(&snippets).Error
	return snippets, err
}

func (m *SnippetModel) Delete(id uint) error {
	return m.DB.Delete(&Snippet{}, id).Error
}

func (m *SnippetModel) Update(snippet *Snippet) error {
	return m.DB.Save(snippet).Error
}
