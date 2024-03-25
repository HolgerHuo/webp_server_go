package middleware

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Headers() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		err := c.Next()

		diff := time.Since(start)
		c.Append("X-Processing-Time", fmt.Sprintf("%dms", diff.Milliseconds()))

		c.Append("X-Powered-By", "DragonCloud Images (DEU)")
		return err
	}
}