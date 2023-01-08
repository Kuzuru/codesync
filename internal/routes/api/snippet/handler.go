package snippet

import (
	"github.com/gofiber/fiber/v2"
)

// RegisterHandlers sets up the routing of the HTTP handlers
func RegisterHandlers(app fiber.Router) {
	res := resource{
		app,
	}

	group := app.Group("/snippet")

	group.Post("/get", res.get)
	group.Post("/new", res.new)
	group.Post("/latest", res.latest)
}

// resource is the structure responsible for representing
// the HTTP request unit for a given package
type resource struct {
	fiber.Router
}

func (r *resource) get(c *fiber.Ctx) error {
	return c.SendString("get\\handler!")
}

func (r *resource) new(c *fiber.Ctx) error {
	return c.SendString("new\\handler!")
}

func (r *resource) latest(c *fiber.Ctx) error {
	return c.SendString("latest\\handler!")
}
