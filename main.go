package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Short URL service started at :3000")
	app := fiber.New()
	Start(app)
}
