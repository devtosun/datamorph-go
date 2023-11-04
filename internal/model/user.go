package model

import (
    // "github.com/google/uuid"
    "gorm.io/gorm"
)

// type User struct {
//     gorm.Model           // Adds some metadata fields to the table
//     ID         		uuid.UUID `gorm:"type:uuid"` // Explicitly specify the type to be uuid
//     UserName      	string
//     Password   		string
// }


type User struct {
	gorm.Model
    ID           uint   `gorm:"primaryKey"`
    Username     string `gorm:"not null"`
    Email        string `gorm:"not null"`
    PasswordHash string `gorm:"not null"`
    Roles        []Role `gorm:"many2many:user_roles"`
    Permissions  []Permission `gorm:"many2many:user_permissions"`
}

type Role struct {
	gorm.Model
    ID          uint   `gorm:"primaryKey"`
    Name        string `gorm:"not null"`
    Permissions []Permission `gorm:"many2many:role_permissions"`
    Users       []User `gorm:"many2many:user_roles"`
}

type Permission struct {
	gorm.Model
    ID    uint   `gorm:"primaryKey"`
    Name  string `gorm:"not null"`
    Roles []Role `gorm:"many2many:role_permissions"`
    Users []User `gorm:"many2many:user_permissions"`
}
