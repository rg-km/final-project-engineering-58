package models

import "time"

type Post struct {
	ID             		string    `dbq:"id"`
	Title          		string    `dbq:"title"`
	Description    		string    `dbq:"description"`
	Content        		string    `dbq:"content"`
	UrlVideo       		string    `dbq:"url_video"`
	CategoryID     		string    `dbq:"category_id"`
	UserID         		string    `dbq:"user_id"`
	ParentID       		string    `dbq:"parent_id"`
	CreatedAt      		time.Time `dbq:"created_at"`
	UpdatedAt      		time.Time `dbq:"updated_at"`
}

func (Post) TableName() string {
	return `posts`
}

func TablePosts() []string {
	return []string{
		"id",
		"title",
		"description",
		"content",
		"url_video",
		"category_id",
		"user_id",
		"parent_id",
		"created_at",
		"updated_at",
	}
}
