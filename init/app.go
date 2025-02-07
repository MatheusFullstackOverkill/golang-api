package init

import "github.com/gofiber/fiber/v2"

var App *fiber.App

func InitApp() {
	App = fiber.New()

	App.Static("/uploads", "uploads")
}
