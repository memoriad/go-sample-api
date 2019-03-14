package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getArticle(c *gin.Context) {
	if id, err := strconv.Atoi(c.Param("article_id")); err == nil {
		db := InitDb()

		if article, err := GetArticleByID(db, id); err == nil {
			c.JSON(200, article)
		} else {
			c.AbortWithStatus(http.StatusNotFound)
		}

		defer func() {
			db.Close()
		}()

	} else {
		// If an invalid article ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func getArticles(c *gin.Context) {
	db := InitDb()

	if articles, err := GetArticles(db); err == nil {
		c.JSON(200, articles)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}

	defer func() {
		db.Close()
	}()
}

func createArticle(c *gin.Context) {
	// Obtain the POSTed title and content values
	var article Article
	c.BindJSON(&article)

	db := InitDb()
	tx := db.Begin()

	if article, err := CreateArticle(tx, &article); err == nil {
		tx.Commit()
		c.JSON(200, article)
	} else {
		tx.Rollback()
		c.AbortWithStatus(http.StatusBadRequest)
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
		tx.Close()
		db.Close()
	}()
}

func updateArticle(c *gin.Context) {
	if id, err := strconv.Atoi(c.Param("article_id")); err == nil {
		db := InitDb()
		tx := db.Begin()

		if article, err := GetArticleByID(tx, id); err == nil {
      var json Article
      c.BindJSON(&json)
      article.Title = json.Title
      article.Content = json.Content
      
			if article, err := UpdateArticle(tx, article); err == nil {
				tx.Commit()
				c.JSON(200, article)
			} else {
				tx.Rollback()
				c.AbortWithStatus(http.StatusBadRequest)
			}
		} else {
			c.AbortWithStatus(http.StatusNotFound)
		}

		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
			}
			tx.Close()
			db.Close()
		}()
	} else {
		// If an invalid article ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func deleteArticle(c *gin.Context) {
	if id, err := strconv.Atoi(c.Param("article_id")); err == nil {
		db := InitDb()
		tx := db.Begin()

		if article, err := GetArticleByID(tx, id); err == nil {
			if article, err := DeleteArticle(tx, article); err == nil {
				tx.Commit()
				c.JSON(200, article)
			} else {
				tx.Rollback()
				c.AbortWithStatus(http.StatusBadRequest)
			}
		} else {
			c.AbortWithStatus(http.StatusNotFound)
		}

		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
			}
			tx.Close()
			db.Close()
		}()
	} else {
		// If an invalid article ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}
