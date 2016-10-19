package controllers

import (
	"github.com/goadesign/goa"
)

var ErrDatabaseError = goa.NewErrorClass("db_error", 500)
