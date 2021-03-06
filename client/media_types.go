//************************************************************************//
// API "cellar": Application Media Types
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/illyabusigin/goa-cellar/design
// --out=$(GOPATH)/src/github.com/illyabusigin/goa-cellar
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package client

import (
	"github.com/goadesign/goa"
	"net/http"
	"time"
	"unicode/utf8"
)

// A tenant account (default view)
//
// Identifier: application/vnd.account+json; view=default
type Account struct {
	// Date of creation
	CreatedAt time.Time `form:"created_at" json:"created_at" xml:"created_at"`
	// Email of account owner
	CreatedBy string `form:"created_by" json:"created_by" xml:"created_by"`
	// API href of account
	Href string `form:"href" json:"href" xml:"href"`
	// ID of account
	ID int `form:"id" json:"id" xml:"id"`
	// Name of account
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the Account media type instance.
func (mt *Account) Validate() (err error) {
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.CreatedBy == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "created_by"))
	}

	if err2 := goa.ValidateFormat(goa.FormatEmail, mt.CreatedBy); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`response.created_by`, mt.CreatedBy, goa.FormatEmail, err2))
	}
	return
}

// A tenant account (link view)
//
// Identifier: application/vnd.account+json; view=link
type AccountLink struct {
	// API href of account
	Href string `form:"href" json:"href" xml:"href"`
	// ID of account
	ID int `form:"id" json:"id" xml:"id"`
}

// Validate validates the AccountLink media type instance.
func (mt *AccountLink) Validate() (err error) {
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}

	return
}

// A tenant account (tiny view)
//
// Identifier: application/vnd.account+json; view=tiny
type AccountTiny struct {
	// API href of account
	Href string `form:"href" json:"href" xml:"href"`
	// ID of account
	ID int `form:"id" json:"id" xml:"id"`
	// Name of account
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the AccountTiny media type instance.
func (mt *AccountTiny) Validate() (err error) {
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}

	return
}

// DecodeAccount decodes the Account instance encoded in resp body.
func (c *Client) DecodeAccount(resp *http.Response) (*Account, error) {
	var decoded Account
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeAccountLink decodes the AccountLink instance encoded in resp body.
func (c *Client) DecodeAccountLink(resp *http.Response) (*AccountLink, error) {
	var decoded AccountLink
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeAccountTiny decodes the AccountTiny instance encoded in resp body.
func (c *Client) DecodeAccountTiny(resp *http.Response) (*AccountTiny, error) {
	var decoded AccountTiny
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// AccountCollection is the media type for an array of Account (default view)
//
// Identifier: application/vnd.account+json; type=collection; view=default
type AccountCollection []*Account

// Validate validates the AccountCollection media type instance.
func (mt AccountCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Href == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "href"))
		}
		if e.Name == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "name"))
		}
		if e.CreatedBy == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "created_by"))
		}

		if err2 := goa.ValidateFormat(goa.FormatEmail, e.CreatedBy); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`response[*].created_by`, e.CreatedBy, goa.FormatEmail, err2))
		}
	}
	return
}

// AccountCollection is the media type for an array of Account (link view)
//
// Identifier: application/vnd.account+json; type=collection; view=link
type AccountLinkCollection []*AccountLink

// Validate validates the AccountLinkCollection media type instance.
func (mt AccountLinkCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Href == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "href"))
		}

	}
	return
}

// AccountCollection is the media type for an array of Account (tiny view)
//
// Identifier: application/vnd.account+json; type=collection; view=tiny
type AccountTinyCollection []*AccountTiny

// Validate validates the AccountTinyCollection media type instance.
func (mt AccountTinyCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Href == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "href"))
		}
		if e.Name == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "name"))
		}

	}
	return
}

// DecodeAccountCollection decodes the AccountCollection instance encoded in resp body.
func (c *Client) DecodeAccountCollection(resp *http.Response) (AccountCollection, error) {
	var decoded AccountCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeAccountLinkCollection decodes the AccountLinkCollection instance encoded in resp body.
func (c *Client) DecodeAccountLinkCollection(resp *http.Response) (AccountLinkCollection, error) {
	var decoded AccountLinkCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeAccountTinyCollection decodes the AccountTinyCollection instance encoded in resp body.
func (c *Client) DecodeAccountTinyCollection(resp *http.Response) (AccountTinyCollection, error) {
	var decoded AccountTinyCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// A bottle of wine (default view)
//
// Identifier: application/vnd.bottle+json; view=default
type Bottle struct {
	// Account that owns bottle
	Account *AccountTiny `form:"account,omitempty" json:"account,omitempty" xml:"account,omitempty"`
	// API href of bottle
	Href string `form:"href" json:"href" xml:"href"`
	// ID of bottle
	ID int `form:"id" json:"id" xml:"id"`
	// Links to related resources
	Links *BottleLinks `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
	Name  string       `form:"name" json:"name" xml:"name"`
	// Rating of bottle between 1 and 5
	Rating   *int   `form:"rating,omitempty" json:"rating,omitempty" xml:"rating,omitempty"`
	Varietal string `form:"varietal" json:"varietal" xml:"varietal"`
	Vineyard string `form:"vineyard" json:"vineyard" xml:"vineyard"`
	Vintage  int    `form:"vintage" json:"vintage" xml:"vintage"`
}

// Validate validates the Bottle media type instance.
func (mt *Bottle) Validate() (err error) {
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Vineyard == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "vineyard"))
	}
	if mt.Varietal == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "varietal"))
	}

	if mt.Account != nil {
		if err2 := mt.Account.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if mt.Links != nil {
		if err2 := mt.Links.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if utf8.RuneCountInString(mt.Name) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, mt.Name, utf8.RuneCountInString(mt.Name), 2, true))
	}
	if mt.Rating != nil {
		if *mt.Rating < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.rating`, *mt.Rating, 1, true))
		}
	}
	if mt.Rating != nil {
		if *mt.Rating > 5 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.rating`, *mt.Rating, 5, false))
		}
	}
	if utf8.RuneCountInString(mt.Varietal) < 4 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.varietal`, mt.Varietal, utf8.RuneCountInString(mt.Varietal), 4, true))
	}
	if utf8.RuneCountInString(mt.Vineyard) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.vineyard`, mt.Vineyard, utf8.RuneCountInString(mt.Vineyard), 2, true))
	}
	if mt.Vintage < 1900 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.vintage`, mt.Vintage, 1900, true))
	}
	if mt.Vintage > 2020 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.vintage`, mt.Vintage, 2020, false))
	}
	return
}

// A bottle of wine (full view)
//
// Identifier: application/vnd.bottle+json; view=full
type BottleFull struct {
	// Account that owns bottle
	Account *Account `form:"account,omitempty" json:"account,omitempty" xml:"account,omitempty"`
	Color   string   `form:"color" json:"color" xml:"color"`
	Country *string  `form:"country,omitempty" json:"country,omitempty" xml:"country,omitempty"`
	// Date of creation
	CreatedAt time.Time `form:"created_at" json:"created_at" xml:"created_at"`
	// API href of bottle
	Href string `form:"href" json:"href" xml:"href"`
	// ID of bottle
	ID int `form:"id" json:"id" xml:"id"`
	// Links to related resources
	Links *BottleLinks `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
	Name  string       `form:"name" json:"name" xml:"name"`
	// Rating of bottle between 1 and 5
	Rating    *int    `form:"rating,omitempty" json:"rating,omitempty" xml:"rating,omitempty"`
	Region    *string `form:"region,omitempty" json:"region,omitempty" xml:"region,omitempty"`
	Review    *string `form:"review,omitempty" json:"review,omitempty" xml:"review,omitempty"`
	Sweetness *int    `form:"sweetness,omitempty" json:"sweetness,omitempty" xml:"sweetness,omitempty"`
	// Date of last update
	UpdatedAt *time.Time `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	Varietal  string     `form:"varietal" json:"varietal" xml:"varietal"`
	Vineyard  string     `form:"vineyard" json:"vineyard" xml:"vineyard"`
	Vintage   int        `form:"vintage" json:"vintage" xml:"vintage"`
}

// Validate validates the BottleFull media type instance.
func (mt *BottleFull) Validate() (err error) {
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Vineyard == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "vineyard"))
	}
	if mt.Varietal == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "varietal"))
	}
	if mt.Color == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "color"))
	}

	if mt.Account != nil {
		if err2 := mt.Account.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if !(mt.Color == "red" || mt.Color == "white" || mt.Color == "rose" || mt.Color == "yellow" || mt.Color == "sparkling") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.color`, mt.Color, []interface{}{"red", "white", "rose", "yellow", "sparkling"}))
	}
	if mt.Country != nil {
		if utf8.RuneCountInString(*mt.Country) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.country`, *mt.Country, utf8.RuneCountInString(*mt.Country), 2, true))
		}
	}
	if mt.Links != nil {
		if err2 := mt.Links.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if utf8.RuneCountInString(mt.Name) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, mt.Name, utf8.RuneCountInString(mt.Name), 2, true))
	}
	if mt.Rating != nil {
		if *mt.Rating < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.rating`, *mt.Rating, 1, true))
		}
	}
	if mt.Rating != nil {
		if *mt.Rating > 5 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.rating`, *mt.Rating, 5, false))
		}
	}
	if mt.Review != nil {
		if utf8.RuneCountInString(*mt.Review) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.review`, *mt.Review, utf8.RuneCountInString(*mt.Review), 3, true))
		}
	}
	if mt.Review != nil {
		if utf8.RuneCountInString(*mt.Review) > 300 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.review`, *mt.Review, utf8.RuneCountInString(*mt.Review), 300, false))
		}
	}
	if mt.Sweetness != nil {
		if *mt.Sweetness < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.sweetness`, *mt.Sweetness, 1, true))
		}
	}
	if mt.Sweetness != nil {
		if *mt.Sweetness > 5 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.sweetness`, *mt.Sweetness, 5, false))
		}
	}
	if utf8.RuneCountInString(mt.Varietal) < 4 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.varietal`, mt.Varietal, utf8.RuneCountInString(mt.Varietal), 4, true))
	}
	if utf8.RuneCountInString(mt.Vineyard) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.vineyard`, mt.Vineyard, utf8.RuneCountInString(mt.Vineyard), 2, true))
	}
	if mt.Vintage < 1900 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.vintage`, mt.Vintage, 1900, true))
	}
	if mt.Vintage > 2020 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.vintage`, mt.Vintage, 2020, false))
	}
	return
}

// A bottle of wine (tiny view)
//
// Identifier: application/vnd.bottle+json; view=tiny
type BottleTiny struct {
	// API href of bottle
	Href string `form:"href" json:"href" xml:"href"`
	// ID of bottle
	ID int `form:"id" json:"id" xml:"id"`
	// Links to related resources
	Links *BottleLinks `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
	Name  string       `form:"name" json:"name" xml:"name"`
	// Rating of bottle between 1 and 5
	Rating *int `form:"rating,omitempty" json:"rating,omitempty" xml:"rating,omitempty"`
}

// Validate validates the BottleTiny media type instance.
func (mt *BottleTiny) Validate() (err error) {
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}

	if mt.Links != nil {
		if err2 := mt.Links.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if utf8.RuneCountInString(mt.Name) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, mt.Name, utf8.RuneCountInString(mt.Name), 2, true))
	}
	if mt.Rating != nil {
		if *mt.Rating < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.rating`, *mt.Rating, 1, true))
		}
	}
	if mt.Rating != nil {
		if *mt.Rating > 5 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.rating`, *mt.Rating, 5, false))
		}
	}
	return
}

// BottleLinks contains links to related resources of Bottle.
type BottleLinks struct {
	Account *AccountLink `form:"account,omitempty" json:"account,omitempty" xml:"account,omitempty"`
}

// Validate validates the BottleLinks type instance.
func (ut *BottleLinks) Validate() (err error) {
	if ut.Account != nil {
		if err2 := ut.Account.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// DecodeBottle decodes the Bottle instance encoded in resp body.
func (c *Client) DecodeBottle(resp *http.Response) (*Bottle, error) {
	var decoded Bottle
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeBottleFull decodes the BottleFull instance encoded in resp body.
func (c *Client) DecodeBottleFull(resp *http.Response) (*BottleFull, error) {
	var decoded BottleFull
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeBottleTiny decodes the BottleTiny instance encoded in resp body.
func (c *Client) DecodeBottleTiny(resp *http.Response) (*BottleTiny, error) {
	var decoded BottleTiny
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// BottleCollection is the media type for an array of Bottle (default view)
//
// Identifier: application/vnd.bottle+json; type=collection; view=default
type BottleCollection []*Bottle

// Validate validates the BottleCollection media type instance.
func (mt BottleCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Href == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "href"))
		}
		if e.Name == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "name"))
		}
		if e.Vineyard == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "vineyard"))
		}
		if e.Varietal == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "varietal"))
		}

		if e.Account != nil {
			if err2 := e.Account.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if e.Links != nil {
			if err2 := e.Links.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if utf8.RuneCountInString(e.Name) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response[*].name`, e.Name, utf8.RuneCountInString(e.Name), 2, true))
		}
		if e.Rating != nil {
			if *e.Rating < 1 {
				err = goa.MergeErrors(err, goa.InvalidRangeError(`response[*].rating`, *e.Rating, 1, true))
			}
		}
		if e.Rating != nil {
			if *e.Rating > 5 {
				err = goa.MergeErrors(err, goa.InvalidRangeError(`response[*].rating`, *e.Rating, 5, false))
			}
		}
		if utf8.RuneCountInString(e.Varietal) < 4 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response[*].varietal`, e.Varietal, utf8.RuneCountInString(e.Varietal), 4, true))
		}
		if utf8.RuneCountInString(e.Vineyard) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response[*].vineyard`, e.Vineyard, utf8.RuneCountInString(e.Vineyard), 2, true))
		}
		if e.Vintage < 1900 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response[*].vintage`, e.Vintage, 1900, true))
		}
		if e.Vintage > 2020 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response[*].vintage`, e.Vintage, 2020, false))
		}
	}
	return
}

// BottleCollection is the media type for an array of Bottle (tiny view)
//
// Identifier: application/vnd.bottle+json; type=collection; view=tiny
type BottleTinyCollection []*BottleTiny

// Validate validates the BottleTinyCollection media type instance.
func (mt BottleTinyCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Href == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "href"))
		}
		if e.Name == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "name"))
		}

		if e.Links != nil {
			if err2 := e.Links.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if utf8.RuneCountInString(e.Name) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response[*].name`, e.Name, utf8.RuneCountInString(e.Name), 2, true))
		}
		if e.Rating != nil {
			if *e.Rating < 1 {
				err = goa.MergeErrors(err, goa.InvalidRangeError(`response[*].rating`, *e.Rating, 1, true))
			}
		}
		if e.Rating != nil {
			if *e.Rating > 5 {
				err = goa.MergeErrors(err, goa.InvalidRangeError(`response[*].rating`, *e.Rating, 5, false))
			}
		}
	}
	return
}

// BottleLinksArray contains links to related resources of BottleCollection.
type BottleLinksArray []*BottleLinks

// Validate validates the BottleLinksArray type instance.
func (ut BottleLinksArray) Validate() (err error) {
	for _, e := range ut {
		if e.Account != nil {
			if err2 := e.Account.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeBottleCollection decodes the BottleCollection instance encoded in resp body.
func (c *Client) DecodeBottleCollection(resp *http.Response) (BottleCollection, error) {
	var decoded BottleCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeBottleTinyCollection decodes the BottleTinyCollection instance encoded in resp body.
func (c *Client) DecodeBottleTinyCollection(resp *http.Response) (BottleTinyCollection, error) {
	var decoded BottleTinyCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeErrorResponse decodes the ErrorResponse instance encoded in resp body.
func (c *Client) DecodeErrorResponse(resp *http.Response) (*goa.ErrorResponse, error) {
	var decoded goa.ErrorResponse
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
