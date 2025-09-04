package middlewares

import (
	"log"

	"github.com/gofiber/fiber"
	"github.com/google/uuid"
)

func LogRequest(c *fiber.Ctx) {
	traceId := uuid.New()
	c.Set("X-Trace-Id", traceId.String())
	log.Println("Request: ", c)

	c.Next()

	log.Println("Response: ", c)
}
