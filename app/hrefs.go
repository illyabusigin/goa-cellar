//************************************************************************//
// API "cellar": Application Resource Href Factories
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/illyabusigin/goa-cellar/design
// --out=$(GOPATH)/src/github.com/illyabusigin/goa-cellar
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"fmt"
	"strings"
)

// AccountHref returns the resource href.
func AccountHref(accountID interface{}) string {
	paramaccountID := strings.TrimLeftFunc(fmt.Sprintf("%v", accountID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/cellar/accounts/%v", paramaccountID)
}

// BottleHref returns the resource href.
func BottleHref(accountID, bottleID interface{}) string {
	paramaccountID := strings.TrimLeftFunc(fmt.Sprintf("%v", accountID), func(r rune) bool { return r == '/' })
	parambottleID := strings.TrimLeftFunc(fmt.Sprintf("%v", bottleID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/cellar/accounts/%v/bottles/%v", paramaccountID, parambottleID)
}
