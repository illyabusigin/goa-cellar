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

// ListBottle returns an array of view: default.
func (m *BottleDB) ListBottle(ctx context.Context, accountID int) []*app.Bottle {
	defer goa.MeasureSince([]string{"goa", "db", "bottle", "listbottle"}, time.Now())

	var native []*Bottle
	var objs []*app.Bottle
	err := m.Db.Scopes(BottleFilterByAccount(accountID, m.Db)).Table(m.TableName()).Preload("Account").Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Bottle", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.BottleToBottle())
	}

	return objs
}

// BottleToBottle loads a Bottle and builds the default view of media type Bottle.
func (m *Bottle) BottleToBottle() *app.Bottle {
	bottle := &app.Bottle{}
	tmp1 := m.Account.AccountToAccountLink()
	bottle.Links = &app.BottleLinks{Account: tmp1}
	tmp2 := &m.Account
	bottle.Account = tmp2.AccountToAccountTiny() // %!s(MISSING)
	bottle.ID = m.ID
	bottle.Name = m.Name
	bottle.Rating = &m.Rating
	bottle.Varietal = m.Varietal
	bottle.Vineyard = m.Vineyard
	bottle.Vintage = m.Vintage

	return bottle
}

// OneBottle loads a Bottle and builds the default view of media type Bottle.
func (m *BottleDB) OneBottle(ctx context.Context, id int, accountID int) (*app.Bottle, error) {
	defer goa.MeasureSince([]string{"goa", "db", "bottle", "onebottle"}, time.Now())

	var native Bottle
	err := m.Db.Scopes(BottleFilterByAccount(accountID, m.Db)).Table(m.TableName()).Preload("Account").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Bottle", "error", err.Error())
		return nil, err
	}

	view := *native.BottleToBottle()
	return &view, err
}

// MediaType Retrieval Functions

// ListBottleFull returns an array of view: full.
func (m *BottleDB) ListBottleFull(ctx context.Context, accountID int) []*app.BottleFull {
	defer goa.MeasureSince([]string{"goa", "db", "bottle", "listbottlefull"}, time.Now())

	var native []*Bottle
	var objs []*app.BottleFull
	err := m.Db.Scopes(BottleFilterByAccount(accountID, m.Db)).Table(m.TableName()).Preload("Account").Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Bottle", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.BottleToBottleFull())
	}

	return objs
}

// BottleToBottleFull loads a Bottle and builds the full view of media type Bottle.
func (m *Bottle) BottleToBottleFull() *app.BottleFull {
	bottle := &app.BottleFull{}
	tmp1 := m.Account.AccountToAccountLink()
	bottle.Links = &app.BottleLinks{Account: tmp1}
	tmp2 := &m.Account
	bottle.Account = tmp2.AccountToAccount() // %!s(MISSING)
	bottle.Color = m.Color
	bottle.Country = m.Country
	bottle.CreatedAt = m.CreatedAt
	bottle.ID = m.ID
	bottle.Name = m.Name
	bottle.Rating = &m.Rating
	bottle.Region = m.Region
	bottle.Review = m.Review
	bottle.Sweetness = m.Sweetness
	bottle.UpdatedAt = &m.UpdatedAt
	bottle.Varietal = m.Varietal
	bottle.Vineyard = m.Vineyard
	bottle.Vintage = m.Vintage

	return bottle
}

// OneBottleFull loads a Bottle and builds the full view of media type Bottle.
func (m *BottleDB) OneBottleFull(ctx context.Context, id int, accountID int) (*app.BottleFull, error) {
	defer goa.MeasureSince([]string{"goa", "db", "bottle", "onebottlefull"}, time.Now())

	var native Bottle
	err := m.Db.Scopes(BottleFilterByAccount(accountID, m.Db)).Table(m.TableName()).Preload("Account").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Bottle", "error", err.Error())
		return nil, err
	}

	view := *native.BottleToBottleFull()
	return &view, err
}

// MediaType Retrieval Functions

// ListBottleTiny returns an array of view: tiny.
func (m *BottleDB) ListBottleTiny(ctx context.Context, accountID int) []*app.BottleTiny {
	defer goa.MeasureSince([]string{"goa", "db", "bottle", "listbottletiny"}, time.Now())

	var native []*Bottle
	var objs []*app.BottleTiny
	err := m.Db.Scopes(BottleFilterByAccount(accountID, m.Db)).Table(m.TableName()).Preload("Account").Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Bottle", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.BottleToBottleTiny())
	}

	return objs
}

// BottleToBottleTiny loads a Bottle and builds the tiny view of media type Bottle.
func (m *Bottle) BottleToBottleTiny() *app.BottleTiny {
	bottle := &app.BottleTiny{}
	tmp1 := m.Account.AccountToAccountLink()
	bottle.Links = &app.BottleLinks{Account: tmp1}
	bottle.ID = m.ID
	bottle.Name = m.Name
	bottle.Rating = &m.Rating

	return bottle
}

// OneBottleTiny loads a Bottle and builds the tiny view of media type Bottle.
func (m *BottleDB) OneBottleTiny(ctx context.Context, id int, accountID int) (*app.BottleTiny, error) {
	defer goa.MeasureSince([]string{"goa", "db", "bottle", "onebottletiny"}, time.Now())

	var native Bottle
	err := m.Db.Scopes(BottleFilterByAccount(accountID, m.Db)).Table(m.TableName()).Preload("Account").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Bottle", "error", err.Error())
		return nil, err
	}

	view := *native.BottleToBottleTiny()
	return &view, err
}
