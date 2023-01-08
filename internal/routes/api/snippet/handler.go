package snippet

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/kuzuru/codesync/internal/entity/snippet"
	snippetEntity "github.com/kuzuru/codesync/internal/entity/snippet/db"
	"github.com/kuzuru/codesync/pkg/db/postgresql"
	"github.com/kuzuru/codesync/pkg/logger"
)

// resource is the structure responsible for representing
// the HTTP request unit for a given package
type resource struct {
	log     *logger.Logger
	app     fiber.Router
	storage snippet.Storage
}

// RegisterHandlers sets up the routing of the HTTP handlers
func RegisterHandlers(app fiber.Router) {
	log := logger.New()

	defer log.Sync()

	// TODO: Idk which context should I use rn
	psqlClient, err := postgresql.NewClient(context.TODO(), 3)
	if err != nil {
		log.Error("Got an error while trying connect to postgres",
			zap.Error(err),
		)
	}

	snippetRepo := snippetEntity.NewRepository(psqlClient)

	res := resource{
		log,
		app,
		snippetRepo,
	}

	group := app.Group("/snippet")

	group.Post("/get", res.get)
	group.Post("/new", res.new)
	group.Post("/latest", res.latest)
}

func (r *resource) get(c *fiber.Ctx) error {
	type Request struct {
		Link string `json:"link"`
	}

	var req Request

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Checking if link valid
	// Currently Link == UUID
	// TODO: Make links not UUIDs (It has to be changeable)
	_, err := uuid.Parse(req.Link)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid link",
		})
	}

	s, err := r.storage.GetByLink(context.TODO(), req.Link)
	if err != nil {
		// Doesn't logs out?
		r.log.Error("/snippet/get", zap.Error(err))

		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": s,
	})
}

func (r *resource) new(c *fiber.Ctx) error {
	// TODO: I should probably get rid of this in this file
	type Request struct {
		Title    string `json:"title"`
		Language string `json:"language"`
		Code     string `json:"code"`
		//Time           int    `json:"time"`            // 0 is lifetime
		//ViewsAvailable int    `json:"views_available"` // 0 is unlimited
	}

	var req Request

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var s snippet.Snippet

	// TODO: Replace with something better
	// TODO: I have to save Request struct in case to get what am I accepting
	s.Title = req.Title
	s.Language = req.Language
	s.Code = req.Code

	err := r.storage.Create(context.TODO(), &s)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"link": s.Link,
	})
}

func (r *resource) latest(c *fiber.Ctx) error {
	return c.SendString("latest\\handler!")
}
