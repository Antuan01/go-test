package main

import (
    "log"
    "github.com/Antuan01/go-test/routes"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/Antuan01/go-test/database"
)

func main() {
    app := fiber.New()

    app.Use(cors.New())

    app.Use(logger.New())

    database.InitDB()

    routes.SetupRoutes(app)

    log.Fatal(app.Listen(":3000"))
}