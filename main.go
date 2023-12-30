package main

import (
	"log"

	data "github.com/devtosun/datamorph-go/database"
	model "github.com/devtosun/datamorph-go/model/auth"
	router "github.com/devtosun/datamorph-go/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	app := fiber.New()

	// data.Hello()

	app.Get("/", func(c *fiber.Ctx) error {

		return c.SendString("Datamorph  5👋!")
	})

	data.ConnectDB()
	data.AutoMigration()
	creteNewUser()

	app.Use(cors.New(cors.Config{
        AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
        AllowOrigins:     "*",
        AllowCredentials: true,
        AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
    }))

	router.SetupRoutes(app)

	app.Listen(":3000")
}

func creteNewUser() {

	// Örnek bir rol oluşturun
	adminRole := &model.Role{
		Name: "admin",
		ID: 1,
	}

	// userRole := &model.Role{
	// 	Name: "user",
	// }

	// roles := []model.Role{*adminRole, *userRole}
    pass := "admin"
	hash, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	newUser := &model.User{
		Username: "admin",
		Email:    "admin@datamorph.com",
		Password: string(hash),
		Role: *adminRole,
		// Permissions: permissions,
	}



	// Rolü ve yetkiyi veritabanına kaydedin
	if err := data.DB.Create(adminRole).Error; err != nil {
		log.Fatalf("Rol kaydedilirken hata oluştu: %v", err)
	}
	println(adminRole.ID)

	// if err := data.DB.Create(userRole).Error; err != nil {
	// 	log.Fatalf("Rol kaydedilirken hata oluştu: %v", err)
	// }

	// Kullanıcıyı veritabanına kaydedin
	if err := data.DB.Create(newUser).Error; err != nil {
	    log.Fatalf("Kullanıcı kaydedilirken hata oluştu: %v", err)
	}

	// Kullanıcıya rol ve yetki atayın

	// if err := data.DB.Create(newUser).Error; err != nil {
	//     log.Fatalf("Kullanıcıya rol ve yetki atama hatası: %v", err)
	// }

	log.Println("Kullanıcı, rol ve yetki başarıyla kaydedildi ve ilişkilendirildi........")
}
