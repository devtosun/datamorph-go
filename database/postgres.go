package database

import (
	"log"
	"github.com/devtosun/datamorph-go/model"
)

func CreateUser(newUser User) {
	if err := DB.Create(newUser).Error; err != nil {
        log.Fatalf("Kullanıcı kaydedilirken hata oluştu: %v", err)
    }

    log.Println("Kullanıcı başarıyla kaydedildi.")
}

