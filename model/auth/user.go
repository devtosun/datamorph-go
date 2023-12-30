package model

import (
	// "github.com/google/uuid"
	"gorm.io/gorm"
)

type LoginInput struct {
	Identity string `json:"identity"`
	Password string `json:"password"`
	gorm.Model 
}

type User struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	RoleID   uint   `gorm:"not null;DEFAULT:3" json:"role_id"`
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Email    string `gorm:"size:255;not null;unique" json:"email"`
	Password string `gorm:"size:255;not null" json:"-"`
	Role     Role   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	gorm.Model
}


// type User struct {
//     gorm.Model           // Adds some metadata fields to the table
//     ID         		uuid.UUID `gorm:"type:uuid"` // Explicitly specify the type to be uuid
//     UserName      	string
//     Password   		string
// }



// type User struct {
// 	ID           uint         `gorm:"primary_key"`
// 	Username     string       `gorm:"not null"`
// 	Email        string       `gorm:"not null"`
// 	PasswordHash string       `gorm:"not null"`
// 	Roles        []Role       `gorm:"many2many:user_roles"`
// 	Permissions  []Permission `gorm:"many2many:user_permissions"`
// 	gorm.Model
// }

// type Role struct {
// 	ID          uint         `gorm:"primary_key"`
// 	Name        string       `gorm:"not null;unique"`
// 	Permissions []Permission `gorm:"many2many:role_permissions"`
// 	Users       []User       `gorm:"many2many:user_roles"`
// 	gorm.Model
// }

// type Permission struct {
// 	ID uint `gorm:"primary_key"`
// 	// ID    uint   `gorm:"primaryKey"`
// 	Name  string `gorm:"not null;unique"`
// 	Roles []Role `gorm:"many2many:role_permissions"`
// 	Users []User `gorm:"many2many:user_permissions"`
// 	gorm.Model
// }
