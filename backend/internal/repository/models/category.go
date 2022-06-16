package models

import "time"

type Category struct {
	ID             		string    `dbq:"id"`
	Name           		string    `dbq:"name"`
	CreatedAt      		time.Time `dbq:"created_at"`
	UpdatedAt      		time.Time `dbq:"updated_at"`
}

func (Category) TableName() string {
	return `categories`
}

func TableCategories() []string {
	return []string{
		"id",
		"name",
		"created_at",
		"updated_at",
	}
}
