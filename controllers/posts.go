package controllers

import (
	"database/sql"
	"go_crud/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func FindPosts(c *gin.Context) {
	posts := []map[string]interface{}{}
	if err := models.DB.NewSelect().Model(&models.Post{}).Column("title", "id").Scan(c, &posts); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"error": "Cannot fetch posts!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": posts})
}

func FindPost(c *gin.Context) {
	var post map[string]interface{}
	if err := models.DB.NewSelect().
		Model(&models.Post{}).
		Where("id = ?", c.Param("id")).
		Limit(1).
		Scan(c, &post); err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Cannot find post!"})
		return
	}

	var comments map[string]interface{}

	if err := models.DB.NewSelect().
		Model(&models.Comment{}).
		ExcludeColumn("post_id").
		Where("post_id = ?", c.Param("id")).
		Scan(c, &comments); err != nil && err != sql.ErrNoRows {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Cannot find post!"})
		return
	}
	if comments != nil {
		post["comments"] = comments
	} else {
		post["comments"] = []int{}
	}

	c.JSON(http.StatusOK, gin.H{"data": post})
}

type CreatePostInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func CreatePost(c *gin.Context) {
	var input CreatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := models.Post{Title: input.Title, Content: input.Content, CreatedAt: time.Now()}
	models.DB.NewInsert().Model(&post).Exec(c)

	c.JSON(http.StatusOK, gin.H{"msg": "Successfully posted"})
}
