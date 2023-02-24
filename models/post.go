package models

import "time"

type Post struct {
	ID    uint   `json:"id" bun:"id,pk,autoincrement"`
	Title string `json:"title" bun:"title"`
	Content   string    `json:"content" bun:"content"`
	CreatedAt time.Time `json:"created_at" bun:"created_at"`
	Comments []*Comment  `json:"comments" bun:"rel:has-many,join:id=post_id"`
}
