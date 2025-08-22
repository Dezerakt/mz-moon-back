package middlewares

import (
	"mz-moon-back/config"
	"net/http"

	"github.com/gofiber/fiber"
)

func Cors(c *fiber.Ctx) {
	origin := c.Get("Origin")
	if origin != "" {
		for _, v := range config.GetTrustedOrigins() {
			if origin == v {
				c.Set("Access-Control-Allow-Origin", v)
			}
		}
	}

	c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Set("Access-Control-Allow-Credentials", "true")
	c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-CSRF-Token, Lng")
	c.Set("Access-Control-Expose-Headers", "Content-Disposition, X-File-Name")
	c.Set("Content-type", "application/json")

	if c.Method() != "OPTIONS" {
		c.Next()
	} else {
		c.Status(http.StatusOK)
	}
}
