package models

import "time"

type Todo struct {
	Id        uint      	`json:id gorm:primaryKey`
	Title     string    	`json:title`
	Note     *string    	`json:note`
	CreatedAt time.Time 	`json:created_at`
	UpdatedAt time.Time 	`json:updated_at`
}