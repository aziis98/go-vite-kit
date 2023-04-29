package routes

import "github.com/gofiber/fiber/v2"

func (r *Router) Api(api fiber.Router) {
	api.Get("/status", func(c *fiber.Ctx) error {
		return c.JSON("ok")
	})
}
