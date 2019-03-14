package main

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Article model
type Article struct {
	ID        int    `gorm:"primary_key" json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// GetArticleByID find by id
func GetArticleByID(db *gorm.DB, id int) (*Article, error) {
	var article Article
	if err := db.Where("id = ?", id).First(&article).Error; err != nil {
		return nil, err
	}

	return &article, nil
}

// GetArticles find all articles
func GetArticles(db *gorm.DB) (*[]Article, error) {
	var articles []Article
	if err := db.Find(&articles).Error; err != nil {
		return nil, err
	}

	return &articles, nil
}

// CreateArticle create article
func CreateArticle(db *gorm.DB, article *Article) (*Article, error) {
	if err := db.Create(&article).Error; err != nil {
		return nil, err
	}

	return article, nil
}

// UpdateArticle update article
func UpdateArticle(db *gorm.DB, article *Article) (*Article, error) {
	if err := db.Save(&article).Error; err != nil {
		return nil, err
	}

	return article, nil
}

// DeleteArticle delete article
func DeleteArticle(db *gorm.DB, article *Article) (*Article, error) {
	if err := db.Delete(&article).Error; err != nil {
		return nil, err
	}

	return article, nil
}
