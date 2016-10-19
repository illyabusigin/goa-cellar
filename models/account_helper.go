//************************************************************************//
// API "cellar": Model Helpers
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/illyabusigin/goa-cellar/design
// --out=$(GOPATH)/src/github.com/illyabusigin/goa-cellar
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package models

import (
	"github.com/goadesign/goa"
	"github.com/illyabusigin/goa-cellar/app"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
	"time"
)

// MediaType Retrieval Functions

// ListAccount returns an array of view: default.
func (m *AccountDB) ListAccount(ctx context.Context) []*app.Account {
	defer goa.MeasureSince([]string{"goa", "db", "account", "listaccount"}, time.Now())

	var native []*Account
	var objs []*app.Account
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Account", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.AccountToAccount())
	}

	return objs
}

// AccountToAccount loads a Account and builds the default view of media type Account.
func (m *Account) AccountToAccount() *app.Account {
	account := &app.Account{}
	account.CreatedAt = m.CreatedAt
	account.ID = m.ID
	account.Name = m.Name

	return account
}

// OneAccount loads a Account and builds the default view of media type Account.
func (m *AccountDB) OneAccount(ctx context.Context, id int) (*app.Account, error) {
	defer goa.MeasureSince([]string{"goa", "db", "account", "oneaccount"}, time.Now())

	var native Account
	err := m.Db.Scopes().Table(m.TableName()).Preload("Bottles").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Account", "error", err.Error())
		return nil, err
	}

	view := *native.AccountToAccount()
	return &view, err
}

// MediaType Retrieval Functions

// ListAccountLink returns an array of view: link.
func (m *AccountDB) ListAccountLink(ctx context.Context) []*app.AccountLink {
	defer goa.MeasureSince([]string{"goa", "db", "account", "listaccountlink"}, time.Now())

	var native []*Account
	var objs []*app.AccountLink
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Account", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.AccountToAccountLink())
	}

	return objs
}

// AccountToAccountLink loads a Account and builds the link view of media type Account.
func (m *Account) AccountToAccountLink() *app.AccountLink {
	account := &app.AccountLink{}
	account.ID = m.ID

	return account
}

// OneAccountLink loads a Account and builds the link view of media type Account.
func (m *AccountDB) OneAccountLink(ctx context.Context, id int) (*app.AccountLink, error) {
	defer goa.MeasureSince([]string{"goa", "db", "account", "oneaccountlink"}, time.Now())

	var native Account
	err := m.Db.Scopes().Table(m.TableName()).Preload("Bottles").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Account", "error", err.Error())
		return nil, err
	}

	view := *native.AccountToAccountLink()
	return &view, err
}

// MediaType Retrieval Functions

// ListAccountTiny returns an array of view: tiny.
func (m *AccountDB) ListAccountTiny(ctx context.Context) []*app.AccountTiny {
	defer goa.MeasureSince([]string{"goa", "db", "account", "listaccounttiny"}, time.Now())

	var native []*Account
	var objs []*app.AccountTiny
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Account", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.AccountToAccountTiny())
	}

	return objs
}

// AccountToAccountTiny loads a Account and builds the tiny view of media type Account.
func (m *Account) AccountToAccountTiny() *app.AccountTiny {
	account := &app.AccountTiny{}
	account.ID = m.ID
	account.Name = m.Name

	return account
}

// OneAccountTiny loads a Account and builds the tiny view of media type Account.
func (m *AccountDB) OneAccountTiny(ctx context.Context, id int) (*app.AccountTiny, error) {
	defer goa.MeasureSince([]string{"goa", "db", "account", "oneaccounttiny"}, time.Now())

	var native Account
	err := m.Db.Scopes().Table(m.TableName()).Preload("Bottles").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Account", "error", err.Error())
		return nil, err
	}

	view := *native.AccountToAccountTiny()
	return &view, err
}
