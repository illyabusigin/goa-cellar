package controllers

import (
	"fmt"
	"io"
	"time"

	"github.com/goadesign/goa"
	"github.com/illyabusigin/goa-cellar/app"
	"github.com/illyabusigin/goa-cellar/models"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/websocket"
)

// BottleController implements the bottle resource.
type BottleController struct {
	*goa.Controller
	storage models.BottleStorage
}

// NewBottle creates a bottle controller.
func NewBottle(service *goa.Service, db *gorm.DB) *BottleController {
	return &BottleController{
		Controller: service.NewController("Bottle"),
	}
}

// List lists all the bottles in the account optionally filtering by year.
func (c *BottleController) List(ctx *app.ListBottleContext) error {
	var bottles []*app.Bottle
	var err error

	if ctx.Years != nil {
		defer goa.MeasureSince([]string{"goa", "db", "bottle", "listbottlebyyears"}, time.Now())

		var native []*models.Bottle
		db := c.storage.DB().(*gorm.DB)
		err = db.Scopes(models.BottleFilterByAccount(ctx.AccountID, db)).Table("bottles").Preload("Account").Find(&native).Where("year in (?)", ctx.Years).Error

		if err != nil {
			goa.LogError(ctx, "error listing Bottle", "error", err.Error())
			return ErrDatabaseError(err)
		}

		for _, t := range native {
			bottles = append(bottles, t.BottleToBottle())
		}
	} else {
		bottles = c.storage.ListBottle(ctx, ctx.AccountID)
	}

	if err != nil {
		return ErrDatabaseError(err)
	}

	return ctx.OK(bottles)
}

// Show retrieves the bottle with the given id.
func (c *BottleController) Show(ctx *app.ShowBottleContext) error {
	bottle, err := c.storage.OneBottle(ctx, ctx.BottleID, ctx.AccountID)
	if err != nil {
		return ErrDatabaseError(err)
	}

	return ctx.OK(bottle)
}

// Watch watches the bottle with the given id.
func (c *BottleController) Watch(ctx *app.WatchBottleContext) error {
	Watcher(ctx.AccountID, ctx.BottleID).ServeHTTP(ctx.ResponseWriter, ctx.Request)
	return nil
}

// Watcher will echo the data received on the WebSocket.
func Watcher(accountID, bottleID int) websocket.Handler {
	return func(ws *websocket.Conn) {
		watched := fmt.Sprintf("Account: %d, Bottle: %d", accountID, bottleID)
		ws.Write([]byte(watched))
		io.Copy(ws, ws)
	}
}

// Create records a new bottle.
func (c *BottleController) Create(ctx *app.CreateBottleContext) error {
	bottle := c.storage.BottleFromCreateBottlePayload(ctx.Payload)

	err := c.storage.Add(ctx, bottle)

	if err != nil {
		return ErrDatabaseError(err)
	}

	return ctx.Created()
}

// Update updates a bottle field(s).
func (c *BottleController) Update(ctx *app.UpdateBottleContext) error {
	err := c.storage.UpdateFromBottlePayload(ctx, ctx.Payload, ctx.BottleID)
	if err != nil {
		return ErrDatabaseError(err)
	}

	return ctx.NoContent()
}

// Delete removes a bottle from the database.
func (c *BottleController) Delete(ctx *app.DeleteBottleContext) error {
	err := c.storage.Delete(ctx, ctx.BottleID)
	if err != nil {
		return ErrDatabaseError(err)
	}

	return ctx.NoContent()
}

// Rate rates a bottle.
func (c *BottleController) Rate(ctx *app.RateBottleContext) error {
	_, err := c.storage.OneBottle(ctx, ctx.BottleID, ctx.AccountID)
	if err != nil {
		return ctx.NotFound()
	}

	updated := &app.BottlePayload{
		Rating: &ctx.Payload.Rating,
	}

	err = c.storage.UpdateFromBottlePayload(ctx, updated, ctx.BottleID)

	if err != nil {
		return ErrDatabaseError(err)
	}

	return ctx.NoContent()
}
