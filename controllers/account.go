package controllers

import (
	"github.com/goadesign/goa"
	"github.com/illyabusigin/goa-cellar/app"
	"github.com/illyabusigin/goa-cellar/models"
	"github.com/jinzhu/gorm"
)

// AccountController implements the account resource.
type AccountController struct {
	*goa.Controller
	storage models.AccountStorage
}

// NewAccount creates a bottle controller.
func NewAccount(service *goa.Service, db *gorm.DB) *AccountController {
	return &AccountController{
		Controller: service.NewController("Account"),
	}
}

// Create runs the create action.
func (c *AccountController) Create(ctx *app.CreateAccountContext) error {
	a := models.Account{}
	a.Name = ctx.Payload.Name
	err := c.storage.Add(ctx.Context, &a)
	if err != nil {
		return ErrDatabaseError(err)
	}
	ctx.ResponseData.Header().Set("Location", app.AccountHref(a.ID))
	return ctx.Created()
}

// Delete runs the delete action.
func (c *AccountController) Delete(ctx *app.DeleteAccountContext) error {
	err := c.storage.Delete(ctx.Context, ctx.AccountID)
	if err != nil {
		return ErrDatabaseError(err)
	}
	return ctx.NoContent()
}

// Show runs the show action.
func (c *AccountController) Show(ctx *app.ShowAccountContext) error {
	account, err := c.storage.OneAccount(ctx.Context, ctx.AccountID)
	if err == gorm.ErrRecordNotFound {
		return ctx.NotFound()
	} else if err != nil {
		return ErrDatabaseError(err)
	}
	account.Href = app.AccountHref(account.ID)
	return ctx.OK(account)
}

// List
func (c *AccountController) List(ctx *app.ListAccountContext) error {
	accounts := c.storage.ListAccount(ctx.Context)
	return ctx.OK(accounts)

}

// Update runs the update action.
func (c *AccountController) Update(ctx *app.UpdateAccountContext) error {
	m, err := c.storage.Get(ctx.Context, ctx.AccountID)
	if err == gorm.ErrRecordNotFound {
		return ctx.NotFound()
	}
	m.Name = ctx.Payload.Name
	err = c.storage.Update(ctx, m)
	if err != nil {
		return ErrDatabaseError(err)
	}
	return ctx.NoContent()
}
