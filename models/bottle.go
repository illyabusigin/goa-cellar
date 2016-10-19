//************************************************************************//
// API "cellar": Models
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

// Bottle Model
type Bottle struct {
	ID        int `gorm:"primary_key"` // primary key
	AccountID int // Belongs To Account
	Color     string
	Country   *string
	CreatedAt time.Time
	DeletedAt *time.Time
	Name      string
	Rating    int
	Region    *string
	Review    *string
	Sweetness *int
	UpdatedAt time.Time
	Varietal  string
	Vineyard  string
	Vintage   int
	Account   Account
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m Bottle) TableName() string {
	return "bottles"

}

// BottleDB is the implementation of the storage interface for
// Bottle.
type BottleDB struct {
	Db *gorm.DB
}

// NewBottleDB creates a new storage type.
func NewBottleDB(db *gorm.DB) *BottleDB {
	return &BottleDB{Db: db}
}

// DB returns the underlying database.
func (m *BottleDB) DB() interface{} {
	return m.Db
}

// BottleStorage represents the storage interface.
type BottleStorage interface {
	DB() interface{}
	List(ctx context.Context) ([]*Bottle, error)
	Get(ctx context.Context, id int) (*Bottle, error)
	Add(ctx context.Context, bottle *Bottle) error
	Update(ctx context.Context, bottle *Bottle) error
	Delete(ctx context.Context, id int) error

	ListBottle(ctx context.Context, accountID int) []*app.Bottle
	OneBottle(ctx context.Context, id int, accountID int) (*app.Bottle, error)

	ListBottleFull(ctx context.Context, accountID int) []*app.BottleFull
	OneBottleFull(ctx context.Context, id int, accountID int) (*app.BottleFull, error)

	ListBottleTiny(ctx context.Context, accountID int) []*app.BottleTiny
	OneBottleTiny(ctx context.Context, id int, accountID int) (*app.BottleTiny, error)

	BottleFromCreateBottlePayload(payload *app.CreateBottlePayload) *Bottle
	UpdateFromBottlePayload(ctx context.Context, payload *app.BottlePayload, id int) error
	UpdateFromCreateBottlePayload(ctx context.Context, payload *app.CreateBottlePayload, id int) error
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *BottleDB) TableName() string {
	return "bottles"

}

// Belongs To Relationships

// BottleFilterByAccount is a gorm filter for a Belongs To relationship.
func BottleFilterByAccount(accountID int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {
	if accountID > 0 {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("account_id = ?", accountID)

		}
	}
	return func(db *gorm.DB) *gorm.DB { return db }
}

// CRUD Functions

// Get returns a single Bottle as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *BottleDB) Get(ctx context.Context, id int) (*Bottle, error) {
	defer goa.MeasureSince([]string{"goa", "db", "bottle", "get"}, time.Now())

	var native Bottle
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &native, err
}

// List returns an array of Bottle
func (m *BottleDB) List(ctx context.Context) ([]*Bottle, error) {
	defer goa.MeasureSince([]string{"goa", "db", "bottle", "list"}, time.Now())

	var objs []*Bottle
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, nil
}

// Add creates a new record.
func (m *BottleDB) Add(ctx context.Context, model *Bottle) error {
	defer goa.MeasureSince([]string{"goa", "db", "bottle", "add"}, time.Now())

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error adding Bottle", "error", err.Error())
		return err
	}

	return nil
}

// Update modifies a single record.
func (m *BottleDB) Update(ctx context.Context, model *Bottle) error {
	defer goa.MeasureSince([]string{"goa", "db", "bottle", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		goa.LogError(ctx, "error updating Bottle", "error", err.Error())
		return err
	}
	err = m.Db.Model(obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *BottleDB) Delete(ctx context.Context, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "bottle", "delete"}, time.Now())

	var obj Bottle

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.LogError(ctx, "error deleting Bottle", "error", err.Error())
		return err
	}

	return nil
}

// BottleFromBottlePayload Converts source BottlePayload to target Bottle model
// only copying the non-nil fields from the source.
func BottleFromBottlePayload(payload *app.BottlePayload) *Bottle {
	bottle := &Bottle{}
	if payload.Color != nil {
		bottle.Color = *payload.Color
	}
	if payload.Country != nil {
		bottle.Country = payload.Country
	}
	if payload.Name != nil {
		bottle.Name = *payload.Name
	}
	if payload.Rating != nil {
		bottle.Rating = *payload.Rating
	}
	if payload.Region != nil {
		bottle.Region = payload.Region
	}
	if payload.Review != nil {
		bottle.Review = payload.Review
	}
	if payload.Sweetness != nil {
		bottle.Sweetness = payload.Sweetness
	}
	if payload.Varietal != nil {
		bottle.Varietal = *payload.Varietal
	}
	if payload.Vineyard != nil {
		bottle.Vineyard = *payload.Vineyard
	}
	if payload.Vintage != nil {
		bottle.Vintage = *payload.Vintage
	}

	return bottle
}

// UpdateFromBottlePayload applies non-nil changes from BottlePayload to the model and saves it
func (m *BottleDB) UpdateFromBottlePayload(ctx context.Context, payload *app.BottlePayload, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "bottle", "updatefrombottlePayload"}, time.Now())

	var obj Bottle
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&obj).Error
	if err != nil {
		goa.LogError(ctx, "error retrieving Bottle", "error", err.Error())
		return err
	}
	if payload.Color != nil {
		obj.Color = *payload.Color
	}
	if payload.Country != nil {
		obj.Country = payload.Country
	}
	if payload.Name != nil {
		obj.Name = *payload.Name
	}
	if payload.Rating != nil {
		obj.Rating = *payload.Rating
	}
	if payload.Region != nil {
		obj.Region = payload.Region
	}
	if payload.Review != nil {
		obj.Review = payload.Review
	}
	if payload.Sweetness != nil {
		obj.Sweetness = payload.Sweetness
	}
	if payload.Varietal != nil {
		obj.Varietal = *payload.Varietal
	}
	if payload.Vineyard != nil {
		obj.Vineyard = *payload.Vineyard
	}
	if payload.Vintage != nil {
		obj.Vintage = *payload.Vintage
	}

	err = m.Db.Save(&obj).Error
	return err
}

// BottleFromCreateBottlePayload Converts source CreateBottlePayload to target Bottle model
// only copying the non-nil fields from the source.
func BottleFromCreateBottlePayload(payload *app.CreateBottlePayload) *Bottle {
	bottle := &Bottle{}
	bottle.Color = payload.Color
	if payload.Country != nil {
		bottle.Country = payload.Country
	}
	bottle.Name = payload.Name
	if payload.Rating != nil {
		bottle.Rating = *payload.Rating
	}
	if payload.Region != nil {
		bottle.Region = payload.Region
	}
	if payload.Review != nil {
		bottle.Review = payload.Review
	}
	if payload.Sweetness != nil {
		bottle.Sweetness = payload.Sweetness
	}
	bottle.Varietal = payload.Varietal
	bottle.Vineyard = payload.Vineyard
	bottle.Vintage = payload.Vintage

	return bottle
}

// UpdateFromCreateBottlePayload applies non-nil changes from CreateBottlePayload to the model and saves it
func (m *BottleDB) UpdateFromCreateBottlePayload(ctx context.Context, payload *app.CreateBottlePayload, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "bottle", "updatefromcreateBottlePayload"}, time.Now())

	var obj Bottle
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&obj).Error
	if err != nil {
		goa.LogError(ctx, "error retrieving Bottle", "error", err.Error())
		return err
	}
	obj.Color = payload.Color
	if payload.Country != nil {
		obj.Country = payload.Country
	}
	obj.Name = payload.Name
	if payload.Rating != nil {
		obj.Rating = *payload.Rating
	}
	if payload.Region != nil {
		obj.Region = payload.Region
	}
	if payload.Review != nil {
		obj.Review = payload.Review
	}
	if payload.Sweetness != nil {
		obj.Sweetness = payload.Sweetness
	}
	obj.Varietal = payload.Varietal
	obj.Vineyard = payload.Vineyard
	obj.Vintage = payload.Vintage

	err = m.Db.Save(&obj).Error
	return err
}
