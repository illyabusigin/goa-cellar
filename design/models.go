package design

import (
	"github.com/goadesign/gorma"
	. "github.com/goadesign/gorma/dsl"
)

var _ = StorageGroup("Cellar", func() {
	Description("This is the global storage group")
	Store("postgres", gorma.MySQL, func() {
		Description("This is the Postgres relational store")
		Model("Account", func() {
			RendersTo(Account)
			Description("Cellar Account")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
			})
			Field("name", gorma.String)
			HasMany("Bottles", "Bottle")
		})

		Model("Bottle", func() {
			BuildsFrom(func() {
				Payload("bottle", "create")
				Payload("bottle", "update")
			})
			RendersTo(Bottle)
			Field("id", gorma.Integer, func() {
				PrimaryKey()
			})
			Field("rating", gorma.Integer)
			Description("Bottle Model")
			BelongsTo("Account")
		})

	})
})
