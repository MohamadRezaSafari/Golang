package main

import (
	"log"

	"github.com/gofiber/fiber"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type Post struct {
	Id          uint
	Title       string
	Description string
}

func main() {

	db, err := gorm.Open(sqlserver.Open("sa:V_f=^t%Fam!2GZ%@tcp(127.0.0.1:1433)/posts_ms"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(Post{})

	app := fiber.New()

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hi")
	// })

	log.Fatal(app.Listen(":5241"))
}
