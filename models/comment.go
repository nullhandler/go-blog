package models

import (
	"time"
)

type Comment struct{
	ID    uint   `json:"id" bun:"id,pk,autoincrement"`
	CreatedAt time.Time `json:"created_at" bun:"created_at"`
	Content string 	`json:"content" bun:"content" binding:"required"`
	PostID uint		`json:"post_id" binding:"required"`
}