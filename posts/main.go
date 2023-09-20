package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Post struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func main() {

	dsn := "host=localhost user=postgres password=1 dbname=posts_ms port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(Post{})

	app := fiber.New()
	app.Use(cors.New())

	app.Get("/api/posts", func(c *fiber.Ctx) error {

		var posts []Post

		db.Find(&posts)

		return c.JSON(posts)
	})

	app.Post("/api/posts/create", func(c *fiber.Ctx) error {

		var post Post

		if err := c.BodyParser(&post); err != nil {
			return err
		}

		db.Create(&post)

		return c.JSON(post)

	})

	log.Fatal(app.Listen(":5241"))
}
