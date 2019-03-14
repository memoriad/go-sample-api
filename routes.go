package main

import (
	"github.com/gin-contrib/cors"
)

func initializeRoutes() {

  config := cors.DefaultConfig()
  config.AllowAllOrigins = true
  config.AddAllowMethods("PATCH", "DELETE")
  router.Use(cors.New(config))
  
  // Handle GET requests at /articles
  router.GET("/articles", getArticles)

  articleRoutes := router.Group("/article")
	{
    // Handle GET requests at /article/article_id
    articleRoutes.GET("/:article_id", getArticle)
    
		// Handle POST requests at /article/add
    articleRoutes.POST("/add", createArticle)
    
    // Handle PATCH requests at /article/update/article_id
    articleRoutes.PATCH("/update/:article_id", updateArticle)
    
    // Handle DELETE requests at /article/update/article_id
		articleRoutes.DELETE("/delete/:article_id", deleteArticle)
	}

}