package models

import "time"

type User struct {
	ID             		string    `dbq:"id"`
	Name           		string    `dbq:"name"`
	Email          		string    `dbq:"email"`
	Role          		string    `dbq:"role"`
	Password       		string    `dbq:"password"`
	CreatedAt      		time.Time `dbq:"created_at"`
	UpdatedAt      		time.Time `dbq:"updated_at"`
}

func (User) TableName() string {
	return `users`
}

func TableUsers() []string {
	return []string{
		"id",
		"name",
		"email",
		"role",
		"password",
		"created_at",
		"updated_at",
	}
}
