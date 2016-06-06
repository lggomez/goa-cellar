//************************************************************************//
// User Types
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --design=github.com/goadesign/goa-cellar/design
// --out=$(GOPATH)/src/github.com/goadesign/goa-cellar
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package client

import (
	"net/http"
	"time"
)

// A tenant account
type Account struct {
	// Date of creation
	CreatedAt time.Time `json:"created_at" xml:"created_at"`
	// Email of account owner
	CreatedBy string `json:"created_by" xml:"created_by"`
	// API href of account
	Href string `json:"href" xml:"href"`
	// ID of account
	ID int `json:"id" xml:"id"`
	// Name of account
	Name string `json:"name" xml:"name"`
}

// DecodeAccount decodes the Account instance encoded in resp body.
func (c *Client) DecodeAccount(resp *http.Response) (*Account, error) {
	var decoded Account
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// BottleCollection media type is a collection of Bottle.
type BottleCollection []*Bottle

// DecodeBottleCollection decodes the BottleCollection instance encoded in resp body.
func (c *Client) DecodeBottleCollection(resp *http.Response) (BottleCollection, error) {
	var decoded BottleCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// A bottle of wine
type Bottle struct {
	// Account that owns bottle
	Account *Account `json:"account,omitempty" xml:"account,omitempty"`
	Color   string   `json:"color" xml:"color"`
	Country *string  `json:"country,omitempty" xml:"country,omitempty"`
	// Date of creation
	CreatedAt time.Time `json:"created_at" xml:"created_at"`
	// API href of bottle
	Href string `json:"href" xml:"href"`
	// ID of bottle
	ID   int    `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
	// Rating of bottle between 1 and 5
	Rating    *int    `json:"rating,omitempty" xml:"rating,omitempty"`
	Region    *string `json:"region,omitempty" xml:"region,omitempty"`
	Review    *string `json:"review,omitempty" xml:"review,omitempty"`
	Sweetness *int    `json:"sweetness,omitempty" xml:"sweetness,omitempty"`
	// Date of last update
	UpdatedAt *time.Time `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	Varietal  string     `json:"varietal" xml:"varietal"`
	Vineyard  string     `json:"vineyard" xml:"vineyard"`
	Vintage   int        `json:"vintage" xml:"vintage"`
}

// DecodeBottle decodes the Bottle instance encoded in resp body.
func (c *Client) DecodeBottle(resp *http.Response) (*Bottle, error) {
	var decoded Bottle
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
