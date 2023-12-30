package database

import (
	. "github.com/devtosun/datamorph-go/model/auth"
	"log"
)

func Create(entity interface{}) {
	log.Println("-------DB")
	log.Println(DB)
	log.Println("-------DB")
	if err := DB.Create(entity).Error; err != nil {
		log.Fatalf("Kullanıcı kaydedilirken hata oluştu: %v", err)
	}

	log.Println("Kullanıcı başarıyla kaydedildi. ")
}

func AutoMigration() {
	
	DB.AutoMigrate(&User{}, &Role{})

	// DB.AutoMigrate(&Role{})
	// DB.AutoMigrate(&Permission{})

	// Migrator oluştur
	migrator := DB.Migrator()
	

	// DropTable ile mevcut tabloları sil
	if err := migrator.DropTable(&User{}, &Role{}); err != nil {
		log.Fatalf("DropTable -> %v", err)
	}

	//AutoMigrate ile yeni tabloları oluştur
	if err := DB.AutoMigrate(&User{}, &Role{}); err != nil {
		log.Fatalf("AutoMigrate -> %v", err)
	}
}
