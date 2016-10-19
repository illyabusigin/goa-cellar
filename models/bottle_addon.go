package models

import (
	"github.com/goadesign/goa"
	"github.com/illyabusigin/goa-cellar/app"
	"golang.org/x/net/context"
	"time"
)

// ListBottleByYears returns an array of view: default.
func (m *BottleDB) ListBottleByYears(ctx context.Context, accountID int, years []int) (bottles []*app.Bottle, err error) {
	defer goa.MeasureSince([]string{"goa", "db", "bottle", "listbottlebyyears"}, time.Now())

	var native []*Bottle
	var objs []*app.Bottle
	err = m.Db.Scopes(BottleFilterByAccount(accountID, m.Db)).Table(m.TableName()).Preload("Account").Find(&native).Where("year in (?)", years).Error

	if err != nil {
		goa.LogError(ctx, "error listing Bottle", "error", err.Error())
		return nil, err
	}

	for _, t := range native {
		objs = append(objs, t.BottleToBottle())
	}

	return objs, nil
}
