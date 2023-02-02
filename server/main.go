package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Todo struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Done bool `json:"done"`
	Body string `json:"body"`
}

func main() {
	fmt.Print("Hello world")

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	todos := []Todo{}

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("ok") //returns ok on endpoint if status code is 200
	})

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{} //equal to the struct

		if err := c.BodyParser(todo); err != nil{
			return err
		}

		todo.ID = len(todos) + 1
		todos = append(todos, *todo) //adds the todo to the list of todos

		return c.JSON(todos)
	})

	app.Patch("/api/todos/:id/done", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id") //take out the id in endpoint as an integer

		if err != nil{
			return c.Status(401).SendString("invalid id")
		}

		for i, t := range todos { //i is index, t is todo
			if t.ID == id{
				todos[i].Done = true //sets the done attribute to true depending on the ID
				break
			}
		} 

		return c.JSON(todos)
	})

	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.JSON(todos)
	})

	log.Fatal(app.Listen(":4000")) //runs on port 4000
}
