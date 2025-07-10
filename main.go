package main

import (
	handlers "api_study/handlers"
	config "api_study/init"
	jwt_pkg "api_study/pkg/jwt"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func ProtectedHandler(c *fiber.Ctx) error {
	ignoredEndpoints := map[string][]string{"/api/login": {"POST"}, "/api/users": {"POST"}}
	endpointPath := string(c.Request().RequestURI())
	requestMethod := string(c.Request().Header.Method())

	for route, methods := range ignoredEndpoints {
		if endpointPath == route {
			for _, method := range methods {
				if requestMethod == method {
					return c.Next()
				}
			}
		}
	}

	tokenString := string(c.Request().Header.Peek("Authorization"))

	if tokenString == "" {
		return c.Status(401).JSON(map[string]string{"error": "Missing authorization header"})
	}

	tokenString = tokenString[len("Bearer "):]

	if tokenString == "" {
		return c.Status(401).JSON(map[string]string{"error": "Invalid token"})
	}

	err := jwt_pkg.VerifyToken(tokenString)
	if err != nil {
		return c.Status(401).JSON(map[string]string{"error": "Invalid token"})
	}

	return c.Next()
}

func main() {
	config.SetConfig()
	config.InitApp()
	config.InitDB()

	config.App.Use(cors.New())
	config.App.Use(ProtectedHandler)

	config.App.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, world!")
	})

	config.App.Post("/api/login", handlers.Login)
	config.App.Post("/api/users", handlers.CreateUser)
	config.App.Get("/api/users/:user_id", handlers.RetriveUser)
	config.App.Put("/api/users/:user_id", handlers.UpdateUser)
	config.App.Get("/api/users", handlers.ListUsers)
	config.App.Delete("/api/users/:user_id", handlers.DeleteUser)

	config.App.Get("/uploads/:category/:category_id/:filename", func(c *fiber.Ctx) error {
		category := c.Params("category")
		categoryId := c.Params("category_id")
		filename := c.Params("filename")

		return c.SendFile(fmt.Sprintf("uploads/%s/%s/%s", category, categoryId, filename))
	})

	config.App.Listen(":" + strconv.Itoa(config.GlobaConfig.Port))
}
