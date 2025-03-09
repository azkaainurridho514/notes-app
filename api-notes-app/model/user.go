package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)
type User struct {
	ID       		uuid.UUID 		`gorm:"type:char(36);primary_key" json:"id"`
	Username 		string    	`json:"username"`
	Phone    		string    	`json:"phone"`
	Email    		string    	`json:"email"`
	Password    	string    	`json:"password" gorm:"text"`
	Address    		string    	`json:"address" gorm:"type:text"`
	CreatedAt    	time.Time   `json:"created_at"`   
	UpdatedAt    	time.Time  	`json:"updated_at"`
}
type UserToFE struct {
	ID       		uuid.UUID 		`json:"id"`
	Username 		string    	`json:"username"`
	Phone    		string    	`json:"phone"`
	Email    		string    	`json:"email"`
	Address    		string    	`json:"address" gorm:"type:text"`
}
type UserLogin struct {
	Email 		string    	`json:"email"`
	Password    		string    	`json:"password"`
}
type UserFromFeRegister struct {
	Username 		string    	`json:"username"`
	Phone    		string    	`json:"phone"`
	Email    		string    	`json:"email"`
	Password    	string    	`json:"password" gorm:"text"`
	Address    		string    	`json:"address" gorm:"type:text"`
}
type UserFromFE struct {
	Username 		string    	`json:"username"`
	Phone    		string    	`json:"phone"`
	Email    		string    	`json:"email"`
	Password    	string    	`json:"password" gorm:"text"`
	Address    		string    	`json:"address" gorm:"type:text"`
}
type UserEditFromFE struct {
	ID       		uuid.UUID 		`json:"id"`
	Username 		string    	`json:"username"`
	Phone    		string    	`json:"phone"`
	Address    		string    	`json:"address" gorm:"type:text"`
}

type RegisterUser struct {
	ID       		uuid.UUID 		`gorm:"type:char(36);primary_key" json:"id"`
	Username 		string    	`json:"username"`
	Phone    		string    	`json:"phone"`
	Email    		string    	`json:"email"`
	Password    	string    	`json:"password" gorm:"text"`
	Address    		string    	`json:"address" gorm:"type:text"`
	CreatedAt    	time.Time   `json:"created_at"`   
	UpdatedAt    	time.Time  	`json:"updated_at"`
}
type DeleteUser struct {
	UserID 			string    	`json:"id"`
}
type Users struct {
 	Users []User `json:"users"`
}
func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	return
}