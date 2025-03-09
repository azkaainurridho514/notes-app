package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)
type Note struct {
	ID       		uuid.UUID 	`gorm:"type:char(36);primary_key" json:"id"`
	UserID      	string 		`gorm:"type:char(36)" json:"user_id"`
	Title 			string    	`json:"title"`
	Description    	string    	`json:"description" gorm:"type:text"`
	CreatedAt    	time.Time   `json:"created_at"`   
	UpdatedAt    	time.Time  	`json:"updated_at"`
}
type NoteFromFE struct {
	Title 			string    	`json:"title"`
	Description    	string    	`json:"description" gorm:"type:text"`
}
type NoteToFE struct {
	ID       		uuid.UUID 	`json:"id"`
	Title 			string    	`json:"title"`
	Description    	string    	`json:"description" gorm:"type:text"`
}
type NoteEditFromFE struct {
	ID       		uuid.UUID 	`json:"id"`
	Title 			string    	`json:"title"`
	Description    	string    	`json:"description" gorm:"type:text"`
}

type CreateNote struct {
	ID       		uuid.UUID 	`gorm:"type:char(36);primary_key" json:"id"`
	Title 			string    	`json:"title"`
	Description    	string    	`json:"description" gorm:"type:text"`
	CreatedAt    	time.Time   `json:"created_at"`   
	UpdatedAt    	time.Time  	`json:"updated_at"`
}
type DeleteNote struct {
	NoteID 			string    	`json:"id"`
}
type Notes struct {
	Notes []Note `json:"notes"`
}
func (note *Note) BeforeCreate(tx *gorm.DB) (err error) {
	note.ID = uuid.New()
	return
}