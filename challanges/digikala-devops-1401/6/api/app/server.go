package main

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func NewServer() error {
	app := fiber.New(fiber.Config{JSONEncoder: json.Marshal, JSONDecoder: json.Unmarshal, DisableStartupMessage: true})

	app.Get("/", isUp)
	app.Post("/signup", signup)
	app.Post("/login", login)
	app.Post("/suggestions", addSuggestions, authorize)
	app.Get("/suggestions", getSuggestions, authorize)

	return app.Listen("0.0.0.0:8080")
}

func isUp(c *fiber.Ctx) error {
	status := http.StatusOK
	response := map[string]bool{"ok": true}
	return c.Status(status).JSON(response)
}

func signup(c *fiber.Ctx) error {
	var user User
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	if len(user.Username) == 0 || len(user.Password) == 0 {
		status := http.StatusBadRequest
		response := map[string]any{"ok": false, "error": "no username or password provided"}
		return c.Status(status).JSON(response)
	}

	return nil
}

func login(c *fiber.Ctx) error {
	return nil
}

func authorize(c *fiber.Ctx) error {
	return nil
}

func addSuggestions(c *fiber.Ctx) error {
	return nil
}

func getSuggestions(c *fiber.Ctx) error {
	return nil
}
