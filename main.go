package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	fmt.Println("Hello World main")

	app := fiber.New()
	todos := []Todo{}


	app.Get("/api/todos/:id", func (c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
		}

		for _, todo := range todos {
			if todo.ID == id {
				return c.Status(200).JSON(todo)
			}
		}

		// If no todo is found with the given ID
		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
	})

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}

		if err := c.BodyParser(todo); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Todo must contain a body"})
		}

		todo.ID = len(todos)
		todos = append(todos, *todo)

		return c.Status(201).JSON(todo)
	})

	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
		}

		for i, todo := range todos {
			if todo.ID == id {
				todos[i].Completed = true
				return c.Status(200).JSON(todos[i])
			}
		}

		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
	})

	app.Delete("/api/todos/:id", func (c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))

		if err != nil {
			return c.Status(400).JSON(fiber.Map{"msg": "Invalid ID"})
		}

		for i, todo := range todos {
			if todo.ID == id {
				todos = append(todos[:i], todos[i+1:]...)
				return c.Status(200).JSON(fiber.Map{"msg": "Todo successfully removed"})
			}
		}
		return c.Status(404).JSON(fiber.Map{"msg": "Todo not found"})
		
	})

	app.Get("/api/todos", func(c *fiber.Ctx) error {
		fmt.Println("GET /api/todos called")
		return c.Status(200).JSON(todos)
	})

	log.Fatal(app.Listen(":9090"))
}
