package models

import "time"

type Book struct {
	Id              *int       `json:"id"`
	Title           *string    `json:"title"`
	Author          *string    `json:"author"`
	ISBN            *string    `json:"isbn"`
	PublicationDate *string    `json:"publication_date"`
	PageCount       *int       `json:"page_count"`
	Genre           *string    `json:"genre"`
	Description     *string    `json:"description"`
	CreatedAt       *time.Time `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
}
